package tower

import "fmt"

type TowerAOI struct {
	towers     [][]*Tower
	minX, maxX int
	minY, maxY int
	towerRange int
	xNum, yNum int
	nodes      map[uint64]*Node
}

func NewTowerAOI(minX, minY, maxX, maxY, towerRange int) *TowerAOI {
	t := &TowerAOI{
		minX:       minX,
		minY:       minY,
		maxX:       maxX,
		maxY:       maxY,
		towerRange: towerRange,
		nodes:      make(map[uint64]*Node),
	}
	t.xNum = int(maxX-minX)/towerRange + 1
	t.yNum = int(maxY-minY)/towerRange + 1
	fmt.Printf("NewTowerAOI:%d %d %d %d %d result:%d %d\n", minX, maxX, minY, maxY, towerRange, t.xNum, t.yNum)
	t.towers = make([][]*Tower, t.xNum)
	for index, _ := range t.towers {
		t.towers[index] = make([]*Tower, t.yNum)
		for yIndex, _ := range t.towers[index] {
			t.towers[index][yIndex] = NewTower(index, yIndex)
			// fmt.Printf("init Tower:%v\n", t.towers[index][yIndex])
		}
	}
	return t
}

func (this *TowerAOI) GetNode(id uint64) *Node {
	return this.nodes[id]
}

func (this *TowerAOI) AddNode(node MyNode, x, y int) {
	n := NewNode(node, x, y)
	for _, tower := range this.getRangeTower(x, y, n.AOIRange()) {
		tower.AddWatcher(n)
	}
	t := this.getTower(x, y)
	t.AddNode(n, nil)
	this.nodes[n.Id()] = n
}

func (this *TowerAOI) LeaveNode(id uint64) {
	node := this.GetNode(id)
	if node == nil {
		return
	}
	node.Tower().RemoveNode(node, true)
	for _, tower := range this.getRangeTower(node.X(), node.Y(), node.AOIRange()) {
		tower.RemoveWatcher(node)
	}
}

func (this *TowerAOI) MoveNode(id uint64, x, y int) {
	node := this.GetNode(id)
	if node == nil {
		return
	}
	oldX := node.X()
	oldY := node.Y()
	node.SetX(x)
	node.SetY(y)
	oldTower := node.Tower()
	newTower := this.getTower(x, y)
	if oldTower != newTower {
		oldTower.RemoveNode(node, false)
		newTower.AddNode(node, oldTower)
	}

	oldXMin, oldXMax, oldYMin, oldYMax := this.getRangeTowerPos(oldX, oldY, node.AOIRange())
	xMin, xMax, yMin, yMax := this.getRangeTowerPos(x, y, node.AOIRange())

	for xi := oldXMin; xi <= oldXMax; xi++ {
		for yi := oldYMin; yi <= oldYMax; yi++ {
			if xi >= xMin && xi <= xMax && yi >= yMin && yi <= yMax {
				continue
			}
			t := this.towers[xi][yi]
			t.RemoveWatcher(node)
		}
	}

	for xi := xMin; xi <= xMax; xi++ {
		for yi := yMin; yi <= yMax; yi++ {
			if xi >= oldXMin && xi <= oldXMax && yi >= oldYMin && yi <= oldYMax {
				continue
			}
			t := this.towers[xi][yi]
			t.AddWatcher(node)
		}
	}

}

func (this *TowerAOI) getTower(x, y int) *Tower {
	xt, yt := this.transPos(x, y)
	return this.towers[xt][yt]
}

func (this *TowerAOI) getRangeTower(x, y, myRange int) []*Tower {
	xMin, xMax, yMin, yMax := this.getRangeTowerPos(x, y, myRange)
	list := make([]*Tower, 0, (xMax-xMin+1)*(yMax-yMin+1))
	for xi := xMin; xi <= xMax; xi++ {
		for yi := yMin; yi <= yMax; yi++ {
			tower := this.towers[xi][yi]
			list = append(list, tower)
		}
	}
	return list
}

func (this *TowerAOI) getRangeTowerPos(x, y, myRange int) (int, int, int, int) {
	xMin, yMin := this.transPos(x-myRange, y-myRange)
	xMax, yMax := this.transPos(x+myRange, y+myRange)
	return xMin, xMax, yMin, yMax
}

func (this *TowerAOI) transPos(x, y int) (int, int) {
	tx := int(x-this.minX) / this.towerRange
	ty := int(y-this.minY) / this.towerRange
	if tx < 0 {
		tx = 0
	} else if tx >= this.xNum {
		tx = this.xNum - 1
	}
	if ty < 0 {
		ty = 0
	} else if ty >= this.yNum {
		ty = this.yNum - 1
	}
	return tx, ty
}

func (this *TowerAOI) PrintAOI() {
	for _, v := range this.nodes {
		towers := this.getRangeTower(v.X(), v.Y(), v.AOIRange())
		fmt.Printf("printAOI:%d list:", v.Id())
		for _, tower := range towers {
			tower.PrintNodes(v)
		}
		fmt.Printf("\n")
	}
}
