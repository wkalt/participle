package lexer_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	. "github.com/alecthomas/participle/v2/lexer"
)

type staticLexer struct {
	tokens []Token
}

func (s *staticLexer) Next() (Token, error) {
	if len(s.tokens) == 0 {
		return EOFToken(Position{}), nil
	}
	t := s.tokens[0]
	s.tokens = s.tokens[1:]
	return t, nil
}

func TestUpgrade(t *testing.T) {
	t0 := Token{Type: 1, Value: "moo"}
	ts := Token{Type: 3, Value: " "}
	t1 := Token{Type: 2, Value: "blah"}
	tokens := []Token{t0, ts, t1}
	l, err := Upgrade(&staticLexer{tokens: tokens}, 3)
	require.NoError(t, err)
	require.Equal(t, t0, mustPeek(t, l, 0))
	require.Equal(t, t0, mustPeek(t, l, 0))
	require.Equal(t, t1, mustPeek(t, l, 1))
	require.Equal(t, t1, mustPeek(t, l, 1))
	require.True(t, mustPeek(t, l, 2).EOF())
	require.True(t, mustPeek(t, l, 3).EOF())
	require.Equal(t, tokens, l.Range(0, 3))
}

func mustPeek(t *testing.T, lexer *PeekingLexer, n int) Token {
	t.Helper()
	token, err := lexer.Peek(n)
	require.NoError(t, err)
	return token
}

func mustNext(t *testing.T, lexer Lexer) Token {
	t.Helper()
	token, err := lexer.Next()
	require.NoError(t, err)
	return token
}
