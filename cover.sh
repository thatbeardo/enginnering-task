clear
go test ./... -coverprofile=coverage.out
echo
go tool cover -func coverage.out
echo
go tool cover -html=coverage.out -o coverage.html
