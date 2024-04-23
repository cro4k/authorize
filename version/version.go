package version

import (
	"encoding/json"
	"net/http"
	"time"
)

const (
	Version = "v0.0.1"
)

var (
	BuildTime = "-"
	BuildHash = "-"
)

func HTTPHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{
			"version":    Version,
			"build_time": BuildTime,
			"build_hash": BuildHash,
			"timestamp":  time.Now().Format(time.DateTime),
		})
	}
}
