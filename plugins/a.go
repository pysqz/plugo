package plugins

import (
	"context"
	"log"

	"github.com/gooo000/plugo/utils"
)

type A struct {
}

func (this *A) Do(ctx *context.Context) (retval uint) {
	log.Printf("this is a.go")
	return utils.RETVAL_CONTINUE
}

func init() {
	utils.Register("a", &A{})
}
