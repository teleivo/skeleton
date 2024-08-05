package dot

import (
	"bufio"
	"errors"
	"io"
	"iter"
	"unicode"

	"github.com/teleivo/skeleton/dot/token"
)

type Lexer struct {
	r   *bufio.Reader
	cur rune
}

func New(r io.Reader) *Lexer {
	lexer := Lexer{r: bufio.NewReader(r)}
	return &lexer
}

func (l *Lexer) readRune() error {
	r, _, err := l.r.ReadRune()
	if err != nil {
		return err
	}

	l.cur = r
	return nil
}

// All returns an iterator over all dot tokens in the given reader.
func (l *Lexer) All() iter.Seq2[token.Token, error] {
	return func(yield func(token.Token, error) bool) {
		for {
			var tok token.Token
			err := l.skipWhitespace()

			if errors.Is(err, io.EOF) {
				return
			}

			switch l.cur {
			case '{':
				tok, err = l.tokenizeRuneAs(token.LeftBrace)
			case '}':
				tok, err = l.tokenizeRuneAs(token.RightBrace)
			case '[':
				tok, err = l.tokenizeRuneAs(token.LeftBracket)
			case ']':
				tok, err = l.tokenizeRuneAs(token.RightBracket)
			case ':':
				tok, err = l.tokenizeRuneAs(token.Colon)
			case ',':
				tok, err = l.tokenizeRuneAs(token.Comma)
			case ';':
				tok, err = l.tokenizeRuneAs(token.Semicolon)
			case '=':
				tok, err = l.tokenizeRuneAs(token.Equal)
			default:
				tok, err = l.tokenizeIdentifier()
			}

			if !yield(tok, err) {
				return
			}
		}
		// TODO handle edge operator which is a two character literal
		// TODO handle illegal runes
		// TODO handle error that is not io.EOF
	}
}

func (l *Lexer) skipWhitespace() (err error) {
	for err = l.readRune(); err == nil && isWhitespace(l.cur); err = l.readRune() {
	}
	return err
}

func isWhitespace(r rune) bool {
	return unicode.IsSpace(r)
}

func (l *Lexer) tokenizeRuneAs(tokenType token.TokenType) (token.Token, error) {
	return token.Token{Type: tokenType, Literal: string(l.cur)}, nil
}

func (l *Lexer) tokenizeIdentifier() (token.Token, error) {
	var tok token.Token
	var err error

	id := []rune{l.cur}
	for err = l.readRune(); err == nil && isIdentifier(l.cur); err = l.readRune() {
		id = append(id, l.cur)
	}

	if err != nil && !errors.Is(err, io.EOF) {
		return tok, err
	}

	literal := string(id)
	tok = token.Token{Type: token.LookupIdentifier(literal), Literal: literal}
	return tok, nil
}

func isIdentifier(r rune) bool {
	return isAlphabetic(r) || r == '_' || r == '-' || r == '.' || r == '"' || r == '\\' || unicode.IsDigit(r)
}

func isAlphabetic(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '\200' && r <= '\377')
}
