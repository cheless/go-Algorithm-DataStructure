package skiplist

import (
	"fmt"
	"math/rand"
)

type Node struct {
	val int
	right, down *Node
}

func NewNode(v int, r, d *Node) *Node {
	return &Node{v, r, d}
}

type Skiplist struct {
	Level int
	Head *Node
}

func NewSkiplist() *Skiplist {
	return &Skiplist{0, nil}
}

// Add *** 插入level+1层 Head列的情况还需完善
func (sl *Skiplist) Add(num int) {
	if sl.Head == nil {
		sl.Head = &Node{num, nil, nil}
		sl.Level = 1
		return
	}

	// 判断索引加在哪一层
	Level := 1
	for Level <= sl.Level && ((rand.Int31()) & 1 == 0) {
		Level++
	}

	if Level > sl.Level {
		sl.Level = Level
		sl.Head = NewNode(num, nil, sl.Head)
	}
	cur := sl.Head
	var last *Node = nil
	for l := sl.Level; l >= 1; l-- {
		for cur.right != nil && cur.right.val < num {
			cur = cur.right
		}
		if l <= Level {
			cur.right = NewNode(num, cur.right, nil)
			if last != nil {
				last.down = cur.right
			}
			last = cur.right
		}
		cur = cur.down
	}
}

func (sl *Skiplist) Delete(num int) bool {
	cur := sl.Head
	if cur == nil {
		return false
	}
	if cur.val == num {
		if cur.right != nil {
			sl.Head = cur.right
		} else if cur.down != nil {
			sl.Head = cur.down.right
		} else {
			sl.Head = nil
		}
		return true
	}

	// 找到删除位置
	flag := false
	for l := sl.Level; l >= 1; l-- {
		for cur.right != nil && cur.right.val < num {
			cur = cur.right
		}
		if cur.right != nil && cur.right.val == num {
			cur.right = cur.right.right
			flag = true
		}
		cur = cur.down
	}
	return flag
}

func (sl *Skiplist) Search(num int) bool {
	cur := sl.Head
	for cur != nil {
		for cur.right != nil && cur.right.val <= num {
			cur = cur.right
		}
		if cur.val == num {
			return true
		}
		cur = cur.down
	}
	return false
}

func (sl *Skiplist) PrintStruct() {
	cur := sl.Head
	for l := sl.Level; l >= 1; l-- {
		t := cur
		for t != nil {
			fmt.Printf("%d ", t.val)
			t = t.right
		}
		fmt.Println()
		cur = cur.down
	}
	fmt.Println()
}