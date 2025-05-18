#pragma once

#include <string>

namespace novac {
namespace compiler {

class Compiler {
 public:
  Compiler(bool debug);

  // Compiles the given source and outputs x86-64 assembly.
  std::string Compile(const std::string& src);

 private:
  bool debug_;
};

}  // namespace compiler
}  // namespace novac
