package shortener

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestShorten(t *testing.T) {
	Convey("Should return a string", t, func() {
		var backend = NewDummyBackend()
		var shortener = NewShortener(backend, shuffledBase62Encoder)
		shortened := shortener.Shorten("http://blog.hsatac.net")
		So(shortened, ShouldEqual, "LAi")
	})
}
