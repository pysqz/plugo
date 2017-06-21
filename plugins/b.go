package plugins

import (
	"context"
	"log"

	"github.com/gooo000/plugo/utils"
)

type B struct {
}

func (this *B) Do(ctx *context.Context) (retval uint) {
	log.Printf("this is b.go")
	return utils.RETVAL_CONTINUE
}

func init() {
	utils.Register("b", &B{})
}
