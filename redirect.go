package main


import (
	"flag"
	"os"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"github.com/VinkDong/redirect/types"
	"github.com/VinkDong/redirect/server"
)

var (
	key       = flag.String("tsl_key", "", "TSL Key")
	cert      = flag.String("tsl_cert", "", "TSL Cert")
	enableSSL = flag.Bool("ssl", false, "Enable SSL")
	config    = flag.String("conf", "", "Config files")
)

func main() {
	flag.Parse()
	readConfig()
}

func readConfig() {
	if *config == "" {
		os.Exit(128)
	}
	data, err := ioutil.ReadFile(*config)

	if err != nil {
		panic(err)
	}

	config := &types.Config{}

	err = yaml.Unmarshal(data, config)
	if err != nil {
		panic(err)
	}

	initConfig(config)

	ctx := server.Context
	ctx.Config = config
	server.Context = ctx
	server.Run()
}

func initConfig(config *types.Config) {
	config.EnableSSL = *enableSSL
	config.Cert = *cert
	config.Key = *key
}