package main

import v1 "github.com/Peterliang233/techtrainingcamp-AppUpgrade/router/v1"

func main() {
	router := v1.InitRouter()

	if err := router.Run(":9090"); err != nil {
		panic(err)
	}
}
