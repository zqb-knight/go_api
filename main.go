package main

import (
	"go_api/router"
)

func main() {
	r := router.InitRouter()

	_ = r.Run()

}
