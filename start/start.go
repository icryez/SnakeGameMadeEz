package start

import (
	"SnakeGame/colors"
	"SnakeGame/structs"
	"fmt"
	"math/rand"
	"time"
)

var g_grid [30][30]structs.Cell
var g_snake structs.SnakeBody

func StartGame(grid [30][30]structs.Cell, snakeBody structs.SnakeBody) {
	
	CallClear()
	g_grid = grid
	g_snake = snakeBody
	
	PrintGrid()
	for true {
		time.Sleep(1 * time.Second)
		moveSnakeHead()
		moveSnakeBody()
		CallClear()
		PrintGrid()
	}
}

func PrintGrid() {
	for r := range g_grid {
		for c, colvalue := range g_grid[r] {
			if r == g_snake.Head[0] && c == g_snake.Head[1] {
				colors.Yellow.Print(" ", " ")
			} else if colvalue.IsBait {
				colors.Red.Print(" ", " ")
			}else if colvalue.IsSnakeBody{
				colors.Green.Print(" ", " ")	
			}else {
				fmt.Print("-", " ")
			}
		}
		fmt.Println()
	}
}

func moveSnakeHead() {
	g_snake.Head[0]++
}

func moveSnakeBody(){
	for range g_snake.Body {

	}
}

