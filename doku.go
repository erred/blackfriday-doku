package doku

import (
	"io"

	"github.com/russross/blackfriday"
)

type Renderer struct {
}

func NewRenderer() *Renderer {
	return &Renderer{}
}

func (r *Renderer) RenderNode(w io.Writer, node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
	if node == nil {
		// handle error
		return blackfriday.Terminate
	}
	switch node.Type {
	case blackfriday.Document:
	case blackfriday.BlockQuote:
	case blackfriday.List:
	case blackfriday.Item:
	case blackfriday.Paragraph:
	case blackfriday.Heading:
	case blackfriday.HorizontalRule:
	case blackfriday.Emph:
	case blackfriday.Strong:
	case blackfriday.Del:
	case blackfriday.Link:
	case blackfriday.Image:
	case blackfriday.Text:
	case blackfriday.HTMLBlock:
	case blackfriday.CodeBlock:
	case blackfriday.Softbreak:
	case blackfriday.Hardbreak:
	case blackfriday.Code:
	case blackfriday.HTMLSpan:
	case blackfriday.Table:
	case blackfriday.TableCell:
	case blackfriday.TableHead:
	case blackfriday.TableBody:
	case blackfriday.TableRow:
	}

	return blackfriday.GoToNext
}

// RenderHeader is a noop for satisfying blackfriday's Renderer
func (r *Renderer) RenderHeader(w io.Writer, ast *blackfriday.Node) {}

// RenderFooter is a noop for satisfying blackfriday's Renderer
func (r *Renderer) RenderFooter(w io.Writer, ast *blackfriday.Node) {}
