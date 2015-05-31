package efabus

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Pos int

type item struct {
	typ itemType
	pos Pos
	val string
}

func (i item) String() string {
	switch {
	case i.typ == itemEOF:
		return "EOF"
	case i.typ == itemError:
		return i.val
	}
	return fmt.Sprintf("%q", i.val)
}

type itemType int

const (
	itemError itemType = iota
	itemEOF
	itemEventName
	itemAttrKey
	itemAttrValue
)

const eof = -1

const (
	leftBracket  = "["
	rightBracket = "]"
	equal        = "="
)

type stateFn func(*lexer) stateFn

// lexer holds the state of the scanner.
type lexer struct {
	name       string    // the name of the input; used only for error reports
	input      string    // the string being scanned
	leftDelim  string    // start of action
	rightDelim string    // end of action
	state      stateFn   // the next lexing function to enter
	pos        Pos       // current position in the input
	start      Pos       // start position of this item
	width      Pos       // width of last rune read from input
	lastPos    Pos       // position of most recent item returned by nextItem
	items      chan item // channel of scanned items
	parenDepth int       // nesting depth of ( ) exprs
}

func (l lexer) Debug() string {
	return fmt.Sprintf("test:%s  pos:%d  start:%d  width:%d  input:%s", l.name, l.pos, l.start, l.width, l.input)
}

// lex creates a new scanner for the input string.
func lex(name, input, left, right string) *lexer {
	if left == "" {
		left = leftBracket
	}
	if right == "" {
		right = rightBracket
	}
	l := &lexer{
		name:       name,
		input:      input,
		leftDelim:  left,
		rightDelim: right,
		items:      make(chan item),
	}
	go l.run()
	return l
}

// run runs the state machine for the lexer.
func (l *lexer) run() {
	for l.state = lexStart; l.state != nil; {
		l.state = l.state(l)
	}
}

// next returns the next rune in the input.
func (l *lexer) next() rune {
	if int(l.pos) >= len(l.input) {
		l.width = 0
		return eof
	}
	r, w := utf8.DecodeRuneInString(l.input[l.pos:])
	l.width = Pos(w)
	l.pos += l.width
	return r
}

// backup steps back one rune. Can only be called once per call of next.
func (l *lexer) backup() {
	l.pos -= l.width
}

// nextItem returns the next item from the input.
func (l *lexer) nextItem() item {
	item := <-l.items
	l.lastPos = item.pos
	return item
}

// lexText scans until an opening action delimiter, "[".
func lexStart(l *lexer) stateFn {

	if strings.HasPrefix(l.input[l.pos:], leftBracket) {
		l.next()
		l.ignore()
		return lexEventName
	}
	l.emit(itemError)
	return nil
}

// lexIdentifier scans an alphanumeric.
func lexEventName(l *lexer) stateFn {
Loop:
	for {
		switch r := l.next(); {
		case isAlphaNumeric(r):
		// absorve
		case r == '=':
			l.backup()
			l.emit(itemEventName)

			// ignore equal
			l.next()
			l.ignore()
			break Loop
		default:
			return l.errorf("Error on EventName! Expecting '=', not %#U", r)
		}
	}
	return lexEventAttrKey
}

func lexEventAttrKey(l *lexer) stateFn {

	r := l.next()

	if r != '[' && r != ']' {
		return l.errorf("Error on Attr Key! Expecting '[' or ']', not %#U", r)
	}

	if r == ']' {
		// end scanner
		// this event has no attributes
		l.ignore()
		l.emit(itemEOF)
		return nil
	}

	if r == '[' {
		l.ignore()
	Loop:
		for {
			switch r := l.next(); {
			case isAlphaNumeric(r):
			// absorb.
			case r == '=':
				l.backup()
				l.emit(itemAttrKey)
				l.next()
				l.ignore()
				break Loop
			default:
				return l.errorf("Error on Attr Key! Expecting '=', not %#U", r)
			}
		}
	}
	return lexEventAttrValue
}

func lexEventAttrValue(l *lexer) stateFn {

Loop:
	for {
		r := l.next()

		if isAlphaNumeric(r) {
			// absorve letter
		} else if r == ']' {
			l.backup()
			l.emit(itemAttrValue)
			l.next()
			l.ignore()
			break Loop
		} else {
			return l.errorf("Error on Attr Value! Expecting ']', not %#U", r)
		}
	}
	return lexEventAttrKey
}

func (l *lexer) atEvtNameTerminator() bool {
	r := l.peek()
	switch r {
	case '=':
		return true
	}
	return false
}

// peek returns but does not consume the next rune in the input.
func (l *lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

func (l *lexer) emit(t itemType) {
	l.items <- item{t, l.start, l.input[l.start:l.pos]}
	l.start = l.pos
}

func (l *lexer) errorf(format string, args ...interface{}) stateFn {
	l.items <- item{itemError, l.start, fmt.Sprintf(format, args...)}
	return nil
}

// isAlphaNumeric reports whether r is an alphabetic, digit, or underscore.
func isAlphaNumeric(r rune) bool {
	return r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r)
}

// ignore skips over the pending input before this point.
func (l *lexer) ignore() {
	l.start = l.pos
}
