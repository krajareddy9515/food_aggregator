package orders

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	restful "github.com/emicklei/go-restful"
	"github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"
)

var fruitSupplier, vegSupplier, grainSupplier = "https://run.mocky.io/v3/c51441de-5c1a-4dc2-a44e-aab4f619926b",
	"https://run.mocky.io/v3/4ec58fbc-e9e5-4ace-9ff0-4e893ef9663c",
	"https://run.mocky.io/v3/e6c77e5c-aec9-403f-821b-e14114220148"

// BuyItem : API to check the item
func BuyItem(req *restful.Request, resp *restful.Response) {

	log.Printf(" In BuyItem API ")

	reqBody, _ := ioutil.ReadAll(req.Request.Body)
	order := Order{}

	err := json.Unmarshal(reqBody, &order)
	if err != nil {
		log.Error(err)
		resp.WriteErrorString(500, "Internal server error")
		return
	}

	name := order.Name
	if len(name) == 0 {
		resp.WriteErrorString(400, "Invalid request body")
		return
	}

	result, err := Suppliers(name)
	if err != nil {
		log.Error(err)
		resp.WriteErrorString(500, err.Error())
		return
	}
	if len(result) > 0 {
		for i := 0; i < len(result); i++ {
			if result[i].Name == name {
				c.Set(time.Now().Format("2006-01-02 15:04:05"), result[i].Name, cache.NoExpiration)
				resp.WriteAsJson(result[i].Name)
				return
			}
		}
	}
	resp.WriteAsJson("NOT_FOUND")
}

//BuyItemQty : API to check quantity
func BuyItemQty(req *restful.Request, resp *restful.Response) {
	log.Printf("In BuyItemQty API ")

	reqBody, _ := ioutil.ReadAll(req.Request.Body)
	order := Order{}

	err := json.Unmarshal(reqBody, &order)
	if err != nil {
		log.Error(err)
		resp.WriteErrorString(500, "Internal server error")
		return
	}

	name := order.Name
	quantity := order.Quantity
	if len(name) == 0 || quantity <= 0 {
		resp.WriteErrorString(400, "Invalid request body")
		return
	}

	result, err := Suppliers(name)
	if err != nil {
		log.Error(err)
		resp.WriteErrorString(500, err.Error())
		return
	}
	if len(result) > 0 {
		for i := 0; i < len(result); i++ {
			if result[i].Name == name && result[i].Quantity == quantity {
				c.Set(time.Now().Format("2006-01-02 15:04:05"), result[i], cache.NoExpiration)
				resp.WriteAsJson(result[i].Name)
				return
			}
		}
	}
	resp.WriteAsJson("NOT_FOUND")
}

//BuyItemQtyPrice : API to check quantity and price
func BuyItemQtyPrice(req *restful.Request, resp *restful.Response) {
	log.Printf(" In BuyItemQtyPrice API ")

	reqBody, _ := ioutil.ReadAll(req.Request.Body)
	order := Order{}

	err := json.Unmarshal(reqBody, &order)
	if err != nil {
		log.Error(err)
		resp.WriteErrorString(500, "Internal server error")
		return
	}

	name := order.Name
	quantity := order.Quantity
	price := order.Price
	log.Println(len(price))
	if len(name) == 0 || len(price) == 0 || quantity <= 0 {
		resp.WriteErrorString(400, "Invalid request body")
		return
	}

	result, err := Suppliers(name)
	if err != nil {
		log.Error(err)
		resp.WriteErrorString(500, err.Error())
		return
	}
	if len(result) > 0 {
		for i := 0; i < len(result); i++ {
			if result[i].Name == name && result[i].Quantity == quantity && result[i].Price == price {
				c.Set(time.Now().Format("2006-01-02 15:04:05"), result[i], cache.NoExpiration)
				resp.WriteAsJson(result[i].Name)
				return
			}
		}
	}
	resp.WriteAsJson("NOT_FOUND")
}

//ShowSummery : API to get summery
func ShowSummery(req *restful.Request, resp *restful.Response) {
	resp.WriteAsJson(c.Items())
}

// Suppliers: API to buy items from suppliers
func Suppliers(name string) ([]Order, error) {
	log.Printf(" In Supplier ")

	var client = &http.Client{}

	for _, v := range urls {
		req, err := http.NewRequest("GET", v, nil)
		if err != nil {
			log.Error(err)
			return nil, err
		}

		res, err := client.Do(req)
		if err != nil {
			log.Error(err)
			return nil, err
		}

		respBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Error(err)
			return nil, err
		}

		order := []Order{}
		errRes := json.Unmarshal(respBody, &order)

		if errRes != nil {
			log.Error(errRes)
			return nil, errRes
		}

		for i := 0; i < len(order); i++ {
			if order[i].Name == name {
				return order, nil
			}
		}
	}
	return nil, nil
}

// FastBuyItem : API to check the item
func FastBuyItem(req *restful.Request, resp *restful.Response) {

	log.Printf(" In BuyItem API ")

	reqBody, _ := ioutil.ReadAll(req.Request.Body)
	order := Order{}

	err := json.Unmarshal(reqBody, &order)
	if err != nil {
		log.Error(err)
		resp.WriteErrorString(500, "Internal server error")
		return
	}

	name := order.Name
	if len(name) == 0 {
		resp.WriteErrorString(400, "Invalid request body")
		return
	}

	c1 := make(chan Order)
	c2 := make(chan Order)
	c3 := make(chan Order)

	go FastSuppliers(name, c1, fruitSupplier)
	go FastSuppliers(name, c2, vegSupplier)
	go FastSuppliers(name, c3, grainSupplier)

	select {
	case result1 := <-c1:
		c.Set(time.Now().Format("2006-01-02 15:04:05"), result1, cache.NoExpiration)
		resp.WriteAsJson(result1.Name)
	case result2 := <-c2:
		c.Set(time.Now().Format("2006-01-02 15:04:05"), result2, cache.NoExpiration)
		resp.WriteAsJson(result2.Name)
	case result3 := <-c3:
		c.Set(time.Now().Format("2006-01-02 15:04:05"), result3, cache.NoExpiration)
		resp.WriteAsJson(result3.Name)
		// default:
		// 	resp.WriteAsJson("NOT_FOUND")
	}
}

// Suppliers: API to buy items from suppliers
func FastSuppliers(name string, c chan Order, url string) {
	log.Printf(" In Supplier ")

	var client = &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Error(err)
	}

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error(err)
	}

	order := []Order{}
	errRes := json.Unmarshal(respBody, &order)

	if errRes != nil {
		log.Error(errRes)
	}

	for i := 0; i < len(order); i++ {
		if order[i].Name == name {
			c <- order[i]
		}
	}
}
