package mysqlproto

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNativePassword(t *testing.T) {
	data := "abcdefghijklmnopqrst"
	pass := "user123"

	hash := nativePassword(data, pass)
	assert.Len(t, hash, 20)
	assert.Equal(t, hex.EncodeToString(hash), "d7dbc13284c74850c777657fc3d7eb80b7185a25")
}
