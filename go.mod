module github.com/kwilteam/kuneiform

go 1.19

require (
	github.com/kwilteam/kwil-db v0.2.1-0.20230530200111-0196b0b65b9c
	github.com/spf13/cobra v1.7.0
)

require (
	github.com/antlr/antlr4/runtime/Go/antlr/v4 v4.0.0-20230512164433-5d1fd1a340c9 // indirect
	github.com/cstockton/go-conv v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/kwilteam/kwil-db/pkg/sql_parser/sql_grammar v0.0.0-00010101000000-000000000000 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/exp v0.0.0-20230419192730-864b3d6c5c2c // indirect
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	// target folder needs a go.mod file, ref: https://github.com/golang/go/wiki/Modules#when-should-i-use-the-replace-directive
	github.com/kwilteam/kwil-db/pkg/sql_parser/sql_grammar => ./sql_grammar
)
