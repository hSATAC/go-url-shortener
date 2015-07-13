package shortener

import (
	"github.com/mattheath/base62"
)

// Shuffled from 0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz
const base62Chars = "cLAxbEUIfRPZWQMeY3tBNSorHhzG2Jv8isDdywnX5qFk49gaKuCjlVpT7106mO"

var shuffledBase62Encoder = base62.NewEncoding(base62Chars)

type Shortener struct {
	backend        Backend
	base62Encoding *base62.Encoding
}

func (s *Shortener) Shorten(url string) string {
	uniqID := s.getUniqID()
	return shuffledBase62Encoder.EncodeInt64(uniqID)
}

func (s *Shortener) getUniqID() int64 {
	return s.backend.GetUniqID()
}

func NewShortener(be Backend, b62 *base62.Encoding) *Shortener {
	return &Shortener{
		backend:        be,
		base62Encoding: b62,
	}
}
