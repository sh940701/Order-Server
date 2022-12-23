package controller

import (
	// "fmt"
	"github.com/codestates/WBABEProject-08/commits/main/util"
	"github.com/gin-gonic/gin"
	"github.com/codestates/WBABEProject-08/commits/main/model"
	"io"
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
	// 전체 주문내역 리스트를 가져온다.
	result := sc.OrderedListModel.GetAll(daycountId)

	c.JSON(200, gin.H{"주문 목록" : result})
}

// 주문 요청에 대한 상태 업데이트 함수
func (sc *SellerController) UpdateOrderStatus(c *gin.Context) {
	orderId := util.ConvertStringToObjectId(c.Param("orderid"))
	order := sc.OrderedListModel.GetOne(orderId)
	
	msg := sc.OrderedListModel.UpdateStatus(order)
	c.JSON(200, gin.H{"msg" : msg})
}

// 새 메뉴를 추가하는 함수
func (sc *SellerController) AddMenu(c *gin.Context) {
	body := c.Request.Body
	byteData, err := io.ReadAll(body)
	util.PanicHandler(err)

	result, err := sc.MenuModel.Add(byteData)
	if err != nil {
		c.JSON(404, gin.H{"msg" : "실패하였습니다.", "err" : err})
	} else {
		c.JSON(200, gin.H{"msg" : "성공하였습니다.", "id" : result})
	}
}

// 메뉴를 삭제하는 함수
func (sc *SellerController) DeleteMenu(c *gin.Context) {
	body := c.Request.Body
	byteDate, err := io.ReadAll(body)
	util.PanicHandler(err)

	result, err := sc.MenuModel.Delete(byteDate)
	if err != nil {
		c.JSON(404, gin.H{"msg" : "실패하였습니다.", "err" : err})
	} else if result == nil {
		c.JSON(404, gin.H{"msg" : "실패하였습니다."})
	} else {
		c.JSON(200, gin.H{"msg" : "성공하였습니다."})
	}
}

// 메뉴 정보를 수정하는 함수
func (sc *SellerController) UpdateMenu(c *gin.Context) {
	body := c.Request.Body
	data, err := io.ReadAll(body)
	util.PanicHandler(err)

	result, err := sc.MenuModel.Update(data)

	if err != nil {
		c.JSON(404, gin.H{"msg" : "실패하였습니다.", "err" : err})
	} else if result == nil {
		c.JSON(404, gin.H{"msg" : "실패하였습니다."})
	} else {
		c.JSON(200, gin.H{"msg" : "성공하였습니다."})
	}

}