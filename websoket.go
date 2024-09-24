package main

import (
	"context"
	"net/http"
	"sync"
	"encoding/json"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	"github.com/nbd-wtf/go-nostr"
)

type Relay struct {
	connections map[*websocket.Conn]bool
	mutex       sync.Mutex
}

func (r *Relay) addConnection(conn *websocket.Conn) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.connections[conn] = true
}

func (r *Relay) removeConnection(conn *websocket.Conn) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	delete(r.connections, conn)
}

func (r *Relay) broadcast(message []byte) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for conn := range r.connections {
		err := conn.Write(context.Background(), websocket.MessageText, message)
		if err != nil {
			pLog(LEng, LWarn, "error broadcasting to connection: ", err)
			conn.CloseNow()
			delete(r.connections, conn)
		}
	}
}

func (r *Relay) handleConnection(conn *websocket.Conn) {
	defer func() {
		r.removeConnection(conn)
		conn.CloseNow()
	}()

	for {
		_, reader, err := conn.Reader(context.Background())
		if err != nil {
			pLog(LEng, LWarn, "read error: ", err)
			break
		}

		message := make([]byte, 512)
		n, err := reader.Read(message)
		if err != nil {
			pLog(LEng, LWarn, "error reading message: ", err)
			break
		}

		message = message[:n]
		pLog(LEng, LDebug, "Received: ", message)

		// Basic Nostr message handling
		switch {
		case string(message[:6]) == "EVENT ":
			// Broadcast the event to all connected clients
			r.broadcast(message)
		case string(message[:4]) == "REQ ":
			// Handle subscription requests
			// You might add logic here to filter events based on the REQ message
			conn.Write(context.Background(),
                websocket.MessageText,
                []byte("Acknowledged REQ"))
		}
	}
}

func eventHandler(ctx context.Context, conn *websocket.Conn, data []json.RawMessage) {
    var event nostr.Event
    if err := event.UnmarshalJSON(data[1]); err != nil {
		wsjson.Write(ctx, conn, [2]string{"NOTICE", "error: invalid EVENT"})
		return
	}

    id := event.GetID()
    res := [4]interface{}{"OK", id, true, ""}

    _ = event.Content

	wsjson.Write(ctx, conn, res)
	pLog(LEng, LDebug, "Sent message: ", res)
}

func wsHandler(w http.ResponseWriter, r *http.Request, ip string, ua string) {
	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
    	InsecureSkipVerify: true,
    })
    if err != nil {
    	pLog(LEng, LWarn, " Error failed to Accept message: ", err)
    	return
    }
	defer conn.Close(websocket.StatusNormalClosure, "Normal closure")

    pLog(LEng, LInfo, ip, " connected ", ua)

    ctx := r.Context()

	for {
		_, msg, err := conn.Read(ctx)
		if websocket.CloseStatus(err) == websocket.StatusNormalClosure {
			pLog(LEng, LDebug, "Client closed the connection normally.")
			return
		} else if err != nil {
			pLog(LEng, LWarn, "Error reading message: ", err)
			return
		}

		pLog(LEng, LDebug, "Received message: ", string(msg))

        var data []json.RawMessage

        if err := json.Unmarshal(msg, &data); err != nil {
			// doesn't looks right
			wsjson.Write(ctx, conn,
			[2]string{
                "NOTICE",
                "error: your json doesn't looks right.",
            })
			continue
		}

        if len(data) < 1 {
			wsjson.Write(ctx, conn,
			[2]string{
			    "NOTICE",
                "error: does not looks like there's something in your message.",
            })
			continue
		}

		var cmd string
		if err := json.Unmarshal(data[0], &cmd); err != nil {
			wsjson.Write(ctx, conn, [2]string{
                "NOTICE",
                "error: please check your command.",
		    })
			continue
		}

		switch cmd {
    		case "REQ":
    			if len(data) < 3 {
    				wsjson.Write(ctx, conn, [2]string{
    				"NOTICE", "error: invalid REQ"})
    				continue
    			}

                // s.ClientREQ <- data
    		case "CLOSE":
    			if len(data) < 2 {
    				wsjson.Write(ctx, conn, [2]string{
    				"NOTICE", "error: invalid CLOSE"})
    				continue
    			}

                // s.ClientCLOSE <- data
    		case "EVENT":
    			if len(data) < 2 {
    				wsjson.Write(ctx, conn, [2]string{
    				"NOTICE", "error: invalid EVENT"})
    				continue
    			}

                eventHandler(ctx, conn, data)
                // s.ClientEVENT <- data
    		default:
    			wsjson.Write(ctx, conn, [2]string{
                    "NOTICE",
                    "error: unknown command " + cmd,
    			})
		}

        // err = conn.Write(r.Context(), websocket.MessageText, data)
        // if err != nil {
        // 	pLog(LEng, LWarn, "Error writing message:", err)
        // 	return
        // }
	}
}
