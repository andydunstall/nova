#pragma once

namespace novac {
namespace types {

struct Type {};

struct Primative : Type {
  enum class Type {
    kBool,
    kI8,
    kU8,
    kI16,
    kU16,
    kI32,
    kU32,
    kI64,
    kU64,
  };

  Type type;

  // ...
};

}  // namespace types
}  // namespace novac
