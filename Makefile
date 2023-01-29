ifeq ($(GOHOSTOS), windows)
	Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
else
	API_PROTO_FILES=$(shell find api -name *.proto)
endif


protoc:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		$(API_PROTO_FILES)