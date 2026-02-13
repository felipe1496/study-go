Go Version: 1.22.8

Bazel Version: 7.1.1


generate crds deepcopy:
```shell
bazel run //:kubegen --spawn_strategy=local
```

generate client to crds:
```shell
bazel run //:clientgen --spawn_strategy=local
```
