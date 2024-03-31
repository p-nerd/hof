doc:
	echo "http://localhost:6060"
	godoc -http=:6060
cover:
	go test -coverprofile=cover.out
	go tool cover -html=cover.out
