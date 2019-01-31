package tower

import "fmt"

type Tower struct {
	x, y     int
	nodes    map[uint64]*Node
	watchers map[uint64]*Node
}

func NewTower(x, y int) *Tower {
	return &Tower{
		x:        x,
		y:        y,
		nodes:    make(map[uint64]*Node),
		watchers: make(map[uint64]*Node),
	}
}

func (this *Tower) GetNode(id uint64) *Node    { return this.nodes[id] }
func (this *Tower) GetWatcher(id uint64) *Node { return this.watchers[id] }

func (this *Tower) AddNode(node *Node, fromTower *Tower) {
	node.SetTower(this)
	this.nodes[node.Id()] = node
	// fmt.Printf("Tower:%d %d watchers:%v AddNode :%d\n", this.x, this.y, this.watchers, node.Id())
	if fromTower == nil {
		for _, watcher := range this.watchers {
			if watcher == node {
				continue
			}
			watcher.OnEnter(node.Id())
		}
	} else {
		for _, watcher := range fromTower.watchers {
			if watcher == node {
				continue
			}
			if this.GetWatcher(watcher.Id()) != nil {
				continue
			}
			watcher.OnLeave(node.Id())
		}

		for _, watcher := range this.watchers {
			if watcher == node {
				continue
			}
			if fromTower.GetWatcher(watcher.Id()) != nil {
				continue
			}
			watcher.OnEnter(node.Id())
		}
	}

}

func (this *Tower) RemoveNode(node *Node, notify bool) {
	// fmt.Printf("Tower:%d %d RemoveNode :%d\n", this.x, this.y, node.Id())
	node.SetTower(nil)
	delete(this.nodes, node.Id())
	if notify {
		for _, watcher := range this.watchers {
			if watcher == node {
				continue
			}
			watcher.OnLeave(node.Id())
		}
	}
}

func (this *Tower) AddWatcher(node *Node) {
	if this.GetWatcher(node.Id()) != nil {
		return
	}
	// fmt.Printf("Tower:%d %d AddWatcher :%d\n", this.x, this.y, node.Id())
	this.watchers[node.Id()] = node
	for id, v := range this.nodes {
		if v == node {
			continue
		}
		node.OnEnter(id)
	}
}

func (this *Tower) RemoveWatcher(node *Node) {
	if this.GetWatcher(node.Id()) == nil {
		return
	}
	// fmt.Printf("Tower:%d %d RemoveWatcher :%d\n", this.x, this.y, node.Id())
	delete(this.watchers, node.Id())
	for id, v := range this.nodes {
		if v == node {
			continue
		}
		node.OnLeave(id)
	}
}

func (this *Tower) PrintNodes(node *Node) {
	for _, v := range this.nodes {
		if v == node {
			continue
		}
		fmt.Printf(" %d", v.Id())
	}
}
