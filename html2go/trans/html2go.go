package trans

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

var elemNameMap = map[string]string{
	"a":          "elem.Anchor(",
	"abbr":       "elem.Abbreviation(",
	"address":    "elem.Address(",
	"area":       "elem.Area(",
	"article":    "elem.Article(",
	"aside":      "elem.ASide(",
	"audio":      "elem.Audio(",
	"b":          "elem.Bold(",
	"base":       "elem.Base(",
	"bdi":        "elem.BidirectionalIsolation(",
	"bdo":        "elem.BidirectionalOverride(",
	"blockquote": "elem.BlockQuote(",
	"body":       "elem.Body(",
	"br":         "elem.Break(",
	"button":     "elem.Button(",
	"canvas":     "elem.Canvas(",
	"caption":    "elem.Caption(",
	"cite":       "elem.Citation(",
	"code":       "elem.Code(",
	"col":        "elem.Column(",
	"colgroup":   "elem.ColumnGroup(",
	"data":       "elem.Data(",
	"datalist":   "elem.DataList(",
	"dd":         "elem.Description(",
	"del":        "elem.DeletedText(",
	"details":    "elem.Details(",
	"dfn":        "elem.Definition(",
	"dialog":     "elem.Dialog(",
	"div":        "elem.Div(",
	"dl":         "elem.DescriptionList(",
	"dt":         "elem.DefinitionTerm(",
	"em":         "elem.Emphasis(",
	"embed":      "elem.Embed(",
	"fieldset":   "elem.FieldSet(",
	"figcaption": "elem.FigureCaption(",
	"figure":     "elem.Figure(",
	"footer":     "elem.Footer(",
	"form":       "elem.Form(",
	"h1":         "elem.Heading1(",
	"h2":         "elem.Heading2(",
	"h3":         "elem.Heading3(",
	"h4":         "elem.Heading4(",
	"h5":         "elem.Heading5(",
	"h6":         "elem.Heading6(",
	"header":     "elem.Header(",
	"hgroup":     "elem.HeadingsGroup(",
	"hr":         "elem.HorizontalRule(",
	"i":          "elem.Italic(",
	"iframe":     "elem.InlineFrame(",
	"img":        "elem.Image(",
	"input":      "elem.Input(",
	"ins":        "elem.InsertedText(",
	"kbd":        "elem.KeyboardInput(",
	"label":      "elem.Label(",
	"legend":     "elem.Legend(",
	"li":         "elem.ListItem(",
	"link":       "elem.Link(",
	"main":       "elem.Main(",
	"map":        "elem.Map(",
	"mark":       "elem.Mark(",
	"menu":       "elem.Menu(",
	"menuitem":   "elem.MenuItem(",
	"meta":       "elem.Meta(",
	"meter":      "elem.Meter(",
	"nav":        "elem.Navigation(",
	"noframes":   "elem.NoFrames(",
	"noscript":   "elem.NoScript(",
	"object":     "elem.Object(",
	"ol":         "elem.OrderedList(",
	"optgroup":   "elem.OptionsGroup(",
	"option":     "elem.Option(",
	"output":     "elem.Output(",
	"p":          "elem.Paragraph(",
	"param":      "elem.Parameter(",
	"picture":    "elem.Picture(",
	"pre":        "elem.Preformatted(",
	"progress":   "elem.Progress(",
	"q":          "elem.Quote(",
	"rp":         "elem.RubyParenthesis(",
	"rt":         "elem.RubyText(",
	"rtc":        "elem.RubyTextContainer(",
	"ruby":       "elem.Ruby(",
	"s":          "elem.Strikethrough(",
	"samp":       "elem.Sample(",
	"script":     "elem.Script(",
	"section":    "elem.Section(",
	"select":     "elem.Select(",
	"slot":       "elem.Slot(",
	"small":      "elem.Small(",
	"source":     "elem.Source(",
	"span":       "elem.Span(",
	"strong":     "elem.Strong(",
	"style":      "elem.Style(",
	"sub":        "elem.Subscript(",
	"summary":    "elem.Summary(",
	"sup":        "elem.Superscript(",
	"table":      "elem.Table(",
	"tbody":      "elem.TableBody(",
	"td":         "elem.TableData(",
	"template":   "elem.Template(",
	"textarea":   "elem.TextArea(",
	"tfoot":      "elem.TableFoot(",
	"th":         "elem.TableHeader(",
	"thead":      "elem.TableHead(",
	"time":       "elem.Time(",
	"tr":         "elem.TableRow(",
	"track":      "elem.Track(",
	"u":          "elem.Underline(",
	"ul":         "elem.UnorderedList(",
	"var":        "elem.Variable(",
	"video":      "elem.Video(",
	"wbr":        "elem.WordBreakOpportunity(",
}

func tab(level int) string {
	return strings.Repeat("\t", level)
}

func markup(w io.Writer, attrs []html.Attribute, level int) {
	fmt.Fprintf(w, "%svecty.Markup(\n", tab(level))
	for _, a := range attrs {
		switch strings.ToLower(a.Key) {
		default:
			fmt.Fprintf(w, "%svecty.Attribute(%q, %q),\n", tab(level+1), a.Key, a.Val)
		case "class":
			classes := strings.Fields(a.Val)
			if len(classes) == 1 {
				fmt.Fprintf(w, "%svecty.Class(%q),\n", tab(level+1), classes[0])
			} else {
				fmt.Fprintf(w, "%svecty.ClassMap{\n", tab(level+1))
				for _, c := range classes {
					fmt.Fprintf(w, "%s%q: true,\n", tab(level+2), c)
				}
				fmt.Fprintf(w, "%s},\n", tab(level+1))
			}
		case "id":
			fmt.Fprintf(w, "%sprop.ID(%q),\n", tab(level+1), a.Val)
		case "href":
			fmt.Fprintf(w, "%sprop.Href(%q),\n", tab(level+1), a.Val)
		case "placeholder":
			fmt.Fprintf(w, "%sprop.Placeholder(%q),\n", tab(level+1), a.Val)
		case "src":
			fmt.Fprintf(w, "%sprop.Src(%q),\n", tab(level+1), a.Val)
		case "type":
			fmt.Fprintf(w, "%sprop.Type(%q),\n", tab(level+1), strings.ToLower(a.Val))
		case "for":
			fmt.Fprintf(w, "%sprop.For(%q),\n", tab(level+1), a.Val)
		case "value":
			fmt.Fprintf(w, "%sprop.Value(%q),\n", tab(level+1), a.Val)
		case "autofocus":
			if a.Val == "" {
				a.Val = "true"
			}
			fmt.Fprintf(w, "%sprop.Autofocus(%s),\n", tab(level+1), a.Val)
		case "checked":
			if a.Val == "" {
				a.Val = "true"
			}
			fmt.Fprintf(w, "%sprop.Checked(%s),\n", tab(level+1), a.Val)
		}
	}
	fmt.Fprintf(w, "%s),\n", tab(level))
}

func walk(w io.Writer, n *html.Node, level int,
	withBody, found bool) error {
	switch n.Type {
	default:
		return fmt.Errorf("unknown html type: %d", n.Type)
	case html.ErrorNode:
		return fmt.Errorf("ErrorNode: %s", n.Data)
	case html.DocumentNode:
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if err := walk(w, c, level, withBody, false); err != nil {
				return err
			}
		}
	case html.DoctypeNode:
		if !found {
			return nil
		}
		fmt.Fprintf(w, "%s%s %s\n", tab(level),
			"DoctypeNode:", n.Data)
	case html.CommentNode:
		if !found {
			return nil
		}
		data := strings.TrimSpace(n.Data)
		fmt.Fprintf(w, "%s// %s\n", tab(level), data)
	case html.TextNode:
		if !found {
			return nil
		}
		data := strings.TrimSpace(n.Data)
		if data != "" {
			fmt.Fprintf(w, "%svecty.Text(%q),\n", tab(level), data)
		}
	case html.ElementNode:
		if n.DataAtom == atom.Body {
			found = true
			if !withBody {
				for c := n.FirstChild; c != nil; c = c.NextSibling {
					walk(w, c, level, withBody, found)
				}
				return nil
			}
		}
		if !found {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				walk(w, c, level, withBody, found)
			}
			return nil
		}
		elem := elemNameMap[n.Data]
		if elem == "" {
			elem = fmt.Sprintf("vecty.Tag(%q,", n.Data)
		}
		if len(n.Attr) > 0 {
			fmt.Fprintf(w, "%s%s\n", tab(level), elem)
			markup(w, n.Attr, level+1)
		} else {
			fmt.Fprintf(w, "%s%s\n", tab(level), elem)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(w, c, level+1, withBody, found)
		}
		if level == 0 {
			fmt.Fprintf(w, "%s)\n", tab(level))
		} else {
			fmt.Fprintf(w, "%s),\n", tab(level))
		}
	}
	return nil
}

// Trans ...
func Trans(w io.Writer, r io.Reader, withBody bool) error {
	doc, err := html.Parse(r)
	if err != nil {
		return err
	}
	return walk(w, doc, 0, withBody, false)
}
