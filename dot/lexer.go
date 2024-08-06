package dot

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"iter"
	"unicode"

	"github.com/teleivo/skeleton/dot/token"
)

type Lexer struct {
	r    *bufio.Reader
	cur  rune
	next rune
	eof  bool
}

func New(r io.Reader) *Lexer {
	lexer := Lexer{r: bufio.NewReader(r)}
	return &lexer
}

func (l *Lexer) readRune() error {
	r, _, err := l.r.ReadRune()
	if err != nil {
		if !errors.Is(err, io.EOF) {
			return err
		}

		l.eof = true
		l.cur = l.next
		l.next = 0
		return nil
	}

	l.cur = l.next
	l.next = r
	return nil
}

// All returns an iterator over all dot tokens in the given reader.
func (l *Lexer) All() iter.Seq2[token.Token, error] {
	return func(yield func(token.Token, error) bool) {
		// initialize current and next runes
		err := l.readRune()
		fmt.Printf("l.cur %q, l.next %q, err %v\n", l.cur, l.next, err)
		if errors.Is(err, io.EOF) {
			return
		}
		err = l.readRune()
		fmt.Printf("l.cur %q, l.next %q, err %v\n", l.cur, l.next, err)
		if errors.Is(err, io.EOF) {
			return
		}
		fmt.Println("initialized")

		for {
			var tok token.Token

			fmt.Println("before skipWhitespace")
			fmt.Printf("l.cur %q, l.next %q, err %v\n", l.cur, l.next, err)
			err := l.skipWhitespace()
			fmt.Println("after skipWhitespace")
			fmt.Printf("l.cur %q, l.next %q, err %v\n", l.cur, l.next, err)
			if err != nil {
				return
			} else if l.eof && l.cur == 0 {
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
				if l.cur == '-' && (l.next == '>' || l.next == '-') {
					tok, err = l.tokenizeEdgeOperator()
				} else {
					tok, err = l.tokenizeIdentifier()
					if !yield(tok, err) || l.eof {
						return
					}
					continue
				}
			}

			if !yield(tok, err) || l.eof {
				return
			}

			fmt.Println("before advance")
			fmt.Printf("l.cur %q, l.next %q, err %v\n", l.cur, l.next, err)
			err = l.readRune()
			fmt.Println("after advance")
			fmt.Printf("l.cur %q, l.next %q, err %v\n", l.cur, l.next, err)
		}
		// TODO handle illegal runes
		// TODO handle error that is not io.EOF
	}
}

func (l *Lexer) skipWhitespace() (err error) {
	for isWhitespace(l.cur) {
		err := l.readRune()
		if err != nil {
			return err
		}
	}

	return nil
}

func isWhitespace(r rune) bool {
	return unicode.IsSpace(r)
}

func (l *Lexer) tokenizeRuneAs(tokenType token.TokenType) (token.Token, error) {
	return token.Token{Type: tokenType, Literal: string(l.cur)}, nil
}

func (l *Lexer) tokenizeEdgeOperator() (token.Token, error) {
	err := l.readRune()
	if l.cur == '-' {
		return token.Token{Type: token.UndirectedEgde, Literal: token.UndirectedEgde}, err
	}
	return token.Token{Type: token.DirectedEgde, Literal: token.DirectedEgde}, err
}

func (l *Lexer) tokenizeIdentifier() (token.Token, error) {
	// TODO should I move this into a func like isIdentifier and check that in the All() loop and
	// direct to the illegal token case there right away?
	// TODO should these read until whitespace/eof and check every rune for the valid set of runes
	// in that particular category?
	if l.cur == '"' { // double-quoted string
		return l.tokenizeQuotedString()
	} else if l.cur == '<' { // HTML string
		return l.tokenizeHTMLString()
	} else if l.cur == '-' || l.cur == '.' || unicode.IsDigit(l.cur) { // numeral
		return l.tokenizeNumeral()
	} else if isAlphabetic(l.cur) || l.cur == '_' { // any valid string
		return l.tokenizeUnquotedString()
	} else {
		// TODO invalid
		var tok token.Token
		return tok, errors.New("invalid token")
	}
}

func isIdentifier(r rune) bool {
	return isAlphabetic(r) || r == '_' || r == '-' || r == '.' || r == '"' || r == '\\' || unicode.IsDigit(r)
}

func isAlphabetic(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '\200' && r <= '\377')
}

func (l *Lexer) tokenizeQuotedString() (token.Token, error) {
	var tok token.Token
	var err error

	id := []rune{l.cur}
	for err = l.readRune(); err == nil && isIdentifier(l.cur); err = l.readRune() {
		id = append(id, l.cur)
	}

	if err != nil {
		return tok, err
	}

	return token.Token{Type: token.Identifier, Literal: string(id)}, nil
}

func (l *Lexer) tokenizeHTMLString() (token.Token, error) {
	var tok token.Token
	var err error

	id := []rune{l.cur}
	for err = l.readRune(); err == nil && isIdentifier(l.cur); err = l.readRune() {
		id = append(id, l.cur)
	}

	if err != nil {
		return tok, err
	}

	return token.Token{Type: token.Identifier, Literal: string(id)}, nil
}

func (l *Lexer) tokenizeNumeral() (token.Token, error) {
	var tok token.Token
	var err error

	id := []rune{l.cur}
	for err = l.readRune(); err == nil && unicode.IsDigit(l.cur); err = l.readRune() {
		id = append(id, l.cur)
	}

	if err != nil {
		return tok, err
	}

	return token.Token{Type: token.Identifier, Literal: string(id)}, nil
}

// tokenizeUnquotedString considers the current rune(s) as an identifier that might be a dot
// keyword.
func (l *Lexer) tokenizeUnquotedString() (token.Token, error) {
	var tok token.Token
	var err error

	id := []rune{l.cur}
	for err = l.readRune(); err == nil && isIdentifier(l.cur); err = l.readRune() {
		id = append(id, l.cur)
	}

	if err != nil {
		return tok, err
	}

	literal := string(id)
	tok = token.Token{Type: token.LookupIdentifier(literal), Literal: literal}

	return tok, err
}
