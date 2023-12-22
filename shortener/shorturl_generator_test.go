package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const UserId = "0dba740-fc4b-4977-872c-d360239e6b1a"

func TestShortLinkGenerator(t *testing.T) {
	initialLink_1 := "https://github.com/ThePrimeagen/harpoon/tree/harpoon2"
	shortLink_1 := GenerateShortLink(initialLink_1, UserId)

	initialLink_2 := "https://interpreterbook.com/"
	shortLink_2 := GenerateShortLink(initialLink_2, UserId)

	initialLink_3 := "http://www.textfiles.com/index.html"
	shortLink_3 := GenerateShortLink(initialLink_3, UserId)

	assert.Equal(t, shortLink_1, "W3Sks5txwfv")
	assert.Equal(t, shortLink_2, "RFh6yzrpued")
	assert.Equal(t, shortLink_3, "9TeFJ6YFXfn")
}
