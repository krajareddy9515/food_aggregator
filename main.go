package main

import (
	"flag"
	"fmt"
	"food_aggregator/orders"
	"log"
	"net/http"
	"time"

	restful "github.com/emicklei/go-restful"
)

var port int

func main() {

	flag.IntVar(&port, "port", 8080, "Server Port")
	flag.Parse()

	var wait time.Duration

	flag.DurationVar(&wait, "graceful-timout", time.Second*15, "The duration server gracefully wait for existing connection to finish")
	flag.Parse()

	container := restful.NewContainer()
	container.Router(restful.CurlyRouter{})

	orders.Register(container)

	ws := new(restful.WebService)
	ws.Route(ws.GET("/health").To(healthCheck))
	container.Add(ws)

	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%d", port),
		WriteTimeout: time.Second * 60,
		ReadTimeout:  time.Second * 60,
		IdleTimeout:  time.Second * 60,
		Handler:      container,
	}
	log.Printf("Food Aggregator API server started at port:%d", port)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

//heathCheck : API to check health
func healthCheck(req *restful.Request, resp *restful.Response) {

	resBody := []byte(`{"status":"OK"}`)
	resp.Header().Set("Content-Type", "application/json")
	resp.Write(resBody)
}