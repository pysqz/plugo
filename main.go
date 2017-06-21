package main

import (
	"context"
	"log"
	"os"

	"github.com/BurntSushi/toml"

	"github.com/gooo000/plugo/config"
	_ "github.com/gooo000/plugo/plugins"
	"github.com/gooo000/plugo/utils"
)

func main() {
	if _, err := toml.DecodeFile("conf.toml", config.Conf); err != nil {
		log.Printf("failed to decode conf.toml, err: [%s]", err)
		os.Exit(1)
	}

	ctx := context.WithValue(context.Background(), "test_key", "test_value")
	for _, i := range config.Conf.Plugins {
		if utils.PluginSet[i].Do(&ctx) == utils.RETVAL_BREAK {
			break
		}
	}
}
