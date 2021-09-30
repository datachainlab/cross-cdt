package types

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"sort"

	"github.com/gogo/protobuf/proto"
)

var _ CDTState = (*GSetCDTState)(nil)

type Element = []byte

func NewGSetCDTState() *GSetCDTState {
	return &GSetCDTState{}
}

func (s *GSetCDTState) Apply(op OP) error {
	switch op := op.(type) {
	case *GSetOP:
		for _, e := range op.Add {
			_, found := findTargetIndex(s.AddSet, e)
			if !found {
				s.AddSet = append(s.AddSet, e)
			}
		}
		s.sort()
	case *GSetInverseOP:
		// remove elements from the set
		for _, e := range op.Add {
			idx, found := findTargetIndex(s.AddSet, e)
			if !found {
				panic("fatal error")
			}
			s.AddSet = append(s.AddSet[:idx], s.AddSet[idx+1:]...)
		}
	default:
		return fmt.Errorf("cannot apply %T", op)
	}
	return nil
}

func (s *GSetCDTState) IsEmpty() bool {
	return len(s.AddSet) == 0
}

func (s *GSetCDTState) Lookup(elements []Element) bool {
	if len(elements) == 0 {
		panic("the elements length must be greater than 0")
	}
	for _, element := range elements {
		_, found := findTargetIndex(s.AddSet, element)
		if !found {
			return false
		}
	}
	return true
}

func (s *GSetCDTState) sort() {
	sort.SliceStable(s.AddSet, func(i, j int) bool {
		return bytes.Compare(s.AddSet[i], s.AddSet[j]) < 0
	})
}

func findTargetIndex(set []Element, target Element) (int, bool) {
	i := sort.Search(len(set), func(x int) bool {
		return bytes.Compare(set[x], target) >= 0
	})
	if i < len(set) && bytes.Equal(set[i], target) {
		return i, true
	} else {
		return 0, false
	}
}

type GSetOPComposer struct{}

var _ OPComposer = (*GSetOPComposer)(nil)

func (GSetOPComposer) Compose(a []OP, b OP) []OP {
	if l := len(a); l == 0 {
		return []OP{b}
	} else if l != 1 {
		panic("fatal error")
	}

	s := a[0]
	return []OP{s.Compose(b)}
}

func NewGSetValue(elements ...Element) *GSetValue {
	var set GSetValue
	set.Add(elements...)
	return &set
}

func NewGSetValueFromBytes(bz []byte) *GSetValue {
	var set GSetValue
	if err := proto.Unmarshal(bz, &set); err != nil {
		panic(err)
	}
	return &set
}

func (s *GSetValue) Add(elements ...Element) {
	orginal := len(s.Elements)
	elm := make(map[string]struct{})
	for _, element := range elements {
		_, found := findTargetIndex(s.Elements[:orginal], element)
		if !found {
			key := hex.EncodeToString(element)
			_, ok := elm[key]
			if ok {
				continue
			}
			elm[key] = struct{}{}
			s.Elements = append(s.Elements, element)
		}
	}
	sort.SliceStable(s.Elements, func(i, j int) bool {
		return bytes.Compare(s.Elements[i], s.Elements[j]) < 0
	})
}

func (s *GSetValue) Lookup(elements []Element) bool {
	if len(elements) == 0 {
		panic("the elements length must be greater than 0")
	}
	for _, element := range elements {
		_, found := findTargetIndex(s.Elements, element)
		if !found {
			return false
		}
	}
	return true
}

func (s *GSetValue) Bytes() []byte {
	bz, err := proto.Marshal(s)
	if err != nil {
		panic(err)
	}
	return bz
}
