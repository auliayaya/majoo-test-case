package api

import (
	"aulia-majoo-test/config"
	"aulia-majoo-test/routes"
	"fmt"
)

var conf *config.APP

func init() {
	conf = config.AppConfig
}

func StartApplication() {
	// PORT For DEV 7610
	routes.Routes().Run(fmt.Sprintf(":%s", conf.Config.GetString("APP_LOCAL_DEV_PORT")))
}
