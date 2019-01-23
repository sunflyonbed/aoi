package olist

import "testing"

/*
	6						2
	5	1
	4
	3					5
	2		4
	1			3
		1	2	3	4	5	6
*/
func TestOlist(t *testing.T) {
	scene := NewMap()
	scene.AddNode(1, 1, 5)
	scene.AddNode(2, 6, 6)
	scene.AddNode(3, 3, 1)
	scene.AddNode(4, 2, 2)
	scene.AddNode(5, 5, 3)
	scene.PrintAOI()

	scene.MoveNode(5, 6, 3)
	scene.MoveNode(1, 1, 1)
	scene.MoveNode(4, 2, 1)
	scene.PrintAOI()
}
