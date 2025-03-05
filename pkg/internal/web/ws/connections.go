package ws

import (
	"fmt"
	"sync"

	"git.solsynth.dev/hypernet/nexus/pkg/internal/directory"
	"git.solsynth.dev/hypernet/nexus/pkg/nex/sec"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"github.com/gofiber/contrib/websocket"
)

var (
	wsMutex sync.Mutex
	wsConn  = make(map[uint]map[string]*websocket.Conn)
)

func ClientRegister(user sec.UserInfo, conn *websocket.Conn) (string, error) {
	wsMutex.Lock()
	if wsConn[user.ID] == nil {
		wsConn[user.ID] = make(map[string]*websocket.Conn)
	}
	var clientId string
	if userDefinedId := conn.Query("clientId"); len(userDefinedId) > 0 && len(userDefinedId) <= 16 {
		clientId = userDefinedId
	} else {
		clientId = uuid.NewString()
	}
	if _, ok := wsConn[user.ID][clientId]; ok {
		return clientId, fmt.Errorf("client already conncted")
	}
	wsConn[user.ID][clientId] = conn
	wsMutex.Unlock()

	log.Debug().
		Str("client_id", clientId).
		Uint("user_id", user.ID).
		Msg("An client connected to stream endpoint...")

	_ = directory.BroadcastEvent("ws.client.register", map[string]any{
		"user": user.ID,
		"id":   clientId,
	})

	return clientId, nil
}

func ClientUnregister(user sec.UserInfo, id string) {
	wsMutex.Lock()
	if wsConn[user.ID] == nil {
		wsConn[user.ID] = make(map[string]*websocket.Conn)
	}
	delete(wsConn[user.ID], id)
	wsMutex.Unlock()

	log.Debug().
		Str("client_id", id).
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

func WebsocketPush(uid uint, body []byte) (count int, successes []string, errs []error) {
	for _, conn := range wsConn[uid] {
		if err := conn.WriteMessage(1, body); err != nil {
			errs = append(errs, err)
		} else {
			successes = append(successes, fmt.Sprintf("%d", uid))
		}
		count++
	}
	return
}

func WebsocketPushDirect(clientId string, body []byte) (count int, successes []string, errs []error) {
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

func WebsocketPushBatch(uidList []uint, body []byte) (count int, successes []string, errs []error) {
	for _, uid := range uidList {
		for _, conn := range wsConn[uid] {
			if err := conn.WriteMessage(1, body); err != nil {
				errs = append(errs, err)
			} else {
				successes = append(successes, fmt.Sprintf("%d", uid))
			}
			count++
		}
	}
	return
}

func WebsocketPushBatchDirect(clientIdList []string, body []byte) (count int, successes []string, errs []error) {
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
