#pragma once

#include <vector>

#include "novac/ast/decl.h"

namespace novac {
namespace ast {

struct File {
  std::vector<Decl*> decls;
};

}  // namespace ast
}  // namespace novac
