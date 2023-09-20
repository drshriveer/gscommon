module github.com/drshriveer/gtools/genum

go 1.21.1

require (
	github.com/drshriveer/gtools/gencommon v0.0.0
	github.com/drshriveer/gtools/rutils v0.0.0
	github.com/drshriveer/gtools/set v0.0.0
	github.com/itzg/go-flagsfiller v1.12.0
	github.com/stretchr/testify v1.8.4
	gopkg.in/yaml.v3 v3.0.1
)

replace (
	github.com/drshriveer/gtools/gencommon v0.0.0 => ../gencommon
	github.com/drshriveer/gtools/rutils v0.0.0 => ../rutils
	github.com/drshriveer/gtools/set v0.0.0 => ../set
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/iancoleman/strcase v0.3.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
)
