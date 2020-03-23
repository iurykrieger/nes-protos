# 🎮 NES Protos

![Built With](https://img.shields.io/badge/built%20with-protobufs-blue?style=flat-square)
![Built With](https://img.shields.io/badge/repo%20status-active-green?style=flat-square)

This repository centralizes all data structures that are used inside NES Project. It uses [Google Protocol Buffers (a.k.a protobufs)](https://developers.google.com/protocol-buffers) to declare this structures and provide a single way to all microservices to comunicate.

## Dependencies

- [Build Essential](https://packages.debian.org/sid/build-essential)

## Development

Use the `proto3` syntax to develop any data structure of your will. You can check [Google official proto3 language guide](https://developers.google.com/protocol-buffers/docs/proto3) to see more examples.

## Building Protos

Just use make command to build protos inside the docker container.

```bash
make 
# Or
make build
```

## Release

```bash
# Not yet.
```