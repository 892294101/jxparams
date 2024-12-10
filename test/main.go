package main

import (
	"fmt"
	"github.com/892294101/jxparams"
	"os"
)

const (
	ServerListenerHost     = "server.listener.host"
	ServerListenerPort     = "server.listener.port"
	ServerModel            = "server.model"
	ServerSSLCert          = "server.ssl.cert.path"
	ServerSSLKey           = "server.ssl.key.path"
	OSSApiAddress          = "oss.api.address"
	OSSBucket              = "oss.bucket"
	OSSApiSSL              = "oss.api.ssl"
	OSSAccessKey           = "oss.access.key"
	OSSSecretKey           = "oss.secret.key"
	StaticFS               = "static.file.system"
	DatabaseType           = "database.type"
	DatabaseUrl            = "database.url"
	DatabaseUsername       = "database.username"
	DatabasePassword       = "database.password"
	DatabaseQueryTimeOut   = "database.query.timeout"
	DatabaseTimeZone       = "database.timezone"
	DatabaseCharset        = "database.charset"
	DatabaseConnectMaxIdle = "database.connect.max.idle"
	DatabaseConnectMax     = "database.connect.max"
	RedisAddress           = "redis.address"
	RedisPassword          = "redis.password"
	RedisDatabase          = "redis.database"
	SentinelAddress        = "redis.sentinel.address"
	SentinelPassword       = "redis.sentinel.password"
	SentinelDatabase       = "redis.sentinel.database"
	SentinelMasterName     = "redis.sentinel.master.name"
	ClusterAddress         = "redis.cluster.address"
	ClusterPassword        = "redis.cluster.password"
	DBMySQLType            = "mysql"
	DBOracleType           = "oracle"
	DBPublic               = "public"
)

func main() {
	p := jxparams.NewParams()
	p.SetConfigFile("/Users/kqwang/development/gowork/src/github.com/892294101/jxapo/aposerver/build/conf/shop.conf")
	/*p.SetParams(ServerListenerHost, jxparams.NewConfig().SetDefault("0.0.0.0"))
	p.SetParams(ServerListenerPort, jxparams.NewConfig().SetDefault("7710"))
	p.SetParams(ServerModel, jxparams.NewConfig().SetDefault("info"))*/
	p.SetParams(ServerSSLCert)
	p.SetParams(ServerSSLKey)

	/*p.SetParams(OSSAccessKey, jxparams.NewConfig().SetMust())
	p.SetParams(OSSSecretKey, jxparams.NewConfig().SetMust())
	p.SetParams(OSSApiAddress, jxparams.NewConfig().SetMust())
	p.SetParams(OSSBucket, jxparams.NewConfig().SetMust())
	p.SetParams(OSSApiSSL, jxparams.NewConfig().SetDefault("true"))

	p.SetParams(StaticFS, jxparams.NewConfig().SetMust())

	p.SetParams(DatabaseType, jxparams.NewConfig().SetMust())
	p.SetParams(DatabaseUrl, jxparams.NewConfig().SetMust())
	p.SetParams(DatabaseUsername, jxparams.NewConfig().SetDefault("root"))
	p.SetParams(DatabasePassword, jxparams.NewConfig().SetMust())
	p.SetParams(DatabaseQueryTimeOut, jxparams.NewConfig().SetDefault("3"))
	p.SetParams(DatabaseTimeZone, jxparams.NewConfig().SetDefault("UTC"))
	p.SetParams(DatabaseCharset, jxparams.NewConfig().SetMust())
	p.SetParams(DatabaseConnectMaxIdle, jxparams.NewConfig().SetDefault("20"))
	p.SetParams(DatabaseConnectMax, jxparams.NewConfig().SetDefault("20"))

	p.SetParams(RedisAddress)
	p.SetParams(RedisPassword)
	p.SetParams(RedisDatabase, jxparams.NewConfig().SetDefault("0"))

	p.SetParams(SentinelAddress)
	p.SetParams(SentinelPassword)
	p.SetParams(SentinelDatabase, jxparams.NewConfig().SetDefault("0"))
	p.SetParams(SentinelMasterName)

	p.SetParams(ClusterAddress)
	p.SetParams(ClusterPassword)*/

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
