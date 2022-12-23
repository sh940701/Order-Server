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

	sRoute := r.Group("/seller")
	{
		sRoute.GET("/orderlist", p.seller.GetOrderList)
		sRoute.GET("/statusupdate/:orderid", p.seller.UpdateOrderStatus)
		sRoute.POST("/addmenu", p.seller.AddMenu)
		sRoute.POST("/delete", p.seller.DeleteMenu)
		sRoute.PUT("/updatemenu", p.seller.UpdateMenu)
	}
	
	bRoute := r.Group("/buyer")
	{
		bRoute.GET("/getlist/:category", p.buyer.GetMenuList)
		bRoute.GET("/getreview/:menuid", p.buyer.GetReview)
		bRoute.GET("/ordered/:orderid", p.buyer.GetOrderStatus)
		bRoute.POST("/addreview/:foodid", p.buyer.AddReview)
		bRoute.POST("/order", p.buyer.Order)
		bRoute.PUT("/addorder", p.buyer.AddOrder)
		bRoute.PUT("/changeorder", p.buyer.ChangeOrder)

	}

	return r
}