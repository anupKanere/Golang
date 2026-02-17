package main

import "fmt"

type Pair struct {
	key   string
	value int
	next  *Pair
}

type HashTable struct {
	buckets []*Pair
	size    int
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		buckets: make([]*Pair, size),
		size:    size,
	}
}

func (ht *HashTable) hash(key string) int {
	hash := 0
	for _, ch := range key {
		hash += int(ch)
	}
	return hash % ht.size
}

func (ht *HashTable) Put(key string, val int) {
	index := ht.hash(key)
	fmt.Println("Hash value := ", index)
	head := ht.buckets[index]

	for head != nil {
		if head.key == key {
			head.value = val
			return
		}
		head = head.next
	}

	newPair := &Pair{key: key, value: val, next: ht.buckets[index]}
	ht.buckets[index] = newPair
}

func (ht *HashTable) Get(key string) (int, bool) {

	index := ht.hash(key)

	head := ht.buckets[index]

	for head != nil {
		if head.key == key {
			return head.value, true
		}
		head = head.next
	}

	return 0, false
}

func (ht *HashTable) Remove(key string) {
	index := ht.hash(key)
	head := ht.buckets[index]
	var prev *Pair

	for head != nil {
		if head.key == key {
			if prev == nil {
				ht.buckets[index] = head.next
			} else {
				prev.next = head.next
			}
		}
		prev = head
		head = head.next
	}

}

func main() {
	ht := NewHashTable(10)
	ht.Put("monday", 1)
	ht.Put("Tuesday", 2)
	ht.Put("wednesday", 3)
	ht.Put("d", 6)
	ht.Put("dd", 8)
	ht.Remove("wednesday")
	val, found := ht.Get("dd")
	fmt.Println(val, found)
}
