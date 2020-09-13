package main

import (
	"SudokuAPI/Router"
	"os"
)

func main() {
	router := Router.SetupRouter()
	// router.Run(":3001")        //uncomment for local
	router.Run(os.Getenv("PORT")) //comment for local
}
