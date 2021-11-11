module github.com/datachainlab/cross-cdt/modules/erc20

go 1.16

require (
	github.com/cosmos/cosmos-sdk v0.43.0-beta1
	github.com/datachainlab/cross v0.2.2
	github.com/datachainlab/cross-cdt v0.0.0-00010101000000-000000000000
	github.com/gogo/protobuf v1.3.3
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/spf13/cobra v1.1.3
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.10
)

replace (
	github.com/datachainlab/cross-cdt => ../../
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
)
