package main

import "github.com/eymen-iron/go-life-hack/pkg/json"

func main() {

	dt, _ := json.JsonGenerator([]byte(`{Name:"data"}`))

	println(dt)
}
