package ws

import (
	"math/rand"
	"sync"

	"git.solsynth.dev/hypernet/nexus/pkg/internal/directory"
	"git.solsynth.dev/hypernet/nexus/pkg/nex/sec"
	"github.com/rs/zerolog/log"

	"github.com/gofiber/contrib/websocket"
)

var (
	wsMutex sync.Mutex
	wsConn  = make(map[uint]map[uint64]*websocket.Conn)
)

func ClientRegister(user sec.UserInfo, conn *websocket.Conn) uint64 {
	wsMutex.Lock()
	if wsConn[user.ID] == nil {
		wsConn[user.ID] = make(map[uint64]*websocket.Conn)
	}
	clientId := rand.Uint64()
	wsConn[user.ID][clientId] = conn
	wsMutex.Unlock()

	log.Debug().
		Uint64("client_id", clientId).
		Uint("user_id", user.ID).
		Msg("An client connected to stream endpoint...")

	_ = directory.BroadcastEvent("ws.client.register", map[string]any{
		"user": user.ID,
		"id":   clientId,
	})

	return clientId
}

func ClientUnregister(user sec.UserInfo, id uint64) {
	wsMutex.Lock()
	if wsConn[user.ID] == nil {
		wsConn[user.ID] = make(map[uint64]*websocket.Conn)
	}
	delete(wsConn[user.ID], id)
	wsMutex.Unlock()

	log.Debug().
		Uint64("client_id", id).
		Uint("user_id", user.ID).
		Msg("An client disconnected from stream endpoint...")

	_ = directory.BroadcastEvent("ws.client.unregister", map[string]any{
		"user": user.ID,
		"id":   id,
	})
}

func ClientCount(uid uint) int {
	return len(wsConn[uid])
}

func WebsocketPush(uid uint, body []byte) (count int, successes []uint64, errs []error) {
	for _, conn := range wsConn[uid] {
		if err := conn.WriteMessage(1, body); err != nil {
			errs = append(errs, err)
		} else {
			successes = append(successes, uint64(uid))
		}
		count++
	}
	return
}

func WebsocketPushDirect(clientId uint64, body []byte) (count int, successes []uint64, errs []error) {
	for _, m := range wsConn {
		if conn, ok := m[clientId]; ok {
			if err := conn.WriteMessage(1, body); err != nil {
				errs = append(errs, err)
			} else {
				successes = append(successes, clientId)
			}
			count++
		}
	}
	return
}

func WebsocketPushBatch(uidList []uint, body []byte) (count int, successes []uint64, errs []error) {
	for _, uid := range uidList {
		for _, conn := range wsConn[uid] {
			if err := conn.WriteMessage(1, body); err != nil {
				errs = append(errs, err)
			} else {
				successes = append(successes, uint64(uid))
			}
			count++
		}
	}
	return
}

func WebsocketPushBatchDirect(clientIdList []uint64, body []byte) (count int, successes []uint64, errs []error) {
	for _, clientId := range clientIdList {
		for _, m := range wsConn {
			if conn, ok := m[clientId]; ok {
				if err := conn.WriteMessage(1, body); err != nil {
					errs = append(errs, err)
				} else {
					successes = append(successes, clientId)
				}
				count++
			}
		}
	}
	return
}
