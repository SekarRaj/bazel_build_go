# Bazel Golang Build

```sh
$ bazel run //:gazelle

$ bazel run //packages/hello-world:hello-world

$ bazel build //packages/hello-world:hello-world

$ bazel-bin/packages/hello-world/darwin_amd64_stripped/hello-world

$ ./bazel-bin/packages/hello-world/darwin_amd64_stripped/hello-world

```