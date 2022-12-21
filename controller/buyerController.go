package controller

import (
	"github.com/codestates/WBABEProject-08/commits/main/model"
)

type BuyerController struct {
	OrderedListModel *model.OrderedListModel
	MenuModel *model.MenuModel
}

func GetBuyerController(Om *model.OrderedListModel, Mm *model.MenuModel) *BuyerController {
	BuyerController := &BuyerController{OrderedListModel : Om, MenuModel : Mm}
	
	return BuyerController
}