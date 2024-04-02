load("@rules_foreign_cc//foreign_cc:defs.bzl", "cmake")

filegroup(
    name = "all_srcs",
    srcs = glob(["**"]),
)

cmake(
    name = "libde265",
    cache_entries = {
        "CMAKE_POSITION_INDEPENDENT_CODE": "ON",
        "BUILD_SHARED_LIBS": "OFF",
        "ENABLE_SDL": "OFF",
    },
    lib_source = ":all_srcs",
    out_static_libs = ["libde265.a"],
    visibility = ["//visibility:public"],
)
