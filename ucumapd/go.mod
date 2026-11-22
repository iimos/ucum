module github.com/iimos/ucum/ucumapd

go 1.21

require (
	github.com/cockroachdb/apd/v3 v3.2.1
	github.com/iimos/ucum v0.0.2
)

replace github.com/iimos/ucum => ../ // local dev: test against the working tree
