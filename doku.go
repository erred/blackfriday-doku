package doku

import (
	"bytes"
	"fmt"
	"io"

	"github.com/russross/blackfriday"
)

type Renderer struct {
	prefix   []byte
	ctrs     [][]byte
	tableSep []byte
}

func NewRenderer() *Renderer {
	return &Renderer{}
}

func (r *Renderer) RenderNode(w io.Writer, node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
	if node == nil {
		return blackfriday.Terminate
	}

	switch node.Type {

	case blackfriday.Heading:
		lvl := node.HeadingData.Level
		b := bytes.Repeat([]byte("="), 7-lvl)
		if 0 < lvl && lvl < 6 {
			if entering {
				fmt.Fprintf(w, "\n%s ", b)
				break
			}
			fmt.Fprintf(w, " %s\n", b)
		}

	case blackfriday.BlockQuote:
		if entering {
			r.prefix = append(r.prefix, []byte("> ")...)
			break
		}
		r.prefix = bytes.TrimSuffix(r.prefix, []byte("> "))
		if len(r.prefix) == 0 {
			w.Write([]byte("\n"))
		}

	case blackfriday.CodeBlock:
		fmt.Fprintf(w, "\n<code>\n%s</code>\n", node.Literal)

	case blackfriday.Code:
		fmt.Fprintf(w, "''%s''", node.Literal)
	case blackfriday.Emph:
		w.Write([]byte("//"))
	case blackfriday.Strong:
		w.Write([]byte("**"))
	case blackfriday.Del:
		if entering {
			w.Write([]byte("<del>"))
			break
		}
		w.Write([]byte("</del>"))

	case blackfriday.Paragraph:
		if entering {
			w.Write([]byte("\n"))
			break
		}
		if len(r.prefix) == 0 && len(r.ctrs) == 0 {
			w.Write([]byte("\n"))
		}
	case blackfriday.Text:
		if len(r.prefix) == 0 && len(r.ctrs) == 0 {
			w.Write(node.Literal)
			break
		}
		po := r.prefix
		if len(r.ctrs) > 0 {
			b := bytes.Repeat([]byte("  "), len(r.ctrs)-1)
			r.prefix = append(r.prefix, b...)
			r.prefix = append(r.prefix, r.ctrs[len(r.ctrs)-1]...)
		}
		w.Write(r.prefix)
		r.prefix = po
		w.Write(bytes.Join(bytes.Split(node.Literal, []byte("\n")), append([]byte("\n"), r.prefix...)))

	case blackfriday.HorizontalRule:
		w.Write([]byte("\n----\n"))

	case blackfriday.Link:
		if entering {
			fmt.Fprintf(w, "[[%s|", node.LinkData.Destination)
			break
		}
		w.Write([]byte("]]"))

	case blackfriday.Image:
		// unimplemented

	case blackfriday.List:
		if entering {
			if node.ListFlags&blackfriday.ListTypeOrdered > 0 {
				r.ctrs = append(r.ctrs, []byte("- "))
				break
			}
			r.ctrs = append(r.ctrs, []byte("* "))
			break
		}
		r.ctrs = r.ctrs[:len(r.ctrs)-1]
	case blackfriday.Item:
		// noop, see Text

	case blackfriday.Table:
		w.Write([]byte("\n"))
	case blackfriday.TableHead:
		if entering {
			r.tableSep = []byte("^")
			break
		}
		r.tableSep = []byte("|")
	case blackfriday.TableRow:
		if entering {
			w.Write([]byte("\n"))
			break
		}
		w.Write(r.tableSep)
	case blackfriday.TableCell:
		if entering {
			w.Write(r.tableSep)
			if node.TableCellData.Align != 0 {
				w.Write([]byte(" "))
			}
			break
		}
		if node.TableCellData.Align != blackfriday.TableAlignmentRight {
			w.Write([]byte(" "))
		}
	case blackfriday.TableBody:
		// noop

	case blackfriday.Document:
		// noop
	case blackfriday.HTMLBlock:
		// noop
	case blackfriday.Softbreak:
		// unimplemented
	case blackfriday.Hardbreak:
		// unimplemented
	case blackfriday.HTMLSpan:
		// noop

	default:
		// unimplemented
	}

	return blackfriday.GoToNext
}

// RenderHeader is a noop for satisfying blackfriday's Renderer
func (r *Renderer) RenderHeader(w io.Writer, ast *blackfriday.Node) {}

// RenderFooter is a noop for satisfying blackfriday's Renderer
func (r *Renderer) RenderFooter(w io.Writer, ast *blackfriday.Node) {}
