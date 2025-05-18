#include "novac/lex/token.h"

namespace novac {
namespace lex {

std::string TokenTypeToString(Token::Type type) {
  switch (type) {
    case Token::Type::kEof:
      return "EOF";

    case Token::Type::kIdent:
      return "IDENT";
    case Token::Type::kInt:
      return "INT";
    case Token::Type::kBool:
      return "BOOL";

    case Token::Type::kAdd:
      return "+";
    case Token::Type::kSub:
      return "-";
    case Token::Type::kMul:
      return "*";
    case Token::Type::kQuo:
      return "/";
    case Token::Type::kRem:
      return "%";

    case Token::Type::kAddAssign:
      return "+=";
    case Token::Type::kSubAssign:
      return "-=";
    case Token::Type::kMulAssign:
      return "*=";
    case Token::Type::kQuoAssign:
      return "/=";
    case Token::Type::kRemAssign:
      return "%=";

    case Token::Type::kAnd:
      return "&";
    case Token::Type::kOr:
      return "|";
    case Token::Type::kXor:
      return "&";
    case Token::Type::kShl:
      return "<<";
    case Token::Type::kShr:
      return ">>";

    case Token::Type::kInc:
      return "++";
    case Token::Type::kDec:
      return "--";

    case Token::Type::kLAnd:
      return "&&";
    case Token::Type::kLOr:
      return "||";

    case Token::Type::kEql:
      return "==";
    case Token::Type::kNeq:
      return "!=";
    case Token::Type::kLss:
      return "<";
    case Token::Type::kGtr:
      return ">";
    case Token::Type::kAssign:
      return "=";
    case Token::Type::kNot:
      return "!";
    case Token::Type::kLeq:
      return "<=";
    case Token::Type::kGeq:
      return ">=";

    case Token::Type::kLParen:
      return "(";
    case Token::Type::kLBrace:
      return "{";
    case Token::Type::kRParen:
      return ")";
    case Token::Type::kRBrace:
      return "}";
    case Token::Type::kColon:
      return ":";
    case Token::Type::kSemiColon:
      return ";";
    case Token::Type::kComma:
      return ",";
    case Token::Type::kArrow:
      return "->";
    case Token::Type::kTilde:
      return "~";

    case Token::Type::kFn:
      return "fn";
    case Token::Type::kReturn:
      return "return";

    case Token::Type::kMut:
      return "mut";
    case Token::Type::kLet:
      return "let";

    case Token::Type::kIf:
      return "if";
    case Token::Type::kElse:
      return "else";

    case Token::Type::kLoop:
      return "loop";
    case Token::Type::kContinue:
      return "continue";
    case Token::Type::kBreak:
      return "break";

    default:
      return "UNKNOWN";
  }
}

}  // namespace lex
}  // namespace novac
