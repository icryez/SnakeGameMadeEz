package main

import (
	"math/rand"
	"SnakeGame/start"
	"SnakeGame/structs"
)

var Snake structs.SnakeBody
var Grid [30][30]structs.Cell
func main() {

	generateRandomBait()
	genRandomStartPoint()
	start.StartGame(Grid,Snake)
}


func generateRandomBait() {
	x := rand.Intn(29)
	y := rand.Intn(29)
	Grid[x][y].IsBait = true
	Grid[x][y].Value = 1
}

func genRandomStartPoint() {
	x := rand.Intn(29)
	y := rand.Intn(29)
	Snake.Head[0],Snake.Head[1] = x,y
	Snake.Body = make(map[[2]int]bool)
	Grid[x][y].IsSnakeHead = true
	Grid[x][y].Value = 2
}