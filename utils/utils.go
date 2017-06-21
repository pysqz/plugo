package utils

import (
	"context"
)

type Plugin interface {
	Do(ctx *context.Context) (retval uint)
}

var PluginSet map[string]Plugin

const (
	RETVAL_BREAK    uint = 0
	RETVAL_CONTINUE uint = 1
)

func init() {
	PluginSet = make(map[string]Plugin)
}

func Register(k string, v Plugin) {
	PluginSet[k] = v
}
