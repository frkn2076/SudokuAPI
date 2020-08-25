package main

import (
	"math"
	"math/rand"
	"strconv"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	app.Post("/sudoku/generate", func(c *fiber.Ctx) {
		level, err := strconv.Atoi(c.Body())

		if err != nil {
			c.Send("Send a number 1 to 10 as request\n1  -> Easiest\n10 -> Hardest")
		} else if level < 1 || level > 10 {
			c.Send("You should send a number 1 to 10.")
		} else {

			sequence := make([]int, 18)
			for i := range sequence {
				sequence[i] = i%9 + 1
			}

			sample := make([][]int, 9)
			for i := 0; i < 9; i++ {
				temp := make([]int, 9)
				copy(temp, sequence[9-int(math.Floor(float64(i/3)))-(i%3)*3:9-int(math.Floor(float64(i/3)))-(i%3)*3+9])
				sample[i] = temp
			}

			for i := 0; i < 25; i++ {

				rand1 := rand.Intn(9) + 1
				rand2 := rand.Intn(9) + 1

				if rand1 != rand2 {
					for a := 0; a < 9; a++ {
						for b := 0; b < 9; b++ {
							if sample[a][b] == rand1 {
								sample[a][b] = rand2
							} else if sample[a][b] == rand2 {
								sample[a][b] = rand1
							}

						}
					}
				}
			}

			for i := 0; i < level*10; i++ {
				sample[rand.Intn(9)][rand.Intn(9)] = 0
			}

			c.Send(sample)
		}

	})

	app.Post("/sudoku/validate", func(c *fiber.Ctx) {
		req := new([][]int)
		if err := c.BodyParser(req); err != nil {
			c.Send("Expected 9x9 integer array")
		} else {
			// expectedSorted := []int{1,2,3,4,5,6,7,8,9}

			// //row check
			// for i := 0; i < 9; i++ {
			// 	temp := (*req)[i]
			// 	if temp != expectedSorted {
			// 		return false
			// 	}
			// }	


		}
	})

	app.Listen(3000)
}
