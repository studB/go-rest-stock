package main

import (
	"fmt"
	"rest-stock/internal/api"
)

var ()

func main() {

	rs, _ := api.NewAuthClient().RequestToken()

	fmt.Println("Current token : ", rs)

	// manger := auth.GetManager()

	// for i := 0; i < 3; i++ {
	// 	fmt.Println("Current token : ", manger.GetToken())
	// 	time.Sleep(2 * time.Second)
	// }

	// READ Required Variable

	// Set Scheduling by go runtine
}
