package student

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data_ref interface{}) *List {
	newNode := &NodeL{Data: data_ref}
	if l == nil {
		return &List{Head: newNode, Tail: newNode}
	} else if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		return l
	}
	mover := l.Head
	for mover.Next != nil {
		mover = mover.Next
	}
	mover.Next = newNode
	l.Tail = newNode
	return l
}

//
func ListRemoveIf(l *List, data_ref interface{}) {
	if l == nil || l.Head == nil {
		return
	}
	prev := l.Head
	for prev != nil {
		if prev.Data == data_ref {
			l.Head = prev.Next
			prev = l.Head
		} else {
			break
		}
	}
	if l.Head == nil {
		return
	} else if prev.Next == nil {
		l.Tail = prev
		return
	}
	mover := l.Head.Next
	for mover.Next != nil {
		if mover.Data == data_ref {
			prev.Next = mover.Next
			for mover.Next != nil {
				if mover.Data == data_ref {
					prev.Next = mover.Next
					mover = mover.Next
				} else {
					break
				}
			}
			continue
		}
		prev = mover
		mover = mover.Next
	}
	if mover.Data == data_ref {
		l.Tail = prev
		prev.Next = nil
	}
}
