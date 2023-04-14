package main

import (
	"log"

	"github.com/CbIPOKGIT/alkomag/pkg/printer"
)

func main() {
	// if err := offers.CollectOffers(); err != nil {
	// 	log.Fatal(err)
	// }

	// if err := details.CollectData(); err != nil {
	// 	log.Fatal(err)
	// }

	if err := printer.PrintResults(); err != nil {
		log.Fatal(err)
	}

	log.Println("Success")
}

// https://rozetka.com.ua/api/product-api/v4/goods/get-characteristic?front-type=xl&country=UA&lang=ua&goodsId=5905449
// https://rozetka.com.ua/api/product-api/v4/goods/get-main?front-type=xl&country=UA&lang=ua&goodsId=351207465
