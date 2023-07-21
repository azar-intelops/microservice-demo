package client

import (
	"io"
	"net/http"
)

func ExecuteNode61() ([]byte, error) {
	response, err := http.Get("one2one:9800/ping")

	if err != nil {
		return nil, err
	}

	return io.ReadAll(response.Body)
}
