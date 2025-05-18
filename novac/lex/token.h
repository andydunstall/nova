#pragma once

#include <string>

#include "novac/lex/position.h"

namespace novac {
namespace lex {

struct Token {
  // Type defines the lexical token types in Nova.
  enum class Type {
    kEof,

    // Identifiers.
    kIdent,  // foo
    kInt,    // 12345
    kBool,   // true

    // Operators.
    kAdd,  // +
    kSub,  // -
    kMul,  // *
    kQuo,  // /
    kRem,  // %

    kAddAssign,  // +=
    kSubAssign,  // -=
    kMulAssign,  // *=
    kQuoAssign,  // /=
    kRemAssign,  // %=

    kAnd,  // &
    kOr,   // |
    kXor,  // ^
    kShl,  // <<
    kShr,  // >>

    kInc,  // ++
    kDec,  // --

    kLAnd,  // &&
    kLOr,   // ||

    kEql,     // ==
    kNeq,     // !=
    kLss,     // <
    kGtr,     // >
    kAssign,  // =
    kNot,     // !
    kLeq,     // <=
    kGeq,     // >=

    kLParen,     // (
    kLBrace,     // {
    kRParen,     // )
    kRBrace,     // }
    kColon,      // ;
    kSemiColon,  // ;
    kComma,      // ,
    kArrow,      // ->
    kTilde,      // ~

    // Keywords.

    kFn,
    kReturn,

    kMut,
    kLet,

    kIf,
    kElse,

    kLoop,
    kContinue,
    kBreak,
  };

  Type type;
  std::string lit;
  Position pos;
};

std::string TokenTypeToString(Token::Type type);

}  // namespace lex
}  // namespace novac
