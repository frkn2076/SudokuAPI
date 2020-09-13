package main

import (
	"app/SudokuAPI/Router"
	// "os"
)

func main() {
	router := Router.SetupRouter()
	router.Run(":3000")        			//uncomment for local
	// router.Run(os.Getenv("PORT")) 	//comment for local

}
