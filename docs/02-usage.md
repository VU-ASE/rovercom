# Installation

## For Go

You can install `rovercom` as a Go package for your Go module as follows:

```bash
# Install latest version
go get github.com/VU-ASE/rovercom
# Install a specific version
go get github.com/VU-ASE/rovercom@v1.5.1
```

Additionally you will need to install the general `proto` package:

```bash
go get google.golang.org/protobuf
```

## For C

You can find the transpiled C header and source files in our repository [here](https://github.com/VU-ASE/rovercom/tree/main/packages/c/gen). Import the header files and make sure to include the C source files when compiling.

## For TypeScirpt

You can install `rovercom` as an NPM package for your TypeScript project as follows:

```bash
# Install latest version
npm install https://gitpkg.vercel.app/VU-ASE/rovercom/packages/typescript
# Install a specific version
npm install https://gitpkg.vercel.app/VU-ASE/rovercom/packages/typescript?v1.5.1
```

After installation, the package will be added as `ase-rovercom`.

## For other languages

You will need to manually compile the given definitions using the `protoc` compiler and its plugins for your language of choice. We recommend taking a look at our [compilation targets](https://github.com/VU-ASE/rovercom/blob/main/Makefile) and [Devcontainer setup](https://github.com/VU-ASE/rovercom/blob/main/.devcontainer/Dockerfile) to understand which tools you need.
