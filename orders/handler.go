package orders

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	restful "github.com/emicklei/go-restful"
	log "github.com/sirupsen/logrus"
)

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
	if name == "" {
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
				resp.WriteAsJson(result[i])
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
	if name == "" || quantity == 0 {
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
				resp.WriteAsJson(result[i])
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
	if name == "" || quantity == 0 {
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
				resp.WriteAsJson(result[i])
				return
			}
		}
	}
	resp.WriteAsJson("NOT_FOUND")
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
