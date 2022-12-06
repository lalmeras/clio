module github.com/lalmeras/clio/pkg

go 1.19

replace github.com/lalmeras/clio/introspect => ../introspect

require (
	github.com/lalmeras/clio/introspect v0.0.0-00010101000000-000000000000
	github.com/mitchellh/go-homedir v1.1.0
	github.com/ovh/go-ovh v1.3.0
	github.com/spf13/cobra v1.6.1
	golang.org/x/exp v0.0.0-20221205204356-47842c84f3db
	golang.org/x/sync v0.1.0
	gopkg.in/ini.v1 v1.67.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)
