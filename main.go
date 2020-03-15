package main

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/file"
	"github.com/micro/go-micro/v2/web"
	conf "github.com/wumuwumu/trace-cloud/config"
	"github.com/wumuwumu/trace-cloud/http"
	"github.com/wumuwumu/trace-cloud/storage"
)

func main(){
	err :=config.Load(
		// base config from env
		//env.NewSource(),
		// override env with flags
		//flag.NewSource(),
		// override flags with file
		file.NewSource(
			file.WithPath("./trace-cloud.toml"),
		),
		//file.NewSource(
		//	file.WithPath("/etc/trace-cloud/config.toml"),
		//	),
		)
	if err != nil{
		panic(err)
	}
	if err = config.Scan(&conf.C);err != nil{
		panic(err)
	}

	storage.Init(conf.C)
	service := web.NewService(
		web.Name("sciento.cn.farm.trace"),
		web.Address(":8055"))
	err =service.Init()
	if err != nil{
		panic(err)
	}
	http.Init(service)
	if err := service.Run(); err != nil {
		panic(err)
	}

}
