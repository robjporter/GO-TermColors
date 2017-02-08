package main

import (
	"../"
	"fmt"
)

func main() {
	termcolors.ClearScreen()
	fmt.Println(termcolors.BannerPrintS("aci", "small"))
	fmt.Println(termcolors.BannerPrintLineS("=", 40))
	fmt.Println(termcolors.BannerPrintLineCommentS("=", "TESTING", 40))
	fmt.Println(termcolors.Bold("BOLD"))
	fmt.Println(termcolors.Italic("ITALIC"))
	fmt.Println(termcolors.Underline("UNDERLINE"))
	fmt.Println(termcolors.StrikeThrough("STRIKETHROUGH"))
	fmt.Println(termcolors.Blink("BLINK"))
	fmt.Println(termcolors.Reversed("REVERSED"))
	fmt.Println(termcolors.Background("BACKGROUND", termcolors.RED))
	fmt.Println(termcolors.Color("COLOR", termcolors.YELLOW))
	fmt.Println(termcolors.Highlight("HIGHLIGHT", "GHLI", termcolors.BRIGHTYELLOW))
	//fmt.Println(MoveTo("TESTING", 40, 0))
	fmt.Println(termcolors.Color(termcolors.Bold("BOLD YELLOW"), termcolors.YELLOW))
	fmt.Println(termcolors.Background(termcolors.Color(termcolors.Bold("BOLD YELLOW WITH BACKGROUND"), termcolors.YELLOW), termcolors.RED))
	fmt.Println(termcolors.Color("BLACK", termcolors.BLACK))
	fmt.Println(termcolors.Color("BRIGHTBLACK", termcolors.BRIGHTBLACK))
	fmt.Println(termcolors.Color("RED", termcolors.RED))
	fmt.Println(termcolors.Color("BRIGHTRED", termcolors.BRIGHTRED))
	fmt.Println(termcolors.Color("GREEN", termcolors.GREEN))
	fmt.Println(termcolors.Color("BRIGHTGREEN", termcolors.BRIGHTGREEN))
	fmt.Println(termcolors.Color("YELLOW", termcolors.YELLOW))
	fmt.Println(termcolors.Color("BRIGHTYELLOW", termcolors.BRIGHTYELLOW))
	fmt.Println(termcolors.Color("BLUE", termcolors.BLUE))
	fmt.Println(termcolors.Color("BRIGHTBLUE", termcolors.BRIGHTBLUE))
	fmt.Println(termcolors.Color("MAGENTA", termcolors.MAGENTA))
	fmt.Println(termcolors.Color("BRIGHTMAGENTA", termcolors.BRIGHTMAGENTA))
	fmt.Println(termcolors.Color("WHITE", termcolors.WHITE))
	fmt.Println(termcolors.Color("BRIGHTWHITE", termcolors.BRIGHTWHITE))
}
