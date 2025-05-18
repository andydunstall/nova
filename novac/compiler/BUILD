cc_library(
    name = "compiler",
    hdrs = glob(["*.h"]),
    srcs = glob(["*.cc"], exclude=["*_test.cc", "*_bench.cc"]),
    visibility = ["//visibility:public"],
    deps = [
        "//novac/lex",
    ],
)
