package main

import (
	"github.com/tyspice/mr-grocery/routes"
)

func main() {
	r := routes.InitRouter()
	r.Run("localhost:12312")
}
