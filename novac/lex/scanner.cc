#include "novac/lex/scanner.h"

#include "fmt/core.h"

namespace novac {
namespace lex {

namespace {

uint8_t Lower(uint8_t c) { return ('a' - 'A') | c; }

bool IsLetter(uint8_t c) {
  return ('a' <= Lower(c) && Lower(c) <= 'z') || c == '_';
}

bool IsDecimal(uint8_t c) { return '0' <= c && c <= '9'; }

}  // namespace

Scanner::Scanner(const std::string& src, bool debug)
    : src_{src}, debug_{debug}, ch_{kEof}, offset_{0} {
  if (src.size() > 0) {
    ch_ = src[0];
    line_ = 1;
    column_ = 1;
  }
}

Token Scanner::Scan() {
  SkipWhitespace();

  auto pos = position();

  uint8_t ch = ch_;
  if (IsLetter(ch)) {
    auto lit = ScanIdentifier();
    Token tok{.type = Lookup(lit), .lit = lit, .pos = pos};
    PrintToken(tok);
    return tok;
  }
  if (IsDecimal(ch)) {
    auto lit = ScanNumber();
    Token tok{.type = Token::Type::kInt, .lit = lit, .pos = pos};
    PrintToken(tok);
    return tok;
  }

  // Move forward so ch_ points to the next character.
  Next();

  Token::Type type;
  switch (ch) {
    case '+':
      if (ch_ == '+') {
        type = Token::Type::kInc;
        Next();
      } else if (ch_ == '=') {
        type = Token::Type::kAddAssign;
        Next();
      } else {
        type = Token::Type::kAdd;
      }
      break;
    case '-':
      if (ch_ == '-') {
        type = Token::Type::kDec;
        Next();
      } else if (ch_ == '=') {
        type = Token::Type::kSubAssign;
        Next();
      } else if (ch_ == '>') {
        type = Token::Type::kArrow;
        Next();
      } else {
        type = Token::Type::kSub;
      }
      break;
    case '*':
      if (ch_ == '=') {
        type = Token::Type::kMulAssign;
        Next();
      } else {
        type = Token::Type::kMul;
      }
      break;
    case '/':
      if (ch_ == '=') {
        type = Token::Type::kQuoAssign;
        Next();
      } else {
        type = Token::Type::kQuo;
      }
      break;
    case '%':
      if (ch_ == '=') {
        type = Token::Type::kRemAssign;
        Next();
      } else {
        type = Token::Type::kRem;
      }
      break;
    case '&':
      if (ch_ == '&') {
        type = Token::Type::kLAnd;
        Next();
      } else {
        type = Token::Type::kAnd;
      }
      break;
    case '|':
      if (ch_ == '|') {
        type = Token::Type::kLOr;
        Next();
      } else {
        type = Token::Type::kOr;
      }
      break;
    case '^':
      type = Token::Type::kXor;
      break;
    case '=':
      if (ch_ == '=') {
        type = Token::Type::kEql;
        Next();
      } else {
        type = Token::Type::kAssign;
      }
      break;
    case '!':
      if (ch_ == '=') {
        type = Token::Type::kNeq;
        Next();
      } else {
        type = Token::Type::kNot;
      }
      break;
    case '<':
      if (ch_ == '<') {
        type = Token::Type::kShl;
        Next();
      } else if (ch_ == '=') {
        type = Token::Type::kLeq;
        Next();
      } else {
        type = Token::Type::kLss;
      }
      break;
    case '>':
      if (ch_ == '>') {
        type = Token::Type::kShr;
        Next();
      } else if (ch_ == '=') {
        type = Token::Type::kGeq;
        Next();
      } else {
        type = Token::Type::kGtr;
      }
      break;
    case '(':
      type = Token::Type::kLParen;
      break;
    case '{':
      type = Token::Type::kLBrace;
      break;
    case ')':
      type = Token::Type::kRParen;
      break;
    case '}':
      type = Token::Type::kRBrace;
      break;
    case ':':
      type = Token::Type::kColon;
      break;
    case ';':
      type = Token::Type::kSemiColon;
      break;
    case ',':
      type = Token::Type::kComma;
      break;
    case '~':
      type = Token::Type::kTilde;
      break;
    case kEof:
      type = Token::Type::kEof;
      break;
    default:
      throw Exception{fmt::format("unexpected character: {:c}", ch), pos};
  }

  Token tok{.type = type, .pos = pos};
  PrintToken(tok);
  return tok;
}

std::string Scanner::ScanIdentifier() {
  size_t i = offset_;
  for (; i < src_.size(); i++) {
    uint8_t b = src_[i];
    if (('a' <= b && b <= 'z') || ('A' <= b && b <= 'Z') || b == '_' ||
        ('0' <= b && b <= '9')) {
      continue;
    }
    break;
  }

  std::string ident = src_.substr(offset_, i - offset_);
  column_ += (i - offset_);
  offset_ = i;
  ch_ = src_[offset_];
  return ident;
}

std::string Scanner::ScanNumber() {
  size_t i = offset_;
  for (; i < src_.size(); i++) {
    uint8_t b = src_[i];
    if ('0' <= b && b <= '9') {
      continue;
    }

    break;
  }

  std::string ident = src_.substr(offset_, i - offset_);
  column_ += (i - offset_);
  offset_ = i;
  ch_ = src_[offset_];
  return ident;
}

void Scanner::SkipWhitespace() {
  while (true) {
    if (ch_ == ' ' || ch_ == '\t' || ch_ == '\n' || ch_ == '\r') {
      // Whitespace.
      Next();
    } else if (offset_ < src_.size() - 1 && ch_ == '/' &&
               src_[offset_ + 1] == '/') {
      // Comment. Skip to next line.
      while (ch_ != '\n') {
        Next();
      }
    } else {
      break;
    }
  }
}

Token::Type Scanner::Lookup(const std::string& s) {
  if (s == "fn") return Token::Type::kFn;
  if (s == "return") return Token::Type::kReturn;
  if (s == "mut") return Token::Type::kMut;
  if (s == "let") return Token::Type::kLet;
  if (s == "if") return Token::Type::kIf;
  if (s == "else") return Token::Type::kElse;
  if (s == "loop") return Token::Type::kLoop;
  if (s == "continue") return Token::Type::kContinue;
  if (s == "break") return Token::Type::kBreak;
  return Token::Type::kIdent;
}

void Scanner::Next() {
  if (ch_ == '\n') {
    line_++;
    column_ = 0;
  }

  column_++;
  if (offset_ < src_.size() - 1) {
    offset_++;
    ch_ = src_[offset_];
  } else {
    offset_ = src_.size();
    ch_ = kEof;
  }
}

void Scanner::PrintToken(Token tok) const {
  if (!debug_) return;

  auto pos = fmt::format("{}:{}", tok.pos.line, tok.pos.column);

  if (tok.lit == "" || TokenTypeToString(tok.type) == tok.lit) {
    fmt::println("{:6} {}", pos,
                 TokenTypeToString(tok.type));
  } else {
    fmt::println("{:6} {} ({})", pos,
                 TokenTypeToString(tok.type), tok.lit);
  }
}

Exception::Exception(const std::string& msg, Position pos) : msg_{msg} {
  msg_ = fmt::format("{} ({}:{})", msg, pos.line, pos.column);
}

}  // namespace lex
}  // namespace novac
