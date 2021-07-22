# food_aggregator


1. buy-item API

    URL: http://localhost:8080/v1/orders/buy-item
    Method: POST
    Request Body:
                {
                     "name": "Apple",

                }
    Status: 200 OK
    Response Body:
                {
                    "id": "24-583-0264",
                    "name": "Apple",
                    "quantity": 30,
                    "price": "$62.02"
                }


2. buy-item-qty API

    URL: http://localhost:8080/v1/orders/buy-item-qty
    Method: POST
    Request Body:
                {
                    "name": "Apple",
                    "quantity": 30

                }
    Status: 200 OK
    Response Body:
                {
                    "id": "24-583-0264",
                    "name": "Apple",
                    "quantity": 30,
                    "price": "$62.02"
                }


3. buy-item-qty-price API

    URL: http://localhost:8080/v1/orders/buy-item-qty-price
    Method: POST
    Request Body:
                {
                    "name": "Apple",
                    "quantity": 30,
                    "price": "$62.02"

                }
    Status: 200 OK
    Response Body:
                {
                    "id": "24-583-0264",
                    "name": "Apple",
                    "quantity": 30,
                    "price": "$62.02"
                }


4. show-summary

    URL: http://localhost:8080/v1/orders/show-summary
    Method: GET
    Status: 200 OK
    Response Body:
                {
                    "id": "24-583-0264",
                    "name": "Apple",
                    "quantity": 30,
                    "price": "$62.02"
                }


5. fast-buy-item

    URL: http://localhost:8080/v1/orders/fast-buy-item
    Method: POST
    Request Body:
                {
                    "name": "Apple",
                    "quantity": 30,
                    "price": "$62.02"

                }
    Status: 200 OK
    Response Body:
                {
                    "id": "24-583-0264",
                    "name": "Apple",
                    "quantity": 30,
                    "price": "$62.02"
                }