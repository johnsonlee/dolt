package types

import (
	"testing"

	"github.com/attic-labs/noms/Godeps/_workspace/src/github.com/stretchr/testify/assert"
	"github.com/attic-labs/noms/chunks"
	"github.com/attic-labs/noms/ref"
)

func TestMeta(t *testing.T) {
	assert := assert.New(t)

	cs := chunks.NewMemoryStore()

	flatList := []Value{UInt32(0), UInt32(1), UInt32(2), UInt32(3), UInt32(4), UInt32(5), UInt32(6), UInt32(7)}
	l0 := NewList(flatList[0])
	lr0 := WriteValue(l0, cs)
	l1 := NewList(flatList[1])
	lr1 := WriteValue(l1, cs)
	l2 := NewList(flatList[2])
	lr2 := WriteValue(l2, cs)
	l3 := NewList(flatList[3])
	lr3 := WriteValue(l3, cs)
	l4 := NewList(flatList[4])
	lr4 := WriteValue(l4, cs)
	l5 := NewList(flatList[5])
	lr5 := WriteValue(l5, cs)
	l6 := NewList(flatList[6])
	lr6 := WriteValue(l6, cs)
	l7 := NewList(flatList[7])
	lr7 := WriteValue(l7, cs)

	mtr := MakeCompoundType(MetaSequenceKind, l0.Type())

	m0 := compoundList{metaSequenceObject{metaSequenceData{{lr0, UInt64(1)}, {lr1, UInt64(2)}}, mtr}, &ref.Ref{}, cs}
	lm0 := WriteValue(m0, cs)
	m1 := compoundList{metaSequenceObject{metaSequenceData{{lr2, UInt64(1)}, {lr3, UInt64(2)}}, mtr}, &ref.Ref{}, cs}
	lm1 := WriteValue(m1, cs)
	m2 := compoundList{metaSequenceObject{metaSequenceData{{lr4, UInt64(1)}, {lr5, UInt64(2)}}, mtr}, &ref.Ref{}, cs}
	lm2 := WriteValue(m2, cs)
	m3 := compoundList{metaSequenceObject{metaSequenceData{{lr6, UInt64(1)}, {lr7, UInt64(2)}}, mtr}, &ref.Ref{}, cs}
	lm3 := WriteValue(m3, cs)

	m00 := compoundList{metaSequenceObject{metaSequenceData{{lm0, UInt64(2)}, {lm1, UInt64(4)}}, mtr}, &ref.Ref{}, cs}
	lm00 := WriteValue(m00, cs)
	m01 := compoundList{metaSequenceObject{metaSequenceData{{lm2, UInt64(2)}, {lm3, UInt64(4)}}, mtr}, &ref.Ref{}, cs}
	lm01 := WriteValue(m01, cs)

	rootList := compoundList{metaSequenceObject{metaSequenceData{{lm00, UInt64(4)}, {lm01, UInt64(8)}}, mtr}, &ref.Ref{}, cs}
	rootRef := WriteValue(rootList, cs)

	rootList = ReadValue(rootRef, cs).(compoundList)

	rootList.IterAll(func(v Value, index uint64) {
		assert.Equal(flatList[index], v)
	})
}
