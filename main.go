package main

import (
	"fmt"
	v1 "myfmw/v1"
	"strconv"
)


func main() {
	fmt.Println("this is main")

    app :=v1.New()

	app.Add("GET", "/member/:name/:age", func(ctx *v1.Context) {
		name :=ctx.Param.ByName("name")
		age,_ := strconv.Atoi(ctx.Param.ByName("age"))

		type ss struct {
			Name string `json:"name"`
			Age int `json:"age"`
		}

		test := ss{name, age}
		ctx.Json(test)
	})
    app.Run(":9988")
}
