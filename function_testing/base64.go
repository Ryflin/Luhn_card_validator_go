package main

import "encoding/base64"

func main() {
	println(base64.StdEncoding.EncodeToString([]byte("")))
}
