module github.com/datachainlab/cross-cdt

go 1.17

require (
	github.com/armon/go-radix v1.0.0
	github.com/cosmos/cosmos-sdk v0.43.0-beta1
	github.com/datachainlab/cross v0.2.0
	github.com/gogo/protobuf v1.3.3
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.10
	github.com/tendermint/tm-db v0.6.4
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
