package main

import (
	"7daysGolang/gee"
	"fmt"
	"net/http"
)

func main() {
	defer fmt.Println("stop running")
	gee := gee.New()
	gee.GET("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "hello my friends, this page is GET ")
	})

	gee.POST("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "hello my friends, this page is POST ")
	})
	fmt.Println("start running")
	gee.Run(":9999")
}
