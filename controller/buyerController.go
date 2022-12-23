package controller

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/codestates/WBABEProject-08/commits/main/util"
	"github.com/codestates/WBABEProject-08/commits/main/model"
	"github.com/gin-gonic/gin"
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
	if (category != "avg") && (category != "suggestion") && (category != "orderedcount") {
		c.JSON(404, gin.H{"error" : "해당 카테고리가 존재하지 않습니다."})
		return
	}
	result := bc.MenuModel.GetList(category)

	c.JSON(200, gin.H{"result" : result})
}


// 메뉴별 평점/리뷰 데이터 조회하는 함수
func (bc *BuyerController) GetReview(c *gin.Context) {
	id := util.ConvertStringToObjectId(c.Param("menuid"))
	menu, err := bc.MenuModel.GetOneMenu(id)
	if err != nil {
		c.JSON(404, gin.H{"error" : err})
		return
	} else if len(menu.Reviews) == 0 {
		c.JSON(400, gin.H{"msg" : "현재 리뷰가 없습니다."})
		return
	} 
	c.JSON(200, gin.H{"평점" : menu.Avg, "리뷰" : menu.Reviews})
}


// 현재 주문의 상태를 조회하는 함수
func (bc *BuyerController) GetOrderStatus(c *gin.Context) {
	orderId := util.ConvertStringToObjectId(c.Param("orderid"))
	order := bc.OrderedListModel.GetOne(orderId)
	// 존재하지 않는 id를 입력했을때의 처리
	if len(order.Orderedmenus) == 0 {
		c.JSON(400, gin.H{"msg" : "존재하지 않는 주문입니다."})
		return
	} else {
		c.JSON(200, gin.H{"현재 주문 상태" : order.Status})
	}
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
		c.JSON(400, gin.H{"msg" : "이미 후기가 작성된 주문입니다."})
		return
	}

	// 음식에 리뷰를 추가한다.
	foodId := c.Param("menuid")
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
		c.JSON(400, gin.H{"msg" : "현재 주문이 불가능합니다."})
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
	changeMenuStruct := &model.ChangeMenuType{}
	err := c.ShouldBindJSON(changeMenuStruct)
	util.PanicHandler(err)

	order := bc.OrderedListModel.GetOne(changeMenuStruct.OrderId)
	if (order.Status == "배달중") || (order.Status == "배달완료") {
		c.JSON(400, gin.H{"msg" : "조리가 완료되어 메뉴 변경이 불가능합니다."})
	} else {
		err := bc.OrderedListModel.ChangeOrder(order, changeMenuStruct)
		if err != nil {
			c.JSON(400, gin.H{"error" : err.Error()})
		} else {
			c.JSON(200, gin.H{"msg" : "메뉴 변경이 완료되었습니다."})
		}
	}
}


// 메뉴 추가하는 함수
func (bc *BuyerController) AddOrder(c *gin.Context) {
	var msg string
	var id primitive.ObjectID

	addStruct := model.AddMenuType{}
	err := c.ShouldBindJSON(&addStruct)
	util.PanicHandler(err)

	// 해당 주문의 현재 상태를 검사한다.
	order := bc.OrderedListModel.GetOne(addStruct.OrderId)
	if (order.Status == "배달중") || (order.Status == "배달완료") {
		id = bc.OrderedListModel.Add(&addStruct.NewOrder)
		msg = "주문 추가가 불가능하여 새 주문으로 접수되었습니다."
	} else {
		id = bc.OrderedListModel.AddOrder(&addStruct, order)
		msg = "주문 추가가 정상적으로 완료되었습니다."
	}
	c.JSON(200, gin.H{"msg" : msg, "id" : id})
}


// 메뉴목록 전체를 가져오는 함수
func (bc *BuyerController) GetAll(c *gin.Context) {
	list := bc.MenuModel.GetAllMenu()

	c.JSON(200, gin.H{"menus" : list})
}