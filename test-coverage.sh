go test -coverpkg=./... -coverprofile=cover.out .
go tool cover -html=cover.out -o coverage.html
open coverage.html
