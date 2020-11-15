package markdown

import (
	"bytes"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
)

type MarkdownContent []byte

func (c MarkdownContent) Convert() ([]byte, error) {
	var ren = html.WithUnsafe()
	var o = goldmark.WithRendererOptions(ren)
	var ext = goldmark.WithExtensions(extension.Table, extension.DefinitionList)
	var m = goldmark.New(ext, o)

	var err error
	var w = &bytes.Buffer{}
	err = m.Convert(c, w)
	if err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}