package main

import (
	"fmt"
	"github.com/892294101/jxparams"
	"os"
)

func main() {
	p := jxparams.NewParams()
	p.SetParams("jx.server.listener.port", jxparams.NewConfig().SetDefault("5100"))
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
	p.SetParams("jx.redis.password")

	p.SetConfigFile("/Users/kqwang/development/gowork/src/github.com/892294101/jxshop/server/build/conf/shop.conf")

	if err := p.Load(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if v, ok := p.GetParams("jx.database.port"); ok {
		fmt.Println(v.ToString())
		fmt.Println(v.ToInt64())
		fmt.Println(v.ToInt())
		fmt.Println(v.ToFloat64())
	}
	p.Println()
}
