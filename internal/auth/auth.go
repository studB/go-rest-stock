package auth

import (
	"fmt"
	"sync"
	"time"
)

type TokenManager struct {
	accessToken string
	expiresAt   time.Time
	mu          sync.RWMutex
}

var (
	instance *TokenManager
	once     sync.Once
)

func GetManager() *TokenManager {
	once.Do(func() {
		instance = &TokenManager{}
		instance.RefreshToken()
	})
	return instance
}

func (tm *TokenManager) RefreshToken() {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	newToken := fmt.Sprintf("token_%d", time.Now().Unix())

	tm.accessToken = newToken
	tm.expiresAt = time.Now().Add(24 * time.Hour)
	fmt.Println("Token refreshed : ", tm.accessToken)
}

func (tm *TokenManager) GetToken() string {
	tm.mu.RLock()
	expired := time.Now().After(tm.expiresAt)
	tm.mu.RUnlock()

	if expired {
		tm.RefreshToken()
	}

	tm.mu.RLock()
	defer tm.mu.RUnlock()
	return tm.accessToken

}
