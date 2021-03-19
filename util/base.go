package util

import (
	"encoding/base32"
	"encoding/base64"
	"strings"
)

func HandleBase32(command string) string {
	words := strings.TrimPrefix(command, "!base32 ")
	encoded := base32.StdEncoding.EncodeToString([]byte(words))
	return encoded
}

func HandleBase64(command string) string {
	words := strings.TrimPrefix(command, "!base64 ")
	encoded := base64.StdEncoding.EncodeToString([]byte(words))
	return encoded
}
