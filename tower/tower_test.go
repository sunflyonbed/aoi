package tower

import (
	"fmt"
	"math/rand"
	"testing"
)

func NewTestNode(id uint64) *TestNode {
	return &TestNode{
		Id: id,
	}
}

type TestNode struct {
	Id uint64
}

func (this *TestNode) UniqueId() uint64 { return this.Id }
func (this *TestNode) OnEnter(id uint64) {
	fmt.Printf("Node:%d see enter:%d\n", this.Id, id)
}
func (this *TestNode) OnLeave(id uint64) {
	fmt.Printf("Node:%d see leave:%d\n", this.Id, id)
}
func (this *TestNode) OnMove(id uint64) {
	fmt.Printf("Node:%d see move:%d\n", this.Id, id)
}

func (this *TestNode) AOIRange() int {
	return 2
}

/*
	6						2
	5	1
	4
	3					5
	2		4
	1			3
		1	2	3	4	5	6
*/
func TestTower(t *testing.T) {
	scene := NewTowerAOI(0, 0, 6, 6, 2)
	scene.AddNode(NewTestNode(1), 1, 5)
	scene.AddNode(NewTestNode(2), 6, 6)
	scene.AddNode(NewTestNode(3), 3, 1)
	scene.AddNode(NewTestNode(4), 2, 2)
	scene.AddNode(NewTestNode(5), 5, 3)
	scene.PrintAOI()

	scene.MoveNode(5, 6, 3)
	scene.PrintAOI()
	scene.MoveNode(1, 1, 1)
	scene.MoveNode(4, 2, 1)
	scene.PrintAOI()
	scene.LeaveNode(4)
	scene.PrintAOI()

}

const (
	BenchmarkNodeSize = 1000
	BenchmarkMapSize  = 100
)

//go test -test.bench=".*" -count=5
func BenchmarkTower(b *testing.B) {
	scene := NewTowerAOI(0, 0, BenchmarkMapSize, BenchmarkMapSize, 10)
	exist := make(map[uint64]bool)
	for i := 0; i < b.N; i++ {
		id := uint64(i%BenchmarkNodeSize) + 1
		if !exist[id] {
			scene.AddNode(NewTestNode(id), 1, 1)
			continue
		}
		result := rand.Intn(10)
		x := rand.Intn(BenchmarkMapSize)
		y := rand.Intn(BenchmarkMapSize)
		switch result {
		case 0:
			scene.LeaveNode(id)
		default:
			scene.MoveNode(id, x, y)
		}
	}
}
