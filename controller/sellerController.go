package controller

import (
	"github.com/codestates/WBABEProject-08/commits/main/model"
)

type SellerController struct {
	OrderedListModel *model.OrderedListModel
	MenuModel *model.MenuModel
}

func GetSellerController(Om *model.OrderedListModel, Mm *model.MenuModel) *SellerController {
	SellerController := &SellerController{OrderedListModel : Om, MenuModel : Mm}
	
	return SellerController
}