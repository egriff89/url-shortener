package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testStoreService = &StorageService{}

func init() {
	testStoreService = InitializeStore()
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}

func TestInsertionAndRetrieval(t *testing.T) {
	initialLink := "https://github.com/ThePrimeagen/harpoon/tree/harpoon2"
	userUUID := "e0dba740-fc4b-4977-872c-d360239e6b1a"
	shortURL := "Jsz4k57oAX"

	SaveUrlMapping(shortURL, initialLink, userUUID)
	retrievedUrl := RetrieveInitialUrl(shortURL)

	assert.Equal(t, initialLink, retrievedUrl)
}
