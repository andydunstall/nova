#include "fmt/core.h"
#include "novac/compiler/compiler.h"
#include "novac/lex/scanner.h"

int main() {
  std::string src = R"(fn main() -> i32 {
	return 10;
})";

  novac::compiler::Compiler compiler{true};
  try {
    compiler.Compile(src);
  } catch (const novac::lex::Exception& e) {
    fmt::println("lex error: {}", e.what());
    exit(EXIT_FAILURE);
  } catch (const std::exception& e) {
    fmt::println("unexpected error: {}", e.what());
    exit(EXIT_FAILURE);
  }
}
