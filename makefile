test:
	go test ./... -cover

bench:
	go test ./... -bench=.
