package configor

import (
	"fmt"
	"github.com/jinzhu/configor"
	"github.com/mix-go/xcli/argv"
)

var Config = struct {
	Mock []struct {
		Method   string
		Path     string
		Response string
	}
}{}

func init() {
	// Conf support YAML, JSON, TOML, Shell Environment
	if err := configor.Load(&Config, fmt.Sprintf("%s/../conf/config.yml", argv.Program().Dir)); err != nil {
		panic(err)
	}
}
