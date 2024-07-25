.PHONY: build start
package-go-install-deps:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

package-go-clean:
	rm -rf ./packages/go

package-go-build: package-go-install-deps package-go-clean
	@mkdir -p ./packages/go
	@protoc --go_out=paths=source_relative:./packages/go -I./definitions definitions/**/*.proto -I./definitions definitions/*.proto

package-ts-install-deps:
	@cd ./packages/typescript && npm install ts-proto

package-ts-clean:
	rm -rf ./packages/typescript/gen

package-ts-build: package-ts-install-deps package-ts-clean
	@mkdir -p ./packages/typescript/gen
	@protoc --plugin=./packages/typescript/node_modules/.bin/protoc-gen-ts_proto --ts_proto_out=paths=source_relative:./packages/typescript/gen --ts_proto_opt=esModuleInterop=true -I./definitions definitions/**/*.proto -I./definitions definitions/*.proto

package-c-install-deps:
	@sudo apt install -y protobuf-c-compiler

package-c-clean:
	rm -rf ./packages/c/gen

package-c-build: package-c-install-deps package-c-clean
	@mkdir -p ./packages/c/gen
	@protoc --c_out=./packages/c/gen -I./definitions definitions/**/*.proto -I./definitions definitions/*.proto

clean: package-go-clean package-ts-clean package-c-clean

build: package-go-build package-ts-build package-c-build
