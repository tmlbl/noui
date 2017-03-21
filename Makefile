GO_SRC=**.go

bin/noui: $(GO_SRC)
	go build -o bin/noui bin/noui.go
