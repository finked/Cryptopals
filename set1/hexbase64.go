package main

import "encoding/hex"
import "encoding/base64"

func HexBase64(data string) string {

	// Transform hex code into bytes
	bytedata, _ := hex.DecodeString(data)

	sEnc := base64.StdEncoding.EncodeToString(bytedata)

	return sEnc
}
