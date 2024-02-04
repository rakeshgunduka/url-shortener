package utils

import (
	"os"
	"sync"
	"url-shortener-go/modules/tokenrange"
	"url-shortener-go/storage"
)

var tokenCounter int64
var startToken int64
var endToken int64

var counterMutex sync.Mutex

func InitTokenCounter() {
	TokenRangeService := &tokenrange.TokenRangeService{DB: storage.DB}
	tokenRange := TokenRangeService.GetNextRange(os.Getenv("SERVER_NAME"))

	startToken = tokenRange.Start
	endToken = tokenRange.End
	tokenCounter = startToken
}

func IncrementTokenCounter() {
	counterMutex.Lock()
	defer counterMutex.Unlock()

	tokenCounter++
}

func GetTokenCounter() int64 {
	counterMutex.Lock()
	defer counterMutex.Unlock()
	if tokenCounter == endToken {
		currentCounter := tokenCounter
		InitTokenCounter()
		return currentCounter
	}

	return tokenCounter
}
