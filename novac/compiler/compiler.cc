#include "novac/compiler/compiler.h"

#include "novac/ast/parser.h"
#include "novac/lex/scanner.h"

namespace novac {
namespace compiler {

Compiler::Compiler(bool debug) : debug_{debug} {}

std::string Compiler::Compile(const std::string& src) {
  lex::Scanner scanner{src, debug_};
  ast::Parser parser{scanner, debug_};
  parser.Parse();

  return "";
}

}  // namespace compiler
}  // namespace novac
