package link_list

// length = 5
// PUT, GET: ABCDBEFG
// FGEBD

// head <-> ... node ... <-> tail

// 头部插入元素，在尾部删除元素
// 线性数据结构
// PUT：元素要在头部插入，容量满了之后，在尾部删除。----> 双向链表
// GET：在O(1)内查到元素，该元素从原位置删除，再插入头部 ----> map(key,node)

type Node struct {
	key, value int
	prev, next *Node
}

type LRUCache struct {
	capacity   int
	head, tail *Node
	hashMap    map[int]*Node
}

func Constructor(capacity int) LRUCache {
	head, tail := &Node{}, &Node{}
	// head <-> tail
	head.next = tail
	tail.prev = head
	return LRUCache{
		capacity: capacity,
		head:     head,
		tail:     tail,
		hashMap:  map[int]*Node{},
	}
}

func (l *LRUCache) Get(key int) int {
	if node, ok := l.hashMap[key]; ok {
		remove(node)
		insert(l.head, node)
		return node.value
	}
	return -1
}

func (l *LRUCache) Put(key int, value int) {
	if node, ok := l.hashMap[key]; ok {
		// 如果在map里面，属于更新
		node.value = value
		remove(node)
		insert(l.head, node)
	} else {
		newNode := &Node{key: key, value: value}
		insert(l.head, newNode)
		l.hashMap[key] = newNode
		if len(l.hashMap) > l.capacity {
			lastNode := l.tail.prev
			remove(lastNode)
			delete(l.hashMap, lastNode.key)
		}
	}

}

func remove(node *Node) {
	// prev <-> node <-> next
	node.prev.next = node.next // prev -> next
	node.next.prev = node.prev // prev <- next
}

func insert(head, node *Node) {
	// head <-> next
	// head <-> node <-> next
	next := head.next
	node.next = next // node -> next
	next.prev = node // node <- next
	head.next = node // head -> node
	node.prev = head // head <- node
}
