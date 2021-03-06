// Generate more fonts her
// http://patorjk.com/software/taag/#p=display&f=Small&t=.

package termcolors

import (
	"fmt"
	"strings"
)

var joker = `
 ___
/ _ \
\// /
  \/
  ()
`

var jokerInf letterInfos
var jokerLines []string

func init() {
	a, b := processOne(joker)
	jokerInf = b
	jokerLines = a
}

type letterInfos struct {
	lineNum  int
	maxWidth int
}

type Banner struct {
	font map[string][]string
	info map[string]letterInfos
}

func processOne(s string) ([]string, letterInfos) {
	lines := strings.Split(s, "\n")
	maxw := 0
	for i, line := range lines {
		if len(line) > maxw {
			maxw = len(line)
		}
		lines[i] = line
	}
	return lines, letterInfos{
		len(lines),
		maxw,
	}
}

func process(m map[string]string) (map[string][]string, map[string]letterInfos) {
	tr := map[string][]string{}
	inf := map[string]letterInfos{}
	for k, v := range m {
		a, b := processOne(v)
		tr[k] = a
		inf[k] = b
	}
	return tr, inf
}

func NewBanner(m map[string]string) Banner {
	trimmed, infos := process(m)
	return Banner{trimmed, infos}
}

func padRight(s string, width int) string {
	if len(s) < width {
		s += strings.Repeat(" ", width-len(s))
	}
	return s
}

func (b Banner) getOne(s string) ([]string, letterInfos) {
	linf, ok := b.info[s]
	if !ok {
		joker, okj := b.info["?"]
		if okj {
			return b.font["?"], joker
		}
		return jokerLines, jokerInf
	}
	return b.font[s], linf
}

func (b Banner) print(text string, printOut bool) string {
	bannerMaxHeight := 0
	// Calculating max height of banner
	for _, v := range text {
		_, linf := b.getOne(string(v))
		if linf.lineNum > bannerMaxHeight {
			bannerMaxHeight = linf.lineNum
		}
	}

	ret := ""
	// Render
	for i := 1; i < bannerMaxHeight-1; i++ {
		thisLin := ""
		for _, v := range text {
			lines, linf := b.getOne(string(v))
			if linf.lineNum <= i {
				thisLin += padRight("", linf.maxWidth)
			} else {
				thisLin += padRight(lines[i], linf.maxWidth)
			}
		}
		if printOut {
			fmt.Println(thisLin)
		} else {
			ret += thisLin + "\n"
		}
	}
	ret = strings.TrimRight(ret, "\n")
	return ret
}

func (b Banner) BannerPrintS(text string) string {
	return b.print(text, false)
}

func (b Banner) BannerPrint(text string) {
	b.print(text, true)
}

func BannerPrint(s string) {
	NewBanner(Ogre).BannerPrint(s)
}

func BannerPrintS(s string, font string) string {
	if font == "ogre" || font == "" {
		return NewBanner(Ogre).BannerPrintS(s)
	} else if font == "small" {
		return NewBanner(Small).BannerPrintS(s)
	}
	return ""
}
