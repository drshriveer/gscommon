module github.com/drshriveer/gtools/gerrors

go 1.21.0

require (
	github.com/drshriveer/gtools/gencommon v0.0.0
	github.com/drshriveer/gtools/genum v0.0.0-20230915011350-8d283eb04d19
	github.com/itzg/go-flagsfiller v1.12.0
	github.com/stretchr/testify v1.8.4
)

replace github.com/drshriveer/gtools/gencommon v0.0.0 => ../gencommon

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/iancoleman/strcase v0.3.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
