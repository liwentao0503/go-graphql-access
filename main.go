package main

import (
	"fmt"
	"net/http"

	"go-graphql-access/router"
)

func main() {
	if err := Init(); err != nil {
		panic(fmt.Sprintf("init err %s", err.Error()))
	}

	for k, v := range router.Router {
		http.Handle(k, v)
	}

	// listen and serve
	http.ListenAndServe(":8080", nil)
}

func Init() error {
	// if err := mysql.InitDB(); err != nil {
	// 	return err
	// }
	// if err := redis.InitDB(); err != nil {
	// 	return err
	// }
	return nil
}
