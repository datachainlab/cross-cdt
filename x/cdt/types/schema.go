package types

import (
	"fmt"

	"github.com/armon/go-radix"
)

type Tree = radix.Tree

type Schema struct {
	tree *Tree
}

func NewSchema() *Schema {
	tree := radix.New()
	return &Schema{tree: tree}
}

func (s *Schema) Set(prefix []byte, tp CDTType) {
	_, ok := s.tree.Insert(string(prefix), tp)
	if ok {
		panic("cannot update the schema")
	}
}

func (s Schema) MatchType(key []byte) (CDTType, bool) {
	_, v, found := s.tree.LongestPrefix(string(key))
	if !found {
		return 0, false
	}
	return v.(CDTType), true
}

func (s Schema) Match(key []byte, tp CDTType) error {
	t, ok := s.MatchType(key)
	if !ok {
		return fmt.Errorf("no type found for the key '%v'", key)
	}
	if t == tp {
		return nil
	}
	return fmt.Errorf("type mismatch: expected=%v got=%v", tp, t)
}

func (s Schema) MatchInt64(key []byte) bool {
	return s.Match(key, CDT_TYPE_INT64) == nil
}

func (s Schema) MustMatchInt64(key []byte) {
	if err := s.Match(key, CDT_TYPE_INT64); err != nil {
		panic(err)
	}
}

func (s Schema) MatchGSet(key []byte) bool {
	return s.Match(key, CDT_TYPE_GSET) == nil
}

func (s Schema) MustMatchGSet(key []byte) {
	if err := s.Match(key, CDT_TYPE_GSET); err != nil {
		panic(err)
	}
}
