package main

import (
	"fmt"
	"time"
	"net/http"
	ctl "github.com/codestates/WBABEProject-08/commits/main/controller"
	logger "github.com/codestates/WBABEProject-08/commits/main/log"
	model "github.com/codestates/WBABEProject-08/commits/main/model"
	route "github.com/codestates/WBABEProject-08/commits/main/route"
	conf "github.com/codestates/WBABEProject-08/commits/main/config"
)

func errPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var err error
	
	config := conf.GetConfig("/Users/sunghyun/Desktop/projects/wemade_project/config/config.toml")
	
		if err := logger.InitLogger(config); err != nil {
			fmt.Printf("init logger failed, err:%v\n", err)
			return
		}
	
		logger.Debug("ready server....")

	port := config.Server.Port
	host := config.Server.Host
	dbName := config.Server.DBname
	Model0 := config.DB["menu"]["model"]
	Model1 := config.DB["orderedlist"]["model"]

	// model 객체 초기화
	mModel, err := model.GetMenuModel(dbName, host, Model0)
	errPanic(err)
	oModel, err := model.GetOrderedListModel(dbName, host, Model1)
	errPanic(err)

	// model 객체를 넣어 controller를 만들어줌
	controller := ctl.GetController(oModel, mModel)

	// controller 객체를 넣어 router를 만들어줌
	router := route.NewRouter(controller)

	mapi := &http.Server {
		Addr : port,
		Handler : router.Idx(),
		ReadTimeout : 5 * time.Second,
		WriteTimeout : 10 * time.Second,
		MaxHeaderBytes : 1 << 20,
	}
	
	mapi.ListenAndServe()
}