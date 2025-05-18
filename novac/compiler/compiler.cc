#include "novac/compiler/compiler.h"

#include "novac/lex/scanner.h"

namespace novac {
namespace compiler {

Compiler::Compiler(bool debug) : debug_{debug} {}

std::string Compiler::Compile(const std::string& src) {
  lex::Scanner scanner{src, debug_};
  lex::Token token = scanner.Scan();
  while (token.type != lex::Token::Type::kEof) {
    token = scanner.Scan();
  }

  return "";
}

}  // namespace compiler
}  // namespace novac
