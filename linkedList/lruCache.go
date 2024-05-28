package linkedList

import "fmt"

type DLLNode struct {
	val, key   int
	prev, next *DLLNode
}

type LRUCache struct {
	capacity, count int
	lruMap          map[int]*DLLNode
	head, tail      *DLLNode
}

func (l *LRUCache) DeleteNode(node *DLLNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (l *LRUCache) AddNode(node *DLLNode) {
	node.next = l.head.next
	node.next.prev = node
	node.prev = l.head
	l.head.next = node
}

func (l *LRUCache) GetNode(key int) int {
	if _, ok := l.lruMap[key]; ok {
		node := l.lruMap[key]
		result := node.val
		l.DeleteNode(node)
		l.AddNode(node)
		fmt.Printf("\nGot the value : %d for the key %d: ", node.val, result)
		return result
	}
	fmt.Printf("\nDid not get any value for the key: %d: ", key)
	return -1
}

func (l *LRUCache) SetNode(key, value int) {
	fmt.Printf("\ngoing to set the (%d %d)", key, value)
	if _, ok := l.lruMap[key]; ok {
		node := l.lruMap[key]
		node.val = value
		l.DeleteNode(node)
		l.AddNode(node)
		return
	}
	node := &DLLNode{
		val: value,
		key: key,
	}
	l.lruMap[key] = node
	if l.count < l.capacity {
		l.count++
		l.AddNode(node)
	} else {
		delete(l.lruMap, l.tail.prev.key)
		l.DeleteNode(l.tail.prev)
		l.AddNode(node)
	}
}

func LRUCacheDriver() {
	cache := LRUCache{
		capacity: 2,
		lruMap:   make(map[int]*DLLNode),
		head:     &DLLNode{},
		tail:     &DLLNode{},
	}
	cache.head.next = cache.tail
	cache.tail.prev = cache.head

	fmt.Println("Going to test the LRU Cache Implementation")
	cache.SetNode(1, 10)
	cache.SetNode(2, 20)
	fmt.Printf("\nValue for the key: 1 is %d", cache.GetNode(1))
	cache.SetNode(3, 30)
	fmt.Printf("\nValue for the key: 2 is %d", cache.GetNode(2))

	cache.SetNode(4, 40)
	fmt.Printf("\nValue for the key: 1 is %d", cache.GetNode(1))
	fmt.Printf("\nValue for the key: 3 is %d", cache.GetNode(3))
	fmt.Printf("\nValue for the key: 4 is %d\n", cache.GetNode(4))
}
