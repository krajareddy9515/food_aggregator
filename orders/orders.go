package orders

import (
	"time"

	"github.com/emicklei/go-restful"
	"github.com/patrickmn/go-cache"
)

type Order struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Price    string `json:"price"`
}

var c = cache.New(5*time.Minute, 10*time.Minute)

func Register(container *restful.Container) {
	ws := Webservice()
	container.Add(ws)
}

func Webservice() *restful.WebService {

	ws := new(restful.WebService)
	ws.Path("v1/orders")

	ws.Route(ws.POST("/buy-item").To(BuyItem))
	ws.Route(ws.POST("/buy-item-qty").To(BuyItemQty))
	ws.Route(ws.POST("/buy-item-qty-price").To(BuyItemQtyPrice))
	ws.Route(ws.GET("/show-summary").To(ShowSummery))
	ws.Route(ws.POST("/fast-buy-item").To(FastBuyItem))
	return ws
}
