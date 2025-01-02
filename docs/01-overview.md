# Overview

**`rovercom` is a collection of [protobuf](https://protobuf.dev/) definitions and their compiled counterparts for various languages. The available packages are used by all official services that run on the ASE Rover.**

There are various types of messages defined in the `rovercom` library, logically grouped together in their respective folder.

- [**service outputs**](https://github.com/VU-ASE/rovercom/tree/main/definitions/outputs) which define many sensor data messages (like camera outputs and RPM readings) and all messages that are sent by one or more official ASE services
- [**debug message wrappers**](https://github.com/VU-ASE/rovercom/tree/main/definitions/debug) which are used for our web-based debugging suite
- [**segmentation messages**](https://github.com/VU-ASE/rovercom/tree/main/definitions/segmentation) which allow to split one byte-encoded message up into multiple smaller messages
- [**tuning messages**](https://github.com/VU-ASE/rovercom/tree/main/definitions/tuning) which are used to tune service Over the Air using our web-based tuning suite

## *wrapper.proto*

In many defintion folders you will see a file called *wrapper.proto*. Because all properties of a protobuf message are optional by design, it is often hard to accurately determine which type of message is received from the binary format alone. Especially in languages that preserve no types at runtime.

Hence, we provide the *wrapper.proto* files which define the actual "wrapper" message that is sent over the wire. Since the wrapper contains all possible values (and acts as a union type), receivers can exhaustively `switch` or `match` on all properties. At all times, only one field shouldd be populated, from which the receiver can deduce the type of message received.
