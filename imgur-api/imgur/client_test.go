package imgur

import (
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
