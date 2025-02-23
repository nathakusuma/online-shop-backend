package main

import "online-shop-backend/internal/bootstrap"

func main() {
	if err := bootstrap.Start(); err != nil {
		panic(err)
	}
}
