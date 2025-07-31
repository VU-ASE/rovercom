import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

# Installation

On every release, all `rovercom` definitions are transpiled to Go, C, Python and TypeScript files.

<Tabs groupId="roverlib-flavor">
  <TabItem value="roverlib-go" label="Go" default>

You can install `rovercom` as a Go package for your Go module as follows:

```bash
# Install latest version
go get github.com/VU-ASE/rovercom/v2
# Install a specific version
go get github.com/VU-ASE/rovercom/v2@v2.0.0
# LEGACY: Install a v1 version
go get github.com/VU-ASE/rovercom@v1.5.1
```

Additionally you will need to install the general `proto` package:

```bash
go get google.golang.org/protobuf
```

  </TabItem>
  <TabItem value="roverlib-python" label="Python">

You can find the transpiled Python files in our repository [here](https://github.com/VU-ASE/rovercom/tree/main/packages/python/gen).

Additionally, you will need to install the `betterproto` package:

```bash
pip install betterproto
```

  </TabItem>
  <TabItem value="roverlib-c" label="C">

You can find the transpiled C header and source files in our repository [here](https://github.com/VU-ASE/rovercom/tree/main/packages/c/gen). Import the header files and make sure to include the C source files when compiling.
  
  </TabItem>
  <TabItem value="roverlib-typescript" label="TypeScript">

You can install `rovercom` as an NPM package for your TypeScript project as follows:

```bash
# Install latest version
npm install https://gitpkg.vercel.app/VU-ASE/rovercom/packages/typescript
# Install a specific version
npm install https://gitpkg.vercel.app/VU-ASE/rovercom/packages/typescript?v1.5.1
```

After installation, the package will be added as `ase-rovercom`.

  </TabItem>

  <TabItem value="other" label="Other">

You will need to manually compile the given definitions using the `protoc` compiler and its plugins for your language of choice. We recommend taking a look at our [compilation targets](https://github.com/VU-ASE/rovercom/blob/main/Makefile) and [Devcontainer setup](https://github.com/VU-ASE/rovercom/blob/main/.devcontainer/Dockerfile) to understand which tools you need.

</TabItem>

</Tabs>



