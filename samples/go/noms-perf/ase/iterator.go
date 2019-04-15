package ase

// SortedEditItr is a KVPIterator implementation that takes two KVPCollItr and merges them as it iterates
type SortedEditItr struct {
	leftItr  *KVPCollItr
	rightItr *KVPCollItr
	done     bool
}

// NewSortedEditItr creates an iterator from two KVPCollection references.  As the iterator iterates it
// merges the collections and iterates in order
func NewSortedEditItr(left, right *KVPCollection) *SortedEditItr {
	leftItr := NewItr(left)
	rightItr := NewItr(right)

	return &SortedEditItr{leftItr, rightItr, false}
}

// Next returns the next KVP
func (itr *SortedEditItr) Next() *KVP {
	if itr.done {
		return nil
	}

	lesser := itr.rightItr
	if itr.leftItr.Less(itr.rightItr) {
		lesser = itr.leftItr
	}

	kvp := lesser.Next()

	itr.done = kvp == nil
	return kvp
}

// Size returns the total number of elements this iterator will iterate over.
func (itr *SortedEditItr) Size() int64 {
	return itr.leftItr.coll.totalSize + itr.rightItr.coll.totalSize
}