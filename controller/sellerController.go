package controller

import (
	"github.com/gin-gonic/gin"
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

// 주문내역 리스트 조회하는 함수
func (sc *SellerController) GetOrderList(c *gin.Context) {

}

// 주문 요청에 대한 상태 업데이트 함수
func (sc *SellerController) UpdateOrderStatus(c *gin.Context) {

}

// 새 메뉴를 추가하는 함수
func (sc *SellerController) AddMenu(c *gin.Context) {

}

// 메뉴를 삭제하는 함수
func (sc *SellerController) DeleteMenu(c *gin.Context) {

}

// 메뉴 정보를 수정하는 함수
func (sc *SellerController) UpdateMenu(c *gin.Context) {

}