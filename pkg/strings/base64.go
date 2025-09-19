package strings

import "encoding/base64"

func Base64Encode(value string) string {
	data := []byte(value)
	str := base64.StdEncoding.EncodeToString(data)
	return str
}

func Base64Decode(value string) (string, error) {
	data, _ := base64.StdEncoding.DecodeString(value)
	return string(data), nil
}

func EncodeURL(value string) string {
	return base64.RawURLEncoding.EncodeToString([]byte(value))
}

func DecodeURL(value string) (string, error) {
	data, _ := base64.RawURLEncoding.DecodeString(value)
	return string(data), nil
}
