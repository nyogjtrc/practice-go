package imgur

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimplyCallAPI(t *testing.T) {
	ic, err := NewFromFile("../config.json")
	assert.NoError(t, err)

	_, err = ic.MyImageCount()
	assert.NoError(t, err)

	_, err = ic.GetImage("abcdef")
	assert.Error(t, err)

	ur, err := ic.UploadImage("../default.jpg")
	assert.NoError(t, err)
	t.Log(ur)

	dr, err := ic.DeleteImage(ur.Data.ID)
	assert.NoError(t, err)
	t.Log(dr)
}
