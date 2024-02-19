package structs

type Cell struct {
	IsVisible   bool
	IsBait      bool
	IsSnakeHead bool
	IsSnakeBody bool
	Value       rune
}
type SnakeBody struct {
	Head [2]int
	Body Node
	HeadNode Node
	BodySize int;
}


type Node struct {
	Value [2]int
	Next *Node
}


