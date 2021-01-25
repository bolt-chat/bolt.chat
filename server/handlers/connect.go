package handlers

import (
	"bytes"
	"encoding/json"
	"net"
	"os"

	"github.com/bolt-chat/protocol"
	"github.com/bolt-chat/protocol/events"
	"github.com/bolt-chat/server/logging"
	"github.com/bolt-chat/util"
)

/*
HandleConnection handles a TCP connection
during its entire lifespan.
*/
func HandleConnection(pool *util.ConnPool, conn *net.TCPConn) {
	var user *protocol.User

	for {
		// a := server.Listener{}
		b := make([]byte, 4096)

		// Wait for and receive incoming events
		_, connErr := conn.Read(b)

		if connErr != nil {
			// Broadcast a disconnect message
			evt := *events.NewLeaveEvent(user) // TODO:
			evtRaw, _ := json.Marshal(evt)
			util.Broadcast(pool, evt) // TODO:
			logging.LogEvent(string(evtRaw))
			pool.RemoveFromPool(conn)
			return
		}

		// Trim empty bytes at the end
		b = bytes.TrimRight(b, "\x00")

		// Log raw events in debug mode
		logging.LogEvent(string(b))

		evt := &events.BaseEvent{}

		// Decode the event
		err := json.Unmarshal(b, evt)

		if err != nil {
			util.WriteJson(conn, *events.NewErrorEvent("invalid_format"))
			conn.Close()
			return
		}

		switch evt.Event.Type {
		case events.MessageType:
			msgEvt := &events.MessageEvent{}
			json.Unmarshal(b, msgEvt)
			util.Broadcast(pool, msgEvt) // TODO: mutate and write
		case events.JoinType:
			joinEvt := &events.JoinEvent{}
			json.Unmarshal(b, joinEvt)
			user = joinEvt.User

			motd, hasMotd := os.LookupEnv("MOTD") // Get MOTD env
			if hasMotd == true {
				// Send MOTD if env var is declared
				util.WriteJson(conn, *events.NewMotdEvent(motd))
			}

			util.Broadcast(pool, joinEvt)
		default:
			util.WriteJson(conn, *events.NewErrorEvent("invalid_event"))
		}
	}
}
