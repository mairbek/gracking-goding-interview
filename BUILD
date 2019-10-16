load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

gazelle(
    name = "gazelle",
    prefix = "github.com/mairbek/gracking-goding-interview",
)

go_library(
    name = "go_default_library",
    srcs = [
        "change.go",
        "lists.go",
        "main.go",
    ],
    importpath = "github.com/mairbek/gracking-goding-interview",
    visibility = ["//visibility:private"],
)

go_binary(
    name = "gracking-goding-interview",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)

load("@io_bazel_rules_go//go:def.bzl", "go_path", "nogo")

go_path(
    name = "gopath",
    mode = "link",
    deps = [
      ":go_default_library",
      ],
)

