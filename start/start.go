package start

import (
	"SnakeGame/colors"
	"SnakeGame/structs"
	"fmt"
	"sync"
	"time"
)

var g_grid [30][30]structs.Cell
var g_snake structs.SnakeBody
var g_searchedCells map[[2]int]bool
var g_gameover bool
var g_searchedCellsMutex sync.RWMutex
var g_SearchOver bool
var g_foundBait [2]int

func StartGame(grid [30][30]structs.Cell, snakeBody structs.SnakeBody) {

	CallClear()
	g_grid = grid
	g_snake = snakeBody
	g_searchedCells = make(map[[2]int]bool)
	// g_snake.Body = *newBodyNode([2]int{g_snake.Head[0] - 1, g_snake.Head[1]})
	// g_snake.Body.Next = newBodyNode([2]int{g_snake.Head[0] - 2, g_snake.Head[1]})
	// g_snake.Body.Next.Next = newBodyNode([2]int{g_snake.Head[0] - 3, g_snake.Head[1]})
	// g_snake.Body.Next.Next.Next = newBodyNode([2]int{g_snake.Head[0] - 4, g_snake.Head[1]})
	// g_grid[g_snake.Head[0]-1][g_snake.Head[1]].IsSnakeBody = true
	// g_grid[g_snake.Head[0]-2][g_snake.Head[1]].IsSnakeBody = true
	// g_grid[g_snake.Head[0]-3][g_snake.Head[1]].IsSnakeBody = true
	// g_grid[g_snake.Head[0]-4][g_snake.Head[1]].IsSnakeBody = true
	PrintGrid()
	for !g_gameover {
		time.Sleep(500 * time.Millisecond)
		moveSnakeHead()
		moveSnakeBody()
		CallClear()
		PrintGrid()
	}
}

func PrintGrid() {
	fmt.Println("  _______________________________________________________________")
	for r := range g_grid {
		fmt.Print("  | ")
		for c, colvalue := range g_grid[r] {
			if r == g_snake.Head[0] && c == g_snake.Head[1] {
				colors.Yellow.Print(" ", " ")
			} else if colvalue.IsBait {
				colors.Red.Print(" ", " ")
				// } else if colvalue.IsSnakeBody {
				// 	colors.Green.Print(" ", " ")
			} else if colvalue.IsVisible {
				colors.Blue.Print(" ", " ")
			} else {
				fmt.Print("-", " ")
			}
		}
		fmt.Println("|")
	}
	fmt.Println("  _______________________________________________________________")
}

func newBodyNode(head [2]int) *structs.Node {
	newNode := *new(structs.Node)
	newNode.Value = head
	return &newNode
}

func moveSnakeHead() {
	temp := &g_snake.Body
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = newBodyNode(g_snake.Head)
	g_grid[g_snake.Head[0]][g_snake.Head[1]].IsSnakeBody = true
	// g_snake.Head[1]++
	// g_snake.Head = nextSnakeHead()
	nextSnakeHead()
}

func nextSnakeHead() {
	searchBait(g_snake.Head[0], g_snake.Head[1])
}

func searchBait(r int, c int) {
	g_grid[r][c].IsVisible = true
	
	time.Sleep(1 * time.Millisecond)
	g_searchedCellsMutex.Lock()
	g_searchedCells[[2]int{r, c}] = true
	g_searchedCellsMutex.Unlock()
	if g_grid[r][c].IsBait {
		g_foundBait = [2]int{r,c}
		g_SearchOver = true
		return
	}
	g_searchedCellsMutex.RLock()
	if !g_SearchOver  && c+1 < 30 && !g_searchedCells[[2]int{r, c + 1}] {
		go searchBait(r, c+1)
	}
	g_searchedCellsMutex.RUnlock()
	g_searchedCellsMutex.RLock()
	if !g_SearchOver && r-1 >= 0 && !g_searchedCells[[2]int{r - 1, c}] {

		go searchBait(r-1, c)
	}
	g_searchedCellsMutex.RUnlock()
	g_searchedCellsMutex.RLock()
	if !g_SearchOver && c-1 >= 0 && !g_searchedCells[[2]int{r, c - 1}] {

		go searchBait(r, c-1)
	}
	g_searchedCellsMutex.RUnlock()
	g_searchedCellsMutex.RLock()
	if !g_SearchOver && r+1 < 30 && !g_searchedCells[[2]int{r + 1, c}] {

		go searchBait(r+1, c)
	}
	g_searchedCellsMutex.RUnlock()
	// return -1, -1
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
