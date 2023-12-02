package main

import "awesomeProject/routers"

func main() {
	r := routers.Router()
	r.Run(":8080")
}
