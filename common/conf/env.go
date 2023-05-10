package conf

import (
	"os"
	"strings"
)

func ExpandEnv(s string) string {
	return os.Expand(s, getEnv)
}

func getEnv(key string) string {
	var k, def = key, ""
	if n := strings.Index(key, ":"); n > 0 {
		k = key[:n]
		def = key[n+1:]
	}
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
