//
// Blackfriday API Blueprint
// Available at http://github.com/Go-ApiBlueprint-Go/blackfriday
//

//
//
// API Blueprint rendering backend
//
//

package blackfriday

import (
	"bytes"
)

// ApiBlueprint is a type that implements the Renderer interface for ApiBlueprint output.
//
// Do not create this directly, instead use the ApiBlueprintRenderer function.
type ApiBlueprint struct {
}

// ApiBlueprintRenderer creates and configures a ApiBlueprint object, which
// satisfies the Renderer interface.
//
// flags is a set of APIBLUEPRINT_* options ORed together (currently no such options
// are defined).
func ApiBlueprintRenderer(flags int) Renderer {
	return &ApiBlueprint{}
}

func (options *ApiBlueprint) GetFlags() int {
	return 0
}

// render code chunks using verbatim, or listings if we have a language
func (options *ApiBlueprint) BlockCode(out *bytes.Buffer, text []byte, lang string) {
	out.Write(text)
}

func (options *ApiBlueprint) TitleBlock(out *bytes.Buffer, text []byte) {

}

func (options *ApiBlueprint) BlockQuote(out *bytes.Buffer, text []byte) {
	out.Write(text)
}

func (options *ApiBlueprint) BlockHtml(out *bytes.Buffer, text []byte) {
	// a pretty lame thing to do...
	out.Write(text)
}

func (options *ApiBlueprint) Header(out *bytes.Buffer, text func() bool, level int, id string) {
	marker := out.Len()

	switch level {
	case 1:
		out.WriteString("\n#")
	case 2:
		out.WriteString("\n##")
	case 3:
		out.WriteString("\n###")
	case 4:
		out.WriteString("\n")
	case 5:
		out.WriteString("\n")
	case 6:
		out.WriteString("\n")
	}
	if !text() {
		out.Truncate(marker)
		return
	}
	out.WriteString("\n")
}

func (options *ApiBlueprint) HRule(out *bytes.Buffer) {
	out.WriteString("\n\n")
}

func (options *ApiBlueprint) List(out *bytes.Buffer, text func() bool, flags int) {
	marker := out.Len()
	if !text() {
		out.Truncate(marker)
		return
	}
}

func (options *ApiBlueprint) ListItem(out *bytes.Buffer, text []byte, flags int) {
	out.WriteString("\n ")
	out.Write(text)
}

func (options *ApiBlueprint) Paragraph(out *bytes.Buffer, text func() bool) {
	marker := out.Len()
	out.WriteString("\n")
	if !text() {
		out.Truncate(marker)
		return
	}
	out.WriteString("\n")
}

func (options *ApiBlueprint) Table(out *bytes.Buffer, header []byte, body []byte, columnData []int) {
	out.WriteString("\n")
	out.WriteString("\n")
	out.Write(header)
	out.WriteString("\n\n")
	out.Write(body)
	out.WriteString("\n\n")
}

func (options *ApiBlueprint) TableRow(out *bytes.Buffer, text []byte) {
	if out.Len() > 0 {
		out.WriteString(" \n")
	}
	out.Write(text)
}

func (options *ApiBlueprint) TableHeaderCell(out *bytes.Buffer, text []byte, align int) {
	if out.Len() > 0 {
		out.WriteString(" & ")
	}
	out.Write(text)
}

func (options *ApiBlueprint) TableCell(out *bytes.Buffer, text []byte, align int) {
	if out.Len() > 0 {
		out.WriteString(" & ")
	}
	out.Write(text)
}

// TODO: this
func (options *ApiBlueprint) Footnotes(out *bytes.Buffer, text func() bool) {

}

func (options *ApiBlueprint) FootnoteItem(out *bytes.Buffer, name, text []byte, flags int) {

}

func (options *ApiBlueprint) AutoLink(out *bytes.Buffer, link []byte, kind int) {
	out.Write(link)
}

func (options *ApiBlueprint) CodeSpan(out *bytes.Buffer, text []byte) {
	escapeSpecialChars(out, text)
}

func (options *ApiBlueprint) DoubleEmphasis(out *bytes.Buffer, text []byte) {
	out.Write(text)
}

func (options *ApiBlueprint) Emphasis(out *bytes.Buffer, text []byte) {
	out.WriteString("*")
	out.Write(text)
	out.WriteString("*")
}

func (options *ApiBlueprint) Image(out *bytes.Buffer, link []byte, title []byte, alt []byte) {
	out.Write(link)
}

func (options *ApiBlueprint) LineBreak(out *bytes.Buffer) {
	out.WriteString("\n")
}

func (options *ApiBlueprint) Link(out *bytes.Buffer, link []byte, title []byte, content []byte) {
	out.Write(link)
	out.Write(content)
}

func (options *ApiBlueprint) RawHtmlTag(out *bytes.Buffer, tag []byte) {
}

func (options *ApiBlueprint) TripleEmphasis(out *bytes.Buffer, text []byte) {
	out.Write(text)
}

func (options *ApiBlueprint) StrikeThrough(out *bytes.Buffer, text []byte) {
	out.Write(text)
}

// TODO: this
func (options *ApiBlueprint) FootnoteRef(out *bytes.Buffer, ref []byte, id int) {

}

// func needsBackslash(c byte) bool {
// 	for _, r := range []byte("_{}%$&\\~#") {
// 		if c == r {
// 			return true
// 		}
// 	}
// 	return false
// }

// func escapeSpecialChars(out *bytes.Buffer, text []byte) {
// 	for i := 0; i < len(text); i++ {
// 		// directly copy normal characters
// 		org := i

// 		for i < len(text) && !needsBackslash(text[i]) {
// 			i++
// 		}
// 		if i > org {
// 			out.Write(text[org:i])
// 		}

// 		// escape a character
// 		if i >= len(text) {
// 			break
// 		}
// 		out.WriteByte('\\')
// 		out.WriteByte(text[i])
// 	}
// }

func (options *ApiBlueprint) Entity(out *bytes.Buffer, entity []byte) {
	// TODO: convert this into a unicode character or something
	out.Write(entity)
}

func (options *ApiBlueprint) NormalText(out *bytes.Buffer, text []byte) {
	escapeSpecialChars(out, text)
}

// header and footer
func (options *ApiBlueprint) DocumentHeader(out *bytes.Buffer) {
	out.WriteString("\nSTART\n")
	out.WriteString(VERSION)
	out.WriteString("\n\n")
}

func (options *ApiBlueprint) DocumentFooter(out *bytes.Buffer) {
	out.WriteString("\nEND\n")
}

// ApiBlueprintBasic is a convenience function for simple rendering.
func ApiBlueprintBasic(input []byte) []byte {
	// set up the HTML renderer
	htmlFlags := HTML_USE_XHTML
	renderer := ApiBlueprintRenderer(htmlFlags)

	// set up the parser
	return MarkdownOptions(input, renderer, Options{Extensions: 0})
}
