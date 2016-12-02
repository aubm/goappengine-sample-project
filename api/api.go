package api

import (
	"fmt"
	"net/http"

	"github.com/aubm/goappengine-sample-project/business"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func NewUUIDHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	log.Infof(ctx, "Starting the generation of a new uuid")
	id := business.GenerateUUID(ctx)
	log.Infof(ctx, "New uuid generated: %s", id)

	fmt.Fprintf(w, "Hello, here is your uuid: %s", id)
}
