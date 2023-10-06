package sourcesans3

import (
	"sync"

	"gio.tools/fonts/sourcesans3/sourcesans3black"
	"gio.tools/fonts/sourcesans3/sourcesans3blackitalic"
	"gio.tools/fonts/sourcesans3/sourcesans3bold"
	"gio.tools/fonts/sourcesans3/sourcesans3bolditalic"
	"gio.tools/fonts/sourcesans3/sourcesans3extrabold"
	"gio.tools/fonts/sourcesans3/sourcesans3extrabolditalic"
	"gio.tools/fonts/sourcesans3/sourcesans3extralight"
	"gio.tools/fonts/sourcesans3/sourcesans3extralightitalic"
	"gio.tools/fonts/sourcesans3/sourcesans3italic"
	"gio.tools/fonts/sourcesans3/sourcesans3light"
	"gio.tools/fonts/sourcesans3/sourcesans3lightitalic"
	"gio.tools/fonts/sourcesans3/sourcesans3medium"
	"gio.tools/fonts/sourcesans3/sourcesans3mediumitalic"
	"gio.tools/fonts/sourcesans3/sourcesans3regular"
	"gio.tools/fonts/sourcesans3/sourcesans3semibold"
	"gio.tools/fonts/sourcesans3/sourcesans3semibolditalic"

	"gioui.org/font"
	"gioui.org/font/opentype"
)

var (
	once       sync.Once
	collection []font.FontFace
)

func Collection() []font.FontFace {
	once.Do(func() {
		register(sourcesans3black.TTF)
		register(sourcesans3blackitalic.TTF)
		register(sourcesans3bold.TTF)
		register(sourcesans3bolditalic.TTF)
		register(sourcesans3extrabold.TTF)
		register(sourcesans3extrabolditalic.TTF)
		register(sourcesans3extralight.TTF)
		register(sourcesans3extralightitalic.TTF)
		register(sourcesans3italic.TTF)
		register(sourcesans3light.TTF)
		register(sourcesans3lightitalic.TTF)
		register(sourcesans3medium.TTF)
		register(sourcesans3mediumitalic.TTF)
		register(sourcesans3regular.TTF)
		register(sourcesans3semibold.TTF)
		register(sourcesans3semibolditalic.TTF)
		// Ensure that any outside appends will not reuse the backing store.
		n := len(collection)
		collection = collection[:n:n]
	})
	return collection
}

func register(src []byte) {
	faces, err := opentype.ParseCollection(src)
	if err != nil {
		panic("failed to parse font: " + err.Error())
	}
	collection = append(collection, faces[0])
}
