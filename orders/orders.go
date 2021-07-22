package orders

import "github.com/emicklei/go-restful"

func Register(container *restful.Container) {
	ws := Webservice()
	container.Add(ws)
}

func Webservice() *restful.WebService {

	ws := new(restful.WebService)
	ws.Path("v1/orders")

	ws.Route(ws.POST("/buy-item").To(BuyItem))
	ws.Route(ws.POST("/buy-item-qty").To(BuyItemQty))
	// ws.Route(ws.GET("/buy-item-qty-price").To(BuyItemQtyPrice))
	// ws.Route(ws.GET("/show-summary").To(ShowSummery))
	// ws.Route(ws.GET("/fast-buy-item").To(FastBuyItem))
	return ws

}
