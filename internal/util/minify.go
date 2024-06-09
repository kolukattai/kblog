package util

import (
	"github.com/charmbracelet/log"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
)

type mf struct{}

func Minify() *mf {
	return &mf{}
}

func (s *mf) Css(fileData []byte) []byte {
	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	minifiedCSS, err := m.Bytes("text/css", fileData)
	if err != nil {
		log.Fatalf("Failed to minify CSS: %v\n", err)
	}

	return minifiedCSS
}

func (s *mf) JS(fileData []byte) []byte {
	m := minify.New()
	m.AddFunc("text/javascript", css.Minify)
	minifiedCSS, err := m.Bytes("text/javascript", fileData)
	if err != nil {
		log.Fatalf("Failed to minify Javascript: %v\n", err)
	}

	return minifiedCSS
}

func (s *mf) HTML(fileData []byte) []byte {
	m := minify.New()
	m.AddFunc("text/html", css.Minify)
	minifiedCSS, err := m.Bytes("text/html", fileData)
	if err != nil {
		log.Fatalf("Failed to minify Javascript: %v\n", err)
	}

	return minifiedCSS
}
