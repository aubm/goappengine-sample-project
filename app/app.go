package app

import (
	"net/http"

	"github.com/aubm/goappengine-sample-project/api"
)

func init() {
	http.HandleFunc("/", api.NewUUIDHandler)
}
