test: fmt
	echo "running tests"
	go test -v -cover ./...
fmt:
	echo "formating codes"
	go vet ./...
	go fmt ./...
doc:
	echo "running docs"
	godoc -http=:6060