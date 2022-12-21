package controller

import (
	"github.com/codestates/WBABEProject-08/commits/main/model"
)

type Controller struct {
	buyer *BuyerController
	seller *SellerController
}

// 이건 메인에서 실행해서 route에 넣어줄거다.
func GetController(Om *model.OrderedListModel, Mm *model.MenuModel) *Controller {
	controller := &Controller{buyer : GetBuyerController(Om, Mm), seller : GetSellerController(Om, Mm)}

	return controller
}

// controller 객체 내의 buyer와 seller가 숨겨져있으므로, 함수를 통해서 사용할 수 있게 해주자
func (c *Controller) GetBuyer() *BuyerController {
	return c.buyer
}

func (c *Controller) GetSeller() *SellerController {
	return c.seller
}