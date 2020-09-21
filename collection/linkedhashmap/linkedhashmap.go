package linkedhashmap

type node struct {
	prev *node
	next *node
	key interface{}
}

type LinkedHashMap struct {
	hash map[interface{}]interface{}
	head *node
	tail *node
	nodes map[interface{}]*node
}

func New() LinkedHashMap {
	return LinkedHashMap{nodes: make(map[interface{}]*node), hash: make(map[interface{}]interface{})}
}

func (m *LinkedHashMap) Put(k, v interface{}) {
	_, exists := m.hash[k]
	if !exists {
		n := newNode(k)
		if m.tail == nil {
			m.tail = n
			m.head = n
		} else {
			m.tail.next = n
			n.prev = m.tail
			m.tail = n
		}
		m.nodes[k] = n
	}
	m.hash[k] = v
}

func (m *LinkedHashMap) Get(k interface{}) (interface{}, bool) {
	v, ok := m.hash[k]
	return v, ok
}

func (m *LinkedHashMap) Remove(k interface{}) (interface{}, bool) {
	v, exists := m.hash[k]
	if !exists {
		return nil, false
	}
	n := m.nodes[k]
	if n == m.tail {
		if m.head == m.tail {
			m.tail = nil
			m.head = nil
		} else {
			m.tail.prev.next = nil
			m.tail = m.tail.prev
		}
	} else if n == m.head {
		m.head.next.prev = nil
		m.head = m.head.next
	} else {
		n.prev.next = n.next
		n.next.prev = n.prev
	}
	delete(m.nodes, k)
	delete(m.hash, k)
	return v, true
}

func (m *LinkedHashMap) Keys() []interface{} {
	var result []interface{}
	t := m.head
	for t != nil {
		result = append(result, t.key)
		t = t.next
	}
	return result
}

func (m *LinkedHashMap) Size() int {
	return len(m.hash)
}

func (m *LinkedHashMap) IsEmpty() bool {
	return len(m.hash) == 0
}

func (m *LinkedHashMap) Each(f func(k, v interface{}) bool)  {
	t := m.head
	for t != nil {
		if !f(t.key, m.hash[t.key]) {
			break
		}
		t = t.next
	}
}

func (m *LinkedHashMap) Clear() {
	m.head = nil
	m.tail = nil
	m.nodes = make(map[interface{}]*node)
	m.hash = make(map[interface{}]interface{})
}

func newNode(k interface{}) *node {
	return &node{key: k}
}

