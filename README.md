Use bazel to build this repo

```
bazel run //:gazelle
bazel test leetcode:all
```

Using with spacemacs

```
bazel build :gopath
GOPATH=bazel-bin/gopath /Applications/Emacs.app/Contents/MacOS/Emacs .
```
