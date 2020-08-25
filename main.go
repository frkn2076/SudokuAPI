package main

import (
	"github.com/gofiber/fiber"
	"math"
	"math/rand"
	"strconv"
)

func main() {
	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) {
		level, err := strconv.Atoi(c.Body())

		if err != nil {
			c.Send(err)
		} else{
			sequence := make([]int, 18)
		for i := range sequence {
			sequence[i] = i % 9 + 1
		}
		sample := make([][]int, 9)
		for i := 0; i < 9; i++ {
			temp := make([]int, 9)
			copy(temp,sequence[9-int(math.Floor(float64(i/3)))-(i%3)*3 : 9-int(math.Floor(float64(i/3)))-(i%3)*3+9])
			sample[i] = temp
		}

		for i := 0; i < 25; i++ {

			rand1 := rand.Intn(9) + 1
			rand2 := rand.Intn(9) + 1
			
			if(rand1 != rand2){
				for a := 0; a < 9; a++ {
					for b := 0; b < 9; b++ {
						if(sample[a][b] == rand1){
							sample[a][b] = rand2
						} else if(sample[a][b] == rand2){
							sample[a][b] = rand1
						}				

					}
				}
			}
		}
		
		for i := 0; i < level; i++ {
			for i := 0; i < level * 2; i++ {
				sample[rand.Intn(9)][rand.Intn(9)] = 0
			}
		}

		}

	
		
	})

	app.Listen(3000)
}

type SudokuLevel struct {
    level int
}
