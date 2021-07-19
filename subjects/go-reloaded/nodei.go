package student

type NodeI struct {
	Data int
	Next *NodeI
}

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	if l == nil {
		return &NodeI{Data: data_ref}
	} else if l.Data >= data_ref {
		return &NodeI{Data: data_ref, Next: l}
	}
	prev := l
	mover := l.Next
	for mover != nil {
		if mover.Data >= data_ref {
			prev.Next = &NodeI{Data: data_ref, Next: mover}
			return l
		}
		prev = mover
		mover = mover.Next
	}
	prev.Next = &NodeI{Data: data_ref}
	return l
}

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	if !IsListSorted(n1) {
		n1 = SlowSortList(n1)
	}
	if !IsListSorted(n2) {
		n2 = SlowSortList(n2)
	}
	if n1 == nil {
		return n2
	} else if n2 == nil {
		return n1
	}
	root := &NodeI{}
	if n1.Data < n2.Data {
		root = n1
		n1 = n1.Next
	} else {
		root = n2
		n2 = n2.Next
	}
	mover := root
	for n1 != nil || n2 != nil {
		if n1 == nil {
			mover.Next = n2
			break
		} else if n2 == nil {
			mover.Next = n1
			break
		}
		if n1.Data < n2.Data {
			mover.Next = n1
			n1 = n1.Next
		} else {
			mover.Next = n2
			n2 = n2.Next
		}
		mover = mover.Next
	}
	return root
}

func IsListSorted(n *NodeI) bool {
	if n == nil {
		return true
	}
	prev := n
	mover := n.Next
	for mover != nil {
		if mover.Data < prev.Data {
			return false
		}
		prev = mover
		mover = mover.Next
	}
	return true
}

func SlowSortList(n *NodeI) *NodeI {
	if n == nil {
		return nil
	}
	tNode := SortListInsert(nil, n.Data)
	n = n.Next
	for n != nil {
		tNode = SortListInsert(tNode, n.Data)
		n = n.Next
	}
	return tNode
}
