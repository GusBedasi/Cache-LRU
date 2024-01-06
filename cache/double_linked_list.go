package cache

type linkedList struct {
	Head *node
	Tail *node
}

func (ll *linkedList) insert(newNode *node) {
	if ll.Head == nil {
		ll.Head = newNode
		ll.Tail = newNode
	} else {
		current := ll.Head
		for current.Next != nil {
			current = current.Next
		}

		current.Next = newNode
		newNode.Prev = current
		ll.Tail = newNode
	}
}

func (ll *linkedList) remove(n *node) {
	if ll.Head == n {
		target := ll.Head
		ll.Head = target.Next
		ll.Head.Prev = nil

		target.Prev = nil
		target.Next = nil

	} else {
		current := ll.Head
		for current.Next != n {
			current = current.Next
		}

		target := current.Next
		current.Next = target.Next

		if current.Next != nil {
			current.Next.Prev = nil
		}

		target.Prev = nil
		target.Next = nil
	}
}
