package actions

import (
	"context"
	"sync"
	"time"

	"github.com/gobuffalo/buffalo"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

var (
	clients      = make(map[string][]*websocket.Conn)
	clientsMutex sync.Mutex
)

// WebSocketHandler handles WebSocket connections
func WebSocketHandler(c buffalo.Context) error {
	taskID := c.Param("taskID")

	conn, err := websocket.Accept(c.Response(), c.Request(), &websocket.AcceptOptions{
		InsecureSkipVerify: true, // Be cautious with this in production
	})
	if err != nil {
		c.Logger().Errorf("Error accepting WebSocket connection: %v", err)
		return err
	}
	defer conn.Close(websocket.StatusInternalError, "the sky is falling")

	clientsMutex.Lock()
	clients[taskID] = append(clients[taskID], conn)
	clientsMutex.Unlock()

	// Keep the connection open
	for {
		_, _, err := conn.Read(context.Background())
		if err != nil {
			if websocket.CloseStatus(err) == websocket.StatusNormalClosure {
				c.Logger().Info("WebSocket connection closed by the client")
			} else {
				c.Logger().Errorf("Unexpected error: %v", err)
			}
			// Remove the client when the connection is closed
			clientsMutex.Lock()
			for i, client := range clients[taskID] {
				if client == conn {
					clients[taskID] = append(clients[taskID][:i], clients[taskID][i+1:]...)
					break
				}
			}
			clientsMutex.Unlock()
			conn.Close(websocket.StatusNormalClosure, "")
			break
		}
	}

	return nil
}

// SendWebSocketUpdate sends an update to all connected clients for a specific task
func SendWebSocketUpdate(taskID string, message interface{}) {
	clientsMutex.Lock()
	defer clientsMutex.Unlock()

	for i, conn := range clients[taskID] {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		err := wsjson.Write(ctx, conn, message)
		cancel()
		if err != nil {
			// Handle error by removing the disconnected client
			clients[taskID] = append(clients[taskID][:i], clients[taskID][i+1:]...)
			conn.Close(websocket.StatusInternalError, "failed to send message")
			// Optionally, log the error
			// log.Printf("Error sending WebSocket message: %v", err)
		}
	}
}
