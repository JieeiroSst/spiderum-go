package utils

import "crypto/rand"

type EmailUtils struct {}

func NewEmailUtils()*EmailUtils {
	return &EmailUtils{}
}

func (e *EmailUtils) ChunkSplit(body string, limit int, end string) string {
	var charSlice []rune
	for _, char := range body {
		charSlice = append(charSlice, char)
	}
	var result string = ""
	for len(charSlice) >= 1 {
		result = result + string(charSlice[:limit]) + end
		charSlice = charSlice[limit:]
		if len(charSlice) < limit {
			limit = len(charSlice)
		}

	}
	return result
}

func (e *EmailUtils) RandStr(strSize int, randType string) string {
	var dictionary string
	if randType == "alphanum" {
		dictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}
	if randType == "alpha" {
		dictionary = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}
	if randType == "number" {
		dictionary = "0123456789"
	}
	var bytes = make([]byte, strSize)
	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	return string(bytes)
}