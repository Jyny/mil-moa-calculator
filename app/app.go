package main

import (
	"log"
	"mil-moa-calculator/pkg/calc"
	"strconv"
	"syscall/js"
)

var (
	window, document, body js.Value
)

func init() {
	window = js.Global()
	document = window.Get("document")
	body = document.Get("body")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("recovered", r)
			main()
		}
	}()

	window.Set("calprecision", js.FuncOf(calPrecision))
	window.Set("caldistance", js.FuncOf(calDistance))

	select {}
}
func calPrecision(this js.Value, args []js.Value) interface{} {
	if len(args) < 2 {
		log.Println("args is not enough")
		return nil
	}

	x, err := strconv.ParseFloat(args[0].String(), 64)
	if err != nil {
		log.Println(err)
		return nil
	}

	y, err := strconv.ParseFloat(args[1].String(), 64)
	if err != nil {
		log.Println(err)
		return nil
	}

	return calc.Calculator{
		Mus: window.Get("mus").String(),
		Aus: window.Get("aus").String(),
	}.CalPrecision(x, y)
}

func calDistance(this js.Value, args []js.Value) interface{} {
	if len(args) < 2 {
		log.Println("args is not enough")
		return nil
	}

	x, err := strconv.ParseFloat(args[0].String(), 64)
	if err != nil {
		log.Println(err)
		return nil
	}

	y, err := strconv.ParseFloat(args[1].String(), 64)
	if err != nil {
		log.Println(err)
		return nil
	}

	return calc.Calculator{
		Mus: window.Get("mus").String(),
		Aus: window.Get("aus").String(),
	}.CalDistance(x, y)
}
