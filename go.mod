module github.com/Eric-Domeier/stigctl

go 1.23.0

toolchain go1.24.7

require github.com/spf13/cobra v1.10.1

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.10 // indirect
)

replace github.com/Eric-Domeier/stigctl/src/cmd/hello => ./src/cmd/hello

replace github.com/Eric-Domeier/stigctl/src/cmd/version => ./src/cmd/version
