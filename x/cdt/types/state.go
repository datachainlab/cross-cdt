package types

import (
	"fmt"

	"github.com/gogo/protobuf/proto"
)

// CDTState manages the provisonal state
type CDTState interface {
	proto.Message
	Apply(OP) error
	IsEmpty() bool
}

func GetEmptyCDTState(ct CDTType) CDTState {
	switch ct {
	case CDT_TYPE_INT64:
		return NewInt64CDTState()
	case CDT_TYPE_GSET:
		return NewGSetCDTState()
	default:
		panic(fmt.Sprintf("unknown type: %v", ct))
	}
}
