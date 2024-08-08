package handlers

import (
	"sync"
)

type FlashMessage struct {
	Type    string
	Content string
}

type MessageStore struct {
	messages map[string]FlashMessage
	mutex    sync.RWMutex
}

var globalMessageStore = &MessageStore{
	messages: make(map[string]FlashMessage),
}

func (ms *MessageStore) Set(key string, message FlashMessage) {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()
	ms.messages[key] = message
}

func (ms *MessageStore) Get(key string) (FlashMessage, bool) {
	ms.mutex.RLock()
	defer ms.mutex.RUnlock()
	message, exists := ms.messages[key]
	return message, exists
}

func (ms *MessageStore) Delete(key string) {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()
	delete(ms.messages, key)
}

type UtilHandler struct{}

// func NewUtilHandler() *UtilHandler {
// 	return &UtilHandler{}
// }

// func (h *UtilHandler) FlashHandler(c *gin.Context) {
// 	taskID := c.Param("taskID")

// 	if result, exists := globalMessageStore.Get(taskID); exists {
// 		c.Set("flash", result)
// 		globalMessageStore.Delete(taskID)
// 	}

// 	c.HTML(http.StatusOK, "_flash.html", gin.H{
// 		"flash": c.MustGet("flash"),
// 	})
// }

// func (h *UtilHandler) RoutesHandler(c *gin.Context) {
// 	routes := []gin.RouteInfo{}
// 	for _, r := range gin.Default().Routes() {
// 		routes = append(routes, r)
// 	}

// 	c.HTML(http.StatusOK, "utils/routes", gin.H{
// 		"routes":  routes,
// 		"Request": c.Request,
// 	})
// }
