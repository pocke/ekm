build:
	go get github.com/tools/godep
	godep save
	godep go build
