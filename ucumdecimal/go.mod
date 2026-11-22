module github.com/iimos/ucum/ucumdecimal

go 1.21

require (
	github.com/iimos/ucum v0.0.4
	github.com/shopspring/decimal v1.4.0
)

replace github.com/iimos/ucum => ../ // local dev: test against the working tree
