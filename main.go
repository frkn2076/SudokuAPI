package main

import (
	"SudokuAPI/Router"
)

func main() {
	router := Router.SetupRouter()
	router.Run(":3001")
}


