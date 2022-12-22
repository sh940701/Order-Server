package controller

import (
	"github.com/codestates/WBABEProject-08/commits/main/util"
	"github.com/codestates/WBABEProject-08/commits/main/model"
	"github.com/gin-gonic/gin"
	"fmt"
)

const daycountId = "63a488f6165051c3589bcbe4"

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
	review := &model.Review{}
	err := c.ShouldBindJSON(review)
	if err != nil {
		util.PanicHandler(err)
	}
	// fmt.Println(review)
	// 먼저 해당 주문에 대한 리뷰가 있는지 확인하고 있다면 멈춘다.
	orderId := review.OrderId
	order := bc.OrderedListModel.GetOne(orderId)
	if order.IsReviewed {
		c.JSON(200, gin.H{"msg" : "이미 후기가 작성된 주문입니다."})
		return
	}

	// 음식에 리뷰를 추가한다.
	foodId := c.Param("foodid")
	bc.MenuModel.AddReview(foodId, review)

	// 주문의 리뷰 작성 여부를 업데이트한다.
	bc.OrderedListModel.UpdateReviewable(review.OrderId)

	// 음식의 평균 점수를 계산하여 넣어준다.
	bc.MenuModel.CalcAvg(foodId)

	c.JSON(200, gin.H{"msg" : "리뷰 작성이 완료되었습니다."})
}

// 메뉴 선택 및 주문하는 함수
func (bc *BuyerController) Order(c *gin.Context) {
	list := &model.OrderedList{}
	err := c.ShouldBindJSON(list)
	util.PanicHandler(err)

	// 주문하기 전에 메뉴가 주문 가능한지 확인
	menu, err := bc.MenuModel.GetOneMenu(list.Orderedmenus[0])
	if err != nil {
		c.JSON(404, gin.H{"error" : err})
		return
	} else if !menu.Orderable {
		c.JSON(200, gin.H{"msg" : "현재 주문이 불가능합니다."})
		return
	} else {
		// 있다면 limit -1, count +1
		bc.MenuModel.LimitAndCountUpdate(menu.ID, menu.Limit, menu.Orderedcount)
		// 그리고 DayOrderCount 1 추가
		dayCount := bc.OrderedListModel.DayOrderCount(daycountId)
		// 주문 추가
		orderNum := bc.OrderedListModel.Add(list)

		c.JSON(200, gin.H{"주문아이디" : orderNum, "오늘 주문번호" : dayCount})
	}
	

}

// 메뉴 변경하는 함수
func (bc *BuyerController) ChangeOrder(c *gin.Context) {

}

// 메뉴 추가하는 함수
func (bc *BuyerController) AddOrder(c *gin.Context) {

}