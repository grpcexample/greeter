update:
	git submodule update --init --recursive

gen:
	docker run --rm -v ${PWD}:/defs namely/protoc-all -d apis/grpcexample/greeter/v1 -o v1 -l go --go-source-relative
	go mod tidy
