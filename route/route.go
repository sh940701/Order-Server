package route

import (
	"github.com/gin-gonic/gin"
	"github.com/codestates/WBABEProject-08/commits/main/controller"
)

type Router struct {
	buyer *controller.BuyerController
	seller *controller.SellerController
}

func NewRouter(ctl *controller.Controller) *Router {
	r := &Router{}
	r.buyer = ctl.GetBuyerController()
	r.seller = ctl.GetSellerController()

	return r
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		//허용할 header 타입에 대해 열거
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, X-Forwarded-For, Authorization, accept, origin, Cache-Control, X-Requested-With")
		//허용할 method에 대해 열거
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (p *Router) Idx() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(CORS())

	sGroup := r.Group("/seller")
	{
		// 전체 주문 목록 가져오기
		sGroup.GET("/order", p.seller.GetOrderList)
		// 주문 상태 변경
		sGroup.PATCH("/order", p.seller.UpdateOrderStatus)
		// 메뉴 추가하기
		sGroup.POST("/menu", p.seller.AddMenu)
		// 메뉴 업데이트하기
		sGroup.PATCH("/menu", p.seller.UpdateMenu)
		// 메뉴 지우기
		sGroup.POST("/delete", p.seller.DeleteMenu)
		// 추천메뉴 설정하기
		sGroup.PATCH("/suggestion", p.seller.SuggestMenu)
	}
	
	bGroup := r.Group("/buyer")
	{
		// 전체 메뉴 가져오기 -> 해야함
		bGroup.GET("/menu", p.buyer.GetAll)
		// 카테고리별 메뉴 정렬하여 가져오기(추천, 평점, 구매횟수)
		bGroup.GET("/menu/:category", p.buyer.GetMenuList)
		// 메뉴별 리뷰 가져오기
		bGroup.GET("/review/:menuid", p.buyer.GetReview)
		// 주문한 음식에 대한 리뷰 작성
		bGroup.POST("/review/:menuid", p.buyer.AddReview)
		// 주문 상태 가져오기
		bGroup.GET("/order/:orderid", p.buyer.GetOrderStatus)
		// 주문하기
		bGroup.POST("/order", p.buyer.Order)
		// 주문에 메뉴 추가하기
		bGroup.PATCH("/order/add", p.buyer.AddOrder)
		// 주문에서 메뉴 바꾸기
		bGroup.PATCH("/order/change", p.buyer.ChangeOrder)
	}

	return r
}