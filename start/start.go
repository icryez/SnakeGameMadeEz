package start

import (
	"SnakeGame/colors"
	"SnakeGame/structs"
	"fmt"
	"time"
)

var g_grid [30][30]structs.Cell
var g_snake structs.SnakeBody

func StartGame(grid [30][30]structs.Cell, snakeBody structs.SnakeBody) {

	CallClear()
	g_grid = grid
	g_snake = snakeBody

	g_snake.Body = *newBodyNode([2]int{g_snake.Head[0]-1,g_snake.Head[1]})
	g_snake.Body.Next = newBodyNode([2]int{g_snake.Head[0]-2,g_snake.Head[1]})
	g_snake.Body.Next.Next = newBodyNode([2]int{g_snake.Head[0]-3,g_snake.Head[1]})
	g_snake.Body.Next.Next.Next = newBodyNode([2]int{g_snake.Head[0]-4,g_snake.Head[1]})
	g_grid[g_snake.Head[0]-1][g_snake.Head[1]].IsSnakeBody = true
	g_grid[g_snake.Head[0]-2][g_snake.Head[1]].IsSnakeBody = true
	g_grid[g_snake.Head[0]-3][g_snake.Head[1]].IsSnakeBody = true
	g_grid[g_snake.Head[0]-4][g_snake.Head[1]].IsSnakeBody = true
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
			} else if colvalue.IsSnakeBody {
				colors.Green.Print(" ", " ")
			} else {
				fmt.Print("-", " ")
			}
		}
		fmt.Println()
	}
}

func newBodyNode(head [2]int) *structs.Node{
	newNode := *new(structs.Node)
	newNode.Value = head
	return &newNode
}

func moveSnakeHead() {
	temp := &g_snake.Body
	for temp.Next != nil{ temp = temp.Next }
	temp.Next = newBodyNode(g_snake.Head)
	g_grid[g_snake.Head[0]][g_snake.Head[1]].IsSnakeBody = true
	g_snake.Head[1]++
}

func moveSnakeBody() {
	temp := &g_snake.Body
	for temp.Next != nil {
		g_grid[temp.Value[0]][temp.Value[1]].IsSnakeBody = true
		temp = temp.Next
	}
	g_grid[g_snake.Body.Value[0]][g_snake.Body.Value[1]].IsSnakeBody = false
	g_snake.Body = *g_snake.Body.Next
}
