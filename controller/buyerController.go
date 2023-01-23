package controller

import (
	"strconv"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/codestates/WBABEProject-08/commits/main/util"
	"github.com/codestates/WBABEProject-08/commits/main/model"
	"github.com/gin-gonic/gin"
)

const daycountID = "63a488f6165051c3589bcbe4"

type BuyerController struct {
	OrderedListModel *model.OrderedListModel
	MenuModel *model.MenuModel
}

func GetBuyerController(Om *model.OrderedListModel, Mm *model.MenuModel) *BuyerController {
	BuyerController := &BuyerController{OrderedListModel : Om, MenuModel : Mm}
	
	return BuyerController
}


// 조건에 맞춰 메뉴 리스트를 조회하는 함수
// GetMenuList godoc
// @Summary 조건에 맞는 메뉴들을 정렬하여 가져옵니다. 각 카테고리는 추천(suggestion), 평점(avg), 주문 수(orderedcount)로 이루어져 있으며, 페이지별 5개의 data를 반환합니다.
// @Tags Buyer
// @Description 카테고리별 메뉴 리스트를 정렬하여 가져오기 위한 함수
// @name GetMenuList
// @Accept  json
// @Produce  json
// @Param category path string true "category"
// @Param page query string true "page"
// @Router /buyer/menu/{category} [get]
// @Success 200 {object} string
func (bc *BuyerController) GetMenuList(c *gin.Context) {
	category := c.Param("category")
	sPage := c.Query("page")
	iPage, _ := strconv.Atoi(sPage)
	if (category != "avg") && (category != "suggestion") && (category != "orderedcount") {
		c.JSON(404, gin.H{"error" : "해당 카테고리가 존재하지 않습니다."})
		return
	}
	result := bc.MenuModel.GetMenuList(category, int64(iPage))

	c.JSON(200, gin.H{"result" : result})
}


// 메뉴별 평점/리뷰 데이터 조회하는 함수
// GetReview godoc
// @Summary 메뉴별 리뷰와 평점을 가져옵니다.
// @Tags Buyer
// @Description 메뉴별로 리뷰와 평점을 가져오기 위한 함수
// @name GetReview
// @Accept  json
// @Produce  json
// @Param menu_id path string true "menuid"
// @Router /buyer/review/{menu_id} [get]
// @Success 200 {object} string
// @failure 400 {object} string
// @failure 404 {object} string
func (bc *BuyerController) GetReview(c *gin.Context) {
	id := util.ConvertStringToObjectId(c.Param("menuid"))
	menu, err := bc.MenuModel.GetOneMenu(id)
	if err != nil {
		c.JSON(404, gin.H{"error" : err.Error()})
		return
	} else if len(menu.Reviews) == 0 {
		c.JSON(400, gin.H{"msg" : "현재 리뷰가 없습니다."})
		return
	} 
	c.JSON(200, gin.H{"평점" : menu.Avg, "리뷰" : menu.Reviews})
}


// 현재 주문의 상태를 조회하는 함수
// GetOrderStatus godoc
// @Summary 현재 주문의 진행상황을 확인합니다.
// @Tags Buyer
// @Description 현재 주문의 진행상황을 확인하기 위한 함수
// @name GetOrderStatus
// @Accept  json
// @Produce  json
// @Param order_id path string true "orderId"
// @Router /buyer/order/{order_id} [get]
// @Success 200 {object} string
// @failure 400 {object} string
func (bc *BuyerController) GetOrderStatus(c *gin.Context) {
	orderId := util.ConvertStringToObjectId(c.Param("orderid"))
	order := bc.OrderedListModel.GetOne(orderId)
	// 존재하지 않는 id를 입력했을때의 처리
	if len(order.OrderedMenus) == 0 {
		c.JSON(400, gin.H{"msg" : "존재하지 않는 주문입니다."})
		return
	} else {
		c.JSON(200, gin.H{"현재 주문 상태" : order.Status})
	}
}


// 메뉴에 대한 평점 및 리뷰를 작성하는 함수
// AddReview godoc
// @Summary 메뉴별 리뷰와 평점을 작성합니다.
// @Tags Buyer
// @Description 메뉴별로 리뷰와 평점 작성하기 위한 함수
// @name AddReview
// @Accept  json
// @Produce  json
// @Param menu_id path string true "menuid"
// @Param review body model.Review true "review"
// @Router /buyer/review/{menu_id} [post]
// @Success 200 {object} string
// @failure 400 {object} string
func (bc *BuyerController) AddReview(c *gin.Context) {
	review := &model.Review{}
	err := c.ShouldBindJSON(review)
	if err != nil {
		util.ErrorHandler(err)
	}
	// 평점이 5점 이상이라면 abort한다.
	if review.Score > 5 {
		c.JSON(400, gin.H{"msg" : "평점은 5점을 넘을 수 없습니다."})
		return
	}
	foodId := c.Param("menuid")
	// 먼저 해당 주문에 대한 리뷰 완료되었는지 확인한다.
	orderId := review.OrderId
	order := bc.OrderedListModel.GetOne(orderId)
	if order.IsReviewed {
		c.JSON(400, gin.H{"msg" : "이미 후기작성이 완료된 주문입니다."})
		return
	}
	// 그리고 해당 주문의 해당 메뉴에 대해서 리뷰가 작성되었는지 확인한다.
	for _, value := range order.OrderedMenus {
		if value.MenuId == util.ConvertStringToObjectId(foodId) && value.IsReviewed {
			c.JSON(400, gin.H{"msg" : "이미 해당 메뉴에 대한 후기가 작성되었습니다."})
			return
		}
	}

	// 음식에 리뷰를 추가한다.
	bc.MenuModel.AddReview(foodId, review)

	// 주문내에서 해당 "메뉴의 리뷰 작성 여부"를 업데이트한다.
	isOrderReviewed := bc.OrderedListModel.UpdateMenuReviewable(order, util.ConvertStringToObjectId(foodId))

	// 만약 모든 메뉴의 IsReviewed가 true라면, "주문의 전체 리뷰 여부"도 업데이트 해준다.
	if isOrderReviewed {
		bc.OrderedListModel.UpdateOrderReviewable(review.OrderId)
	}

	// 음식의 평균 점수를 계산하여 넣어준다.
	bc.MenuModel.CalcAvg(foodId)

	c.JSON(200, gin.H{"msg" : "메뉴에 대한 리뷰 작성이 완료되었습니다."})
}


// 메뉴 선택 및 주문하는 함수
// Order godoc
// @Summary 주문을 진행합니다.
// @Tags Buyer
// @Description 주문을 하는 함수
// @name Order
// @Accept  json
// @Produce  json
// @Param orderData body model.AddOrderType true "orderData"
// @Router /buyer/order [post]
// @Success 200 {object} string
// @failure 400 {object} string
// @failure 404 {object} string
func (bc *BuyerController) Order(c *gin.Context) {
	list := &model.OrderedList{
		IsReviewed : false,
		Status : "주문접수",
	}
	err := c.ShouldBindJSON(list)
	util.ErrorHandler(err)

	menus := []model.Menu{}

	// 모든 메뉴는 주문 시점에 리뷰가 작성되어있지 않기 때문에 false로 초기화해줌
	for i := 0; i < len(list.OrderedMenus); i++ {
		list.OrderedMenus[i].IsReviewed = false
	}

	// 주문하기 전에 메뉴가 주문 가능한지 확인
	// 메뉴 리스트를 돌면서 모든 메뉴가 주문 가능한 상태인지 확인한다.
	for i := 0; i < len(list.OrderedMenus); i++ {
		menu, err := bc.MenuModel.GetOneMenu(list.OrderedMenus[i].MenuId)
		if err != nil {
			c.JSON(404, gin.H{"error" : err})
			return
		} else if !menu.IsOrderable {
			c.JSON(400, gin.H{"해당 메뉴는 현재 주문이 불가능합니다." : menu.Name})
			return
		// 주문 가능한 메뉴 수량 체크 로직 추가
		} else if menu.Limit < list.OrderedMenus[i].Amount {
			c.JSON(400, gin.H{"msg" : "해당 메뉴의 수량이 부족합니다.", "남은 수량" : menu.Limit})
			return
		} else {
			menus = append(menus, *menu)
		}
	}
	var dayCount int
	var orderNum primitive.ObjectID
	// 메뉴별로 판매 수량 및 판매 가능 수량을 업데이트해준다.
	for idx, menu := range menus {
		// 있다면 수량만큼 limit--, count++
		// goroutine으로 실행할지 여부 생각하기
		bc.MenuModel.LimitAndCountUpdate(menu.ID, menu.Limit, menu.Orderedcount, list.OrderedMenus[idx].Amount)
	}
	// 한 개의 주문이기 때문에 dayCount++ 와 주문 추가 작업은 한 번만 해준다.
	// 그리고 DayOrderCount 1 추가
	dayCount = bc.OrderedListModel.DayOrderCount(daycountID)
	// 주문 추가
	orderNum = bc.OrderedListModel.Add(list)

	c.JSON(200, gin.H{"주문아이디" : orderNum, "오늘 주문번호" : dayCount})
}


// 메뉴 변경하는 함수
// AddOrder godoc
// @Summary 주문에서 메뉴를 변경합니다.
// @Tags Buyer
// @Description 주문서 메뉴를 변경하는 함수 -> "배달중", "배달완료" 상태일 시에는 에러 반환
// @name ChangeOrder
// @Accept  json
// @Produce  json
// @Param changeData body model.ChangeMenuType true "changeData"
// @Router /buyer/order/change [patch]
// @Success 200 {object} string
// @failure 400 {object} string
func (bc *BuyerController) ChangeOrder(c *gin.Context) {
	changeMenuStruct := &model.ChangeMenuType{}
	err := c.ShouldBindJSON(changeMenuStruct)
	util.ErrorHandler(err)

	order := bc.OrderedListModel.GetOne(changeMenuStruct.OrderId)
	legacyMenu, _ := bc.MenuModel.GetOneMenu(changeMenuStruct.LegacyFoodId)
	var legacyFoodAmount int
	for _, menu := range order.OrderedMenus {
		if menu.MenuId == legacyMenu.ID {
			legacyFoodAmount = menu.Amount
		}
	}
	newMenu, _ := bc.MenuModel.GetOneMenu(changeMenuStruct.NewMenu.MenuId)
	// 추가하려는 메뉴가 주문 가능한 수량을 가지고 있는지 확인하는 로직 추가
	if newMenu.Limit < changeMenuStruct.NewMenu.Amount {
		c.JSON(400, gin.H{"msg" : "해당 메뉴의 수량이 부족합니다.", "남은 수량" : newMenu.Limit})
		return
	} else if (order.Status == "배달중") || (order.Status == "배달완료") {
		c.JSON(400, gin.H{"msg" : "조리가 완료되어 메뉴 변경이 불가능합니다."})
		return
	} else {
		err := bc.OrderedListModel.ChangeOrder(order, changeMenuStruct)
		if err != nil {
			// 원래 주문에 legacyfood가 포함되어 있지 않을 경우 에러 반환
			c.JSON(400, gin.H{"error" : err.Error()})
			return
		} else {
			// 주문에서 제외한 음식과 변경된 음식의 limit, orderedcount 업데이트 로직 추가
			// legacy 메뉴에서는 limit++, count--를 해줘야 한다.
			bc.MenuModel.LimitAndCountUpdate(legacyMenu.ID, legacyMenu.Limit, legacyMenu.Orderedcount, -legacyFoodAmount)
			// new 메뉴에서는 limit--, count++를 해줘야 한다.
			bc.MenuModel.LimitAndCountUpdate(newMenu.ID, newMenu.Limit, newMenu.Orderedcount, changeMenuStruct.NewMenu.Amount)
			c.JSON(200, gin.H{"msg" : "메뉴 변경이 완료되었습니다."})
		}
	}
}


// 메뉴 추가하는 함수
// AddOrder godoc
// @Summary 주문에 메뉴를 추가합니다.
// @Tags Buyer
// @Description 주문에 메뉴를 추가하는 함수 -> "배달중", "배달완료" 상태일 시에는 에러 반환
// @name AddOrder
// @Accept  json
// @Produce  json
// @Param addData body model.AddMenuType true "addData"
// @Router /buyer/order/add [patch]
// @Success 200 {object} string
func (bc *BuyerController) AddOrder(c *gin.Context) {
	var msg string
	var id primitive.ObjectID

	addStruct := model.AddMenuType{}
	err := c.ShouldBindJSON(&addStruct)
	util.ErrorHandler(err)

	// 해당 주문의 현재 상태를 검사한다.
	order := bc.OrderedListModel.GetOne(addStruct.OrderId)
	// 이전 주문 메뉴에 현재 추가하려는 메뉴가 있는지 검사
	for _, value := range order.OrderedMenus {
		if value.MenuId == addStruct.NewItem.MenuId {
			c.JSON(400, gin.H{"msg" : "이미 해당 메뉴가 포함되어 있습니다."})
			return
		}
	}
	newMenu, _ := bc.MenuModel.GetOneMenu(addStruct.NewItem.MenuId)
	// 새로 추가하고자 하는 메뉴에 대해서 주문 가능한 limit이 있는지를 먼저 확인하는 로직 추가
	if newMenu.Limit < addStruct.NewItem.Amount {
		c.JSON(400, gin.H{"msg" : "해당 메뉴의 수량이 부족합니다.", "남은 수량" : newMenu.Limit})
		return
	}
	if (order.Status == "배달중") || (order.Status == "배달완료") {
		// 주문 수량만큼 limit--, count++ 해주는 로직 추가
		bc.MenuModel.LimitAndCountUpdate(newMenu.ID, newMenu.Limit, newMenu.Orderedcount, addStruct.NewItem.Amount)
		id = bc.OrderedListModel.Add(&addStruct.NewOrder)
		msg = "주문 추가가 불가능하여 새 주문으로 접수되었습니다."
	} else {
		// 추가한 메뉴에 대하여 limit과 count를 업로드해주는 로직 추가
		bc.MenuModel.LimitAndCountUpdate(newMenu.ID, newMenu.Limit, newMenu.Orderedcount, addStruct.NewItem.Amount)
		id = bc.OrderedListModel.AddOrder(&addStruct, order)
		msg = "주문 추가가 정상적으로 완료되었습니다."
	}
	c.JSON(200, gin.H{"msg" : msg, "id" : id})
}


// 메뉴목록 전체를 가져오는 함수
// GetAll godoc
// @Summary 모든 메뉴를 가져옵니다. 페이지당 5개의 데이터를 가지고 있습니다.
// @Tags Buyer
// @Description 모든 메뉴 리스트를 가져오기 위한 함수
// @name GetAll
// @Accept  json
// @Produce  json
// @Param page query string true "page"
// @Router /buyer/menu [get]
// @Success 200 {object} string
func (bc *BuyerController) GetAll(c *gin.Context) {
	sPage := c.Query("page")
	iPage, _ := strconv.Atoi(sPage)
	list := bc.MenuModel.GetAllMenu(int64(iPage))

	c.JSON(200, gin.H{"menus" : list})
}