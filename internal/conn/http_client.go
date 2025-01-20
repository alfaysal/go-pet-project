package conn

import (
	"net/http"
	"time"
)

var defaultClient *http.Client

func InitHttpClient() {
	defaultClient = &http.Client{
		Timeout: time.Second * 5,
	}
}

func GetHttpClient() *http.Client {
	return defaultClient
}
