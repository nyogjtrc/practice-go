package imgur

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseError(t *testing.T) {
	json := []byte(`{"data":{"error":"Unable to find an image","request":"","method":"GET"},"success":false,"status":404}`)
	err := parseError(json)

	assert.Equal(t, errors.New("Unable to find an image"), err)
}
