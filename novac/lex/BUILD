cc_library(
    name = "lex",
    hdrs = glob(["*.h"]),
    srcs = glob(["*.cc"], exclude=["*_test.cc", "*_bench.cc"]),
    visibility = ["//visibility:public"],
    deps = [
        "@fmt",
    ],
)
