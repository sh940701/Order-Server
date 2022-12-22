package controller

import (
	"github.com/codestates/WBABEProject-08/commits/main/model"
	"github.com/gin-gonic/gin"
	"fmt"
)

type BuyerController struct {
	OrderedListModel *model.OrderedListModel
	MenuModel *model.MenuModel
}

func GetBuyerController(Om *model.OrderedListModel, Mm *model.MenuModel) *BuyerController {
	BuyerController := &BuyerController{OrderedListModel : Om, MenuModel : Mm}
	
	return BuyerController
}

// 메뉴 리스트를 조회하는 함수
func (bc *BuyerController) GetMenuList(c *gin.Context) {
	category := c.Param("category")
	fmt.Println("category: ", category)
	result := bc.MenuModel.GetList(category)

	c.JSON(200, gin.H{"result" : result})
}

// 메뉴별 평점/리뷰 데이터 조회하는 함수
func (bc *BuyerController) GetReview(c *gin.Context) {

}

// 현재 주문의 상태를 조회하는 함수
func (bc *BuyerController) GetOrderStatus(c *gin.Context) {

}

// 메뉴에 대한 평점 및 리뷰를 작성하는 함수
func (bc *BuyerController) AddReview(c *gin.Context) {

}

// 메뉴 선택 및 주문하는 함수
func (bc *BuyerController) Order(c *gin.Context) {

}

// 메뉴 변경하는 함수
func (bc *BuyerController) ChangeOrder(c *gin.Context) {

}

// 메뉴 추가하는 함수
func (bc *BuyerController) AddOrder(c *gin.Context) {

}