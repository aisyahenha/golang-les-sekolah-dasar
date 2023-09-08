package main

import (
	"fmt"

	"github.com/aisyahenha/golang-les-sekolah-dasar/delivery"
)

func main ()  {
	fmt.Println("hello gais")
	delivery.NewServer().Run()
}

