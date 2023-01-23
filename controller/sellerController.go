package controller

import (
	// "fmt"
	"github.com/codestates/WBABEProject-08/commits/main/util"
	"github.com/gin-gonic/gin"
	"github.com/codestates/WBABEProject-08/commits/main/model"
	"io"
	"strconv"
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
// GetOrderList godoc
// @Summary 주문 전체 내역을 page에 맞춰서 가져옵니다. 각 페이지당 5개의 data를 포함합니다.
// @Tags Seller
// @Description 주문 내역을 가져오기 위한 함수
// @name GetOrderList
// @Accept  json
// @Produce  json
// @Param page query string true "page"
// @Router /seller/order [get]
// @Success 200 {array} model.Menu
func (sc *SellerController) GetOrderList(c *gin.Context) {
	sPage := c.Query("page")
	iPage, _ := strconv.Atoi(sPage)
	// 전체 주문내역 리스트를 가져온다.
	result := sc.OrderedListModel.GetAll(daycountID, int64(iPage))

	c.JSON(200, gin.H{"주문 목록" : result})
}


// 주문 요청에 대한 상태 업데이트 함수
// UpdateOrderStatus godoc
// @Summary 현재 주문의 상태를 업데이트 합니다. 상태는 "주문접수 -> 조리중 -> 배달중 -> 배달완료" 의 순서로 진행되고, 상태가 배달중 이상인 경우 메뉴 변경 및 추가 등의 작업이 제한됩니다.
// @Tags Seller
// @Description 주문의 상태를 업데이트 하기 위한 함수
// @name UpdateOrderStatus
// @Accept  json
// @Produce  json
// @Param id body model.IdType true "id"
// @Router /seller/order [patch]
// @Success 200 {object} string
func (sc *SellerController) UpdateOrderStatus(c *gin.Context) {
	body := c.Request.Body
	data, err := io.ReadAll(body)
	util.ErrorHandler(err)
	orderId, _, _ := util.GetJsonIdKeyValue(data)
	order := sc.OrderedListModel.GetOne(orderId)
	
	msg := sc.OrderedListModel.UpdateStatus(order)
	c.JSON(200, gin.H{"msg" : msg})
}


// 새 메뉴를 추가하는 함수
// AddMenu godoc
// @Summary DB에 새 메뉴를 추가합니다.
// @Tags Seller
// @Description 새 메뉴를 추가하기 위한 함수
// @name AddMenu
// @Accept  json
// @Produce  json
// @Param menu body model.AddMenuDataType true "menu"
// @Router /seller/menu [post]
// @Success 200 {object} string
// @failure 404 {object} string
func (sc *SellerController) AddMenu(c *gin.Context) {
	body := c.Request.Body
	byteData, err := io.ReadAll(body)
	util.ErrorHandler(err)

	result, err := sc.MenuModel.AddMenu(byteData)
	if err != nil {
		c.JSON(404, gin.H{"msg" : "실패하였습니다.", "err" : err})
	} else {
		c.JSON(200, gin.H{"msg" : "성공하였습니다.", "id" : result})
	}
}


// 메뉴를 삭제하는 함수
// DeleteMenu godoc
// @Summary DB에서 메뉴를 삭제합니다.
// @Tags Seller
// @Description 메뉴를 삭제하기 위한 함수
// @name DeleteMenu
// @Accept  json
// @Produce  json
// @Param menuid query string true "menuid"
// @Router /seller/menu [delete]
// @Success 200 {object} string
// @failure 404 {object} string
func (sc *SellerController) DeleteMenu(c *gin.Context) {
	id := c.Query("menuid")
	menuId := util.ConvertStringToObjectId(id)

	result, err := sc.MenuModel.DeleteMenu(menuId)
	if err != nil {
		c.JSON(404, gin.H{"msg" : "실패하였습니다.", "err" : err})
	} else if result == nil {
		c.JSON(404, gin.H{"msg" : "실패하였습니다."})
	} else {
		c.JSON(200, gin.H{"msg" : "성공하였습니다."})
	}
}


// 메뉴 정보를 수정하는 함수
// UpdateMenu godoc
// @Summary 메뉴 정보를 업데이트 합니다. ex) 가격 : 10000 -> 가격 : 9000
// @Tags Seller
// @Description 메뉴를 수정하기 위한 함수
// @name UpdateMenu
// @Accept  json
// @Produce  json
// @Param updateData body model.UpdateMenuDataType true "update data"
// @Router /seller/menu [patch]
// @Success 200 {object} string
// @failure 404 {object} string
func (sc *SellerController) UpdateMenu(c *gin.Context) {
	body := c.Request.Body
	data, err := io.ReadAll(body)
	util.ErrorHandler(err)

	result, err := sc.MenuModel.UpdateMenu(data)

	if err != nil {
		c.JSON(404, gin.H{"msg" : "실패하였습니다.", "err" : err})
	} else if result == nil {
		c.JSON(404, gin.H{"msg" : "실패하였습니다."})
	} else {
		c.JSON(200, gin.H{"msg" : "성공하였습니다."})
	}
}


// 추천메뉴를 변경하는 함수
// SuggestMenu godoc
// @Summary 추천 메뉴를 업데이트 합니다. ex) 햄버거, 삼겹살 -> 스테이크, 컵밥
// @Tags Seller
// @Description 추천 메뉴를 업데이트 하기 위한 함수
// @name SuggestMenu
// @Accept  json
// @Produce  json
// @Param suggestionIds body model.SuggestionType true "suggest data"
// @Router /seller/suggestion [patch]
// @Success 200 {object} string
// @failure 404 {object} string
func (sc *SellerController) SuggestMenu(c *gin.Context) {
	var suggestion model.SuggestionType
	err := c.ShouldBindJSON(&suggestion)
	util.ErrorHandler(err)

	err = sc.MenuModel.SuggestionUpdate(&suggestion)
	if err != nil {
		c.JSON(400, gin.H{"error" : err.Error()})
		return
	}
	c.JSON(200, gin.H{"msg" : "추천메뉴 업데이트가 완료되었습니다."})
}
