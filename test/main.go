package main

import (
	"fmt"
	"github.com/892294101/jxparams"
	"os"
)

func main() {
	p := jxparams.NewParams()
	/*	p.SetParams("jx.server.listener.port", jxparams.NewConfig().SetDefault("5100"))
		p.SetParams("jx.server.model")
		p.SetParams("jx.filesystem.image.path")
		p.SetParams("jx.filesystem.host")
		p.SetParams("jx.database.type")

		p.SetParams("jx.database.host")

		p.SetParams("jx.database.port")
		p.SetParams("jx.database.name")
		p.SetParams("jx.database.username")
		p.SetParams("jx.database.password")
		p.SetParams("jx.database.charset")
		p.SetParams("jx.database.connect.max.idle")
		p.SetParams("jx.database.connect.min.idle")
		p.SetParams("jx.database.connect.max")
		p.SetParams("jx.redis.address")
		p.SetParams("jx.redis.password")*/
	p.SetParams("auth", jxparams.NewConfig().SetPrefix())
	p.SetParams("appcode", jxparams.NewConfig().SetSuffix())
	p.SetParams("ksajfkasjdf", jxparams.NewConfig().SetSuffix().SetDefault("12312"))
	p.SetParams("port", jxparams.NewConfig().SetDefault("1521"))
	p.SetConfigFile("/tmp/test.conf")

	if err := p.Load(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	p.Println()
	fmt.Println("=====================================================")
	v, _ := p.GetPrefix("auth")
	for i, params := range v {
		fmt.Println(i, params.ToString())
	}
	fmt.Println("=====================================================")
	v, _ = p.GetSuffix("appcode")
	for i, params := range v {
		fmt.Println(i, params.ToString())
	}

}
