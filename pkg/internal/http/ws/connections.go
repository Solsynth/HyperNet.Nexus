package ws

import (
	"git.solsynth.dev/hypernet/nexus/pkg/internal/directory"
	"math/rand"
	"sync"

	"git.solsynth.dev/hypernet/nexus/pkg/internal/models"
	"github.com/gofiber/contrib/websocket"
)

var (
	wsMutex sync.Mutex
	wsConn  = make(map[uint]map[uint64]*websocket.Conn)
)

func ClientRegister(user models.Account, conn *websocket.Conn) uint64 {
	wsMutex.Lock()
	if wsConn[user.ID] == nil {
		wsConn[user.ID] = make(map[uint64]*websocket.Conn)
	}
	clientId := rand.Uint64()
	wsConn[user.ID][clientId] = conn
	wsMutex.Unlock()

	directory.BroadcastEvent("ws.client.register", map[string]any{
		"user": user.ID,
		"id":   clientId,
	})

	return clientId
}

func ClientUnregister(user models.Account, id uint64) {
	wsMutex.Lock()
	if wsConn[user.ID] == nil {
		wsConn[user.ID] = make(map[uint64]*websocket.Conn)
	}
	delete(wsConn[user.ID], id)
	wsMutex.Unlock()

	directory.BroadcastEvent("ws.client.unregister", map[string]any{
		"user": user.ID,
		"id":   id,
	})
}

func ClientCount(uid uint) int {
	return len(wsConn[uid])
}

func WebsocketPush(uid uint, body []byte) (count int, success int, errs []error) {
	for _, conn := range wsConn[uid] {
		if err := conn.WriteMessage(1, body); err != nil {
			errs = append(errs, err)
		} else {
			success++
		}
		count++
	}
	return
}

func WebsocketPushDirect(clientId uint64, body []byte) (count int, success int, errs []error) {
	for _, m := range wsConn {
		if conn, ok := m[clientId]; ok {
			if err := conn.WriteMessage(1, body); err != nil {
				errs = append(errs, err)
			} else {
				success++
			}
			count++
		}
	}
	return
}

func WebsocketPushBatch(uidList []uint, body []byte) (count int, success int, errs []error) {
	for _, uid := range uidList {
		for _, conn := range wsConn[uid] {
			if err := conn.WriteMessage(1, body); err != nil {
				errs = append(errs, err)
			} else {
				success++
			}
			count++
		}
	}
	return
}

func WebsocketPushBatchDirect(clientIdList []uint64, body []byte) (count int, success int, errs []error) {
	for _, clientId := range clientIdList {
		for _, m := range wsConn {
			if conn, ok := m[clientId]; ok {
				if err := conn.WriteMessage(1, body); err != nil {
					errs = append(errs, err)
				} else {
					success++
				}
				count++
			}
		}
	}
	return
}
