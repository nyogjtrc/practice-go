package imgur

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	ic := New("123", "abc", "aaaaa", "bbbbb")
	oc := ic.OAuth2Config()

	assert.Equal(t, ic.ClientID, oc.ClientID)
	assert.Equal(t, ic.ClientSecret, oc.ClientSecret)

	_, err := ic.HTTPClient()
	assert.NoError(t, err)
}

func TestParseError(t *testing.T) {
	ic := New("123", "abc", "aaaaa", "bbbbb")

	json := []byte(`{"data":{"error":"Unable to find an image","request":"","method":"GET"},"success":false,"status":404}`)
	err := ic.parseError(json)

	assert.Equal(t, errors.New("Unable to find an image"), err)
}
