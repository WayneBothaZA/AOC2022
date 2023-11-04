type Node struct {
	prev *Node
	next *Node
	key  interface{}
}

// return 0 if n = m, -1 or n < m and 1 if n > m
func (n *Node) Compare(m *Node) int {
	if n.y > m.y {
		return 1
	} else if n.y < m.y {
		return -1
	} else {
		if n.x > m.x {
			return 1
		} else if n.x < m.x {
			return -1
		}

	}

}

type List struct {
	head *Node
	tail *Node
}

func (L *List) Insert(key interface{}) {
	node := &Node{
		next: L.head,
		key:  key,
	}
	if L.head != nil {
		L.head.prev = node
	} else {
		L.tail = node
	}
	L.head = node
}

func (L *List) InsertSorted(key interface{}) {
	node := &Node{
		next: L.head,
		key:  key,
	}
	if L.head != nil {
		L.head.prev = node
	} else {
		L.tail = node
	}
	L.head = node
}

func (L *List) Display() {
	fmt.Printf("head = %v, tail = %v\n", L.head.key, L.tail.key)
	node := L.head
	for node != nil {
		fmt.Printf("%+v", node.key)
		if node != L.tail {
			fmt.Printf(" -> ")
		}
		node = node.next
	}
	fmt.Println()
}

func (L *List) Sort() {

}

func Display(node *Node) {
	for node != nil {
		fmt.Printf("%v ->", node.key)
		node = node.next
	}
	fmt.Println()
}
