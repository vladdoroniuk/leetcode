package lru_cache

type Node struct {
	Key   int
	Value int
	Prev  *Node
	Next  *Node
}

type LRUCache struct {
	capacity int
	head     *Node
	tail     *Node
	cache    map[int]*Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	node, exists := this.cache[key]
	if !exists {
		return -1
	}

	this.removeNode(node)
	this.addNode(node)

	return node.Value
}

func (this *LRUCache) Put(key int, value int) {
	node, exists := this.cache[key]

	if exists {
		node.Value = value
		this.removeNode(node)
	} else {
		node = &Node{
			Value: value,
			Key:   key,
		}
		if this.capacity == 0 {
			evictKey := this.head.Key
			this.removeNode(this.head)
			delete(this.cache, evictKey)
		} else {
			this.capacity--
		}
		this.cache[key] = node
	}

	this.addNode(node)
}

func (this *LRUCache) removeNode(n *Node) {
	prev := n.Prev
	next := n.Next

	if prev != nil {
		prev.Next = next
		if this.tail == n {
			this.tail = prev
		}
	}
	if next != nil {
		next.Prev = prev
		if this.head == n {
			this.head = next
		}
	}
	if prev == nil && next == nil {
		this.head = nil
		this.tail = nil
	}

	n.Prev = nil
	n.Next = nil
}

func (this *LRUCache) addNode(n *Node) {
	if this.head == nil {
		this.head = n
		this.tail = n
	} else {
		n.Prev = this.tail
		this.tail.Next = n
		this.tail = n
	}
}
