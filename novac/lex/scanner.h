#pragma once

#include "novac/lex/token.h"

namespace novac {
namespace lex {

// Scanner parses a Nova file and outputs a stream of token literals.
//
// Comments are discarded.
class Scanner {
 public:
  Scanner(const std::string& src, bool debug);

  Token Scan();

 private:
  static constexpr uint8_t kEof = 0xff;

  Position position() const {
    return Position{
        .line = line_,
        .column = column_,
    };
  }

  std::string ScanIdentifier();
  std::string ScanNumber();

  void SkipWhitespace();
  void SkipComments();

  Token::Type Lookup(const std::string& s);

  void Next();

  void PrintToken(Token tok) const;

  std::string src_;

  bool debug_;

  // Current character.
  uint8_t ch_;
  // Current position.
  int line_;
  int column_;
  size_t offset_;
};

class Exception : public std::exception {
 public:
  Exception(const std::string& msg, Position pos);

  const char* what() const noexcept override { return msg_.c_str(); }

 private:
  std::string msg_;
};

}  // namespace lex
}  // namespace novac
