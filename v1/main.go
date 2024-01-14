package main

import (
	"douyin/v1/config"
	"douyin/v1/router"
	"fmt"
)

func main() {
	fmt.Println(11)
	r := router.InitDouyinRouter()
	err := r.Run(fmt.Sprintf(":%d", config.Info.Port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
}
