package start

import (
	"SnakeGame/colors"
	"SnakeGame/structs"
	"fmt"
	"math/rand"
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
var tempBait structs.Node
var canSearch bool
var firstBait bool

func StartGame(grid [30][30]structs.Cell, snakeBody structs.SnakeBody) {

	CallClear()
	g_grid = grid
	g_snake = snakeBody
	g_searchedCells = make(map[[2]int]bool)
	firstBait = true
	PrintGrid()
	for !g_gameover {
		moveSnakeHead()
		CallClear()
		PrintGrid()

	}
}
func GenerateRandomBait() {
	x := rand.Intn(29)
	y := rand.Intn(29)
	g_grid[x][y].IsBait = true
	g_grid[x][y].Value = 1
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
			} else if colvalue.IsSnakeBody {
				colors.Green.Print(" ", " ")
			} else if colvalue.IsVisible && !colvalue.IsPath {
				colors.Blue.Print(" ", " ")
			} else if colvalue.IsPath {
				colors.Red.Print(" ", " ")
			} else {
				fmt.Print(" ", " ")
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
	time.Sleep(50 * time.Millisecond)
	nextSnakeHead()
	if pathFromSearch != nil {
		for _, v := range pathFromSearch {
			time.Sleep(130 * time.Millisecond)
			moveSnakeBody()
			g_snake.Head = v
			CallClear()
			PrintGrid()
			if g_snake.Head == g_foundBait {
				GenerateRandomBait()
				g_grid[tempBait.Value[0]][tempBait.Value[1]].IsBait = false
				pathFromSearch = nil
				g_searchedCells = make(map[[2]int]bool)
				g_SearchOver = false
				temp := &g_snake.Body
				if temp != nil && temp.Next != nil {
					temp = temp.Next
				}
				if temp != nil {
					temp.Next = newBodyNode(g_snake.Head)
				}
				if temp == nil {
					g_snake.Body = *newBodyNode(g_snake.Head)
				}
			}
		}

	}
}

func newSearchNode(r int, c int, prevNode *structs.Node) *structs.Node {
	newNode := *new(structs.Node)
	newNode.Next = prevNode
	newNode.Value = [2]int{r, c}
	return &newNode
}

func nextSnakeHead() {
	searchBait(g_snake.Head[0], g_snake.Head[1], newSearchNode(g_snake.Head[0], g_snake.Head[1], nil))
	time.Sleep(50 * time.Millisecond)

	makeSnakePath()
}

var pathFromSearch [][2]int

func makeSnakePath() {
	temp := tempBait
	for temp.Next != nil {
		pathFromSearch = append(pathFromSearch, temp.Value)
		temp = *temp.Next
	}
	for i, j := 0, len(pathFromSearch)-1; i < j; i, j = i+1, j-1 {
		pathFromSearch[i], pathFromSearch[j] = pathFromSearch[j], pathFromSearch[i]
	}
}

func searchBait(r int, c int, prevNode *structs.Node) {
	newNode := newSearchNode(r, c, prevNode)
	time.Sleep(1 * time.Millisecond)
	g_searchedCellsMutex.Lock()
	g_searchedCells[[2]int{r, c}] = true
	g_searchedCellsMutex.Unlock()
	if g_grid[r][c].IsBait && !g_SearchOver {
		tempBait = *newNode
		g_foundBait = [2]int{r, c}
		g_SearchOver = true
		return
	}
	g_searchedCellsMutex.RLock()
	if !g_SearchOver && c+1 < 30 && !g_searchedCells[[2]int{r, c + 1}] && !g_grid[r][c+1].IsSnakeBody {
		go searchBait(r, c+1, newNode)
	}
	g_searchedCellsMutex.RUnlock()
	g_searchedCellsMutex.RLock()
	if !g_SearchOver && r-1 >= 0 && !g_searchedCells[[2]int{r - 1, c}] && !g_grid[r-1][c].IsSnakeBody {

		go searchBait(r-1, c, newNode)
	}
	g_searchedCellsMutex.RUnlock()
	g_searchedCellsMutex.RLock()
	if !g_SearchOver && c-1 >= 0 && !g_searchedCells[[2]int{r, c - 1}] && !g_grid[r][c-1].IsSnakeBody {

		go searchBait(r, c-1, newNode)
	}
	g_searchedCellsMutex.RUnlock()
	g_searchedCellsMutex.RLock()
	if !g_SearchOver && r+1 < 30 && !g_searchedCells[[2]int{r + 1, c}] && !g_grid[r+1][c].IsSnakeBody {

		go searchBait(r+1, c, newNode)
	}
	g_searchedCellsMutex.RUnlock()
}

func moveSnakeBody() {
	if &g_snake.Body == nil {
		g_snake.Body = *newBodyNode(g_snake.Head)
	}
	temp := &g_snake.Body
	for temp.Next != nil {
		g_grid[temp.Value[0]][temp.Value[1]].IsSnakeBody = true
		temp = temp.Next
	}
	g_grid[g_snake.Body.Value[0]][g_snake.Body.Value[1]].IsSnakeBody = false
	if g_snake.Body.Next != nil {
		g_snake.Body = *g_snake.Body.Next
	}
}
