.PHONY: build start

INPUTS = -I./definitions definitions/**/*.proto

package-go-install-deps:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

package-go-clean:
	rm -rf ./packages/go

package-go-build: package-go-install-deps package-go-clean
	@mkdir -p ./packages/go
	@protoc --go_out=paths=source_relative:./packages/go $(INPUTS)

package-ts-install-deps:
	@mkdir -p ./packages/typescript/gen
	@cd ./packages/typescript && npm install ts-proto

package-ts-clean:
	rm -rf ./packages/typescript/gen

package-ts-build: package-ts-install-deps package-ts-clean
	@mkdir -p ./packages/typescript/gen
	@protoc --plugin=./packages/typescript/node_modules/.bin/protoc-gen-ts_proto --ts_proto_out=paths=source_relative:./packages/typescript/gen --ts_proto_opt=esModuleInterop=true $(INPUTS)

package-c-install-deps:
	@apt install -y protobuf-c-compiler

package-c-clean:
	rm -rf ./packages/c/gen

package-c-build: package-c-install-deps package-c-clean
	@mkdir -p ./packages/c/gen
	@protoc --c_out=./packages/c/gen $(INPUTS)

package-python-clean:
	rm -rf ./packages/python/gen

package-python-build: package-c-clean
	@mkdir -p ./packages/python/gen
	@protoc --python_betterproto_out=./packages/python/gen $(INPUTS)

package-csharp-clean:
	rm -rf ./packages/csharp/gen

package-csharp-build: package-csharp-clean
	@mkdir -p ./packages/csharp/gen
	@protoc --csharp_out=./packages/csharp/gen $(INPUTS)

clean: package-go-clean package-ts-clean package-c-clean package-python-clean package-csharp-clean

build: package-go-build package-ts-build package-c-build package-python-build package-csharp-build
