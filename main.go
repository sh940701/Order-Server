package main

import (
	"time"
	"net/http"
	ctl "github.com/codestates/WBABEProject-08/commits/main/controller"
	model "github.com/codestates/WBABEProject-08/commits/main/model"
	route "github.com/codestates/WBABEProject-08/commits/main/route"
)

func errPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// model 객체 초기화
	mModel, err := model.GetMenuModel("order-system", "mongodb://localhost:27017", "menu")
	errPanic(err)
	oModel, err := model.GetOrderedListModel("order-system", "mongodb://localhost::27017", "orderedlist")
	errPanic(err)

	// model 객체를 넣어 controller를 만들어줌
	controller := ctl.GetController(oModel, mModel)

	// controller 객체를 넣어 router를 만들어줌
	router := route.NewRouter(controller)

	mapi := &http.Server {
		Addr : ":8080",
		Handler : router.Idx(),
		ReadTimeout : 5 * time.Second,
		WriteTimeout : 10 * time.Second,
		MaxHeaderBytes : 1 << 20,
	}
	
	mapi.ListenAndServe()
}