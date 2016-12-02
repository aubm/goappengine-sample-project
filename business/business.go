package business

import (
	"github.com/satori/go.uuid"
	"golang.org/x/net/context"
)

func GenerateUUID(ctx context.Context) string {
	return uuid.NewV4().String()
}
