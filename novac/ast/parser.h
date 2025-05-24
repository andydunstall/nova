#pragma once

#include "novac/ast/file.h"
#include "novac/lex/scanner.h"

namespace novac {
namespace ast {

class Parser {
 public:
  Parser(lex::Scanner scanner, bool debug);

  File* Parse();
};

}  // namespace ast
}  // namespace novac
