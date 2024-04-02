load("@rules_foreign_cc//foreign_cc:defs.bzl", "cmake")

filegroup(
    name = "all_srcs",
    srcs = glob(["**"]),
)

cmake(
    name = "libheif",
    cache_entries = {
        "CMAKE_POSITION_INDEPENDENT_CODE": "ON",
        "BUILD_SHARED_LIBS": "OFF",
        "LIBDE265_LIBRARY": "$EXT_BUILD_DEPS/libde265/libde265.a",
        "LIBDE265_INCLUDE_DIR": "$EXT_BUILD_DEPS/libde265/include",
        "WITH_LIBDE265": "ON",
        "WITH_GDK_PIXBUF": "OFF",
        "WITH_DAV1D": "OFF",
        "WITH_AOM_ENCODER": "OFF",
        "WITH_AOM_DECODER": "OFF",
        "WITH_X265": "OFF",
        "WITH_KVAZAAR": "OFF",
        "WITH_SvtEnc": "OFF",
        "WITH_RAV1E": "OFF",
        "WITH_JPEG_ENCODER": "OFF",
        "WITH_JPEG_DECODER": "OFF",
        "WITH_OpenJPEG_ENCODER": "OFF",
        "WITH_OpenJPEG_DECODER": "OFF",
        "WITH_FFMPEG_DECODER": "OFF",
        "WITH_LIBSHARPYUV": "OFF",
        "WITH_EXAMPLES": "OFF",
        "BUILD_TESTING": "OFF",
        "ENABLE_PLUGIN_LOADING": "OFF",
    },
    lib_source = ":all_srcs",
    out_static_libs = ["libheif.a"],
    visibility = ["//visibility:public"],
    deps = ["@libde265"],
)
