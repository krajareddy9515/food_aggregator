package orders

import "github.com/emicklei/go-restful"

type Order struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Price    string `json:"price"`
}

var urls = []string{"https://run.mocky.io/v3/c51441de-5c1a-4dc2-a44e-aab4f619926b",
	"https://run.mocky.io/v3/4ec58fbc-e9e5-4ace-9ff0-4e893ef9663c",
	"https://run.mocky.io/v3/e6c77e5c-aec9-403f-821b-e14114220148",
}

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
	// ws.Route(ws.GET("/show-summary").To(ShowSummery))
	// ws.Route(ws.GET("/fast-buy-item").To(FastBuyItem))
	return ws

}
