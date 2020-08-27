package main

import (
	"math"
	"math/rand"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
	// "net/http"
)

func main() {
	router := gin.Default()

	router.GET("/sudoku/generate/:level", func(c *gin.Context) {
		param := c.Param("level")
		level, err := strconv.Atoi(param)
		if err != nil {
			c.JSON(500, gin.H{
				"Error":   "Send a number 1 to 10 as request",
				"Info":    "1 -> Easiest, 10 -> Hardest",
				"Message": "Invalid format",
			})
		} else if level < 1 || level > 10 {
			c.JSON(500, gin.H{
				"Error": "You should send a number 1 to 10",
				"Info":  "1 -> Easiest, 10 -> Hardest",
			})
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
			c.JSON(400, gin.H{
				"result": sample,
			})
		}

	})

	router.POST("/sudoku/validate", func(c *gin.Context) {
		var req [][]int
		if c.BindJSON(&req) != nil {
			c.JSON(500, gin.H{
				"Error": "Expected 9x9 integer array",
			})
		} else {
			expectedSorted := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

			for i := 0; i < 9; i++ {

				column := make([]int, 9)
				for j := 0; j < 9; j++ {
					column[j] = req[j][i]
				}
				sort.Ints(column)
				row := make([]int, 9)
				copy(row, req[i])
				sort.Ints(column)
				sort.Ints(row)

				//row check
				if !Equal(row, expectedSorted) {
					c.JSON(400, gin.H{
						"isSuccess": false,
					})
					return
				}

				//column check
				if !Equal(column, expectedSorted) {
					c.JSON(400, gin.H{
						"isSuccess": false,
					})
					return
				}

				//3x3 square check
				square := GetInnerSquareAsLine(req, int(math.Floor(float64(i/3))),i%3)
				if !Equal(square, expectedSorted){
					c.JSON(400, gin.H{
						"isSuccess": false,
					})
					return
				}
			}

			c.JSON(400, gin.H{
				"isSuccess": true,
			})
		}
	})

	router.Run(":3000")
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for index, value := range a {
		if value != b[index] {
			return false
		}
	}
	return true
}

func GetInnerSquareAsLine(a [][]int, row int, column int) []int {
	innerSquare := make([]int, 9)
	for i := row*3; i < row*3+3; i++ {
		for j := column*3; j < column*3+3; j++ {
			innerSquare[(i%3)*3+j%3] = a[i][j]
		}
	}
	sort.Ints(innerSquare)
	return innerSquare
}
