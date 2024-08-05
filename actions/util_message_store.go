package actions

import (
	"sync"
)

type MessageStore struct {
	messages map[string]StoreMessage
	mutex    sync.RWMutex
}

type StoreMessage struct {
	Type    string
	Content string
}

var globalMessageStore = &MessageStore{
	messages: make(map[string]StoreMessage),
}

func (ms *MessageStore) Set(jobID string, message StoreMessage) {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()
	ms.messages[jobID] = message
}

func (ms *MessageStore) Get(jobID string) (StoreMessage, bool) {
	ms.mutex.RLock()
	defer ms.mutex.RUnlock()
	message, exists := ms.messages[jobID]
	return message, exists
}

func (ms *MessageStore) Delete(jobID string) {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()
	delete(ms.messages, jobID)
}
