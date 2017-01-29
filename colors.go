package colors

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	BOLD          = "\033[1m"
	ITALIC        = "\033[3m"
	UNDERLINE     = "\033[4m"
	STRIKETHROUGH = "\033[9m"
	REVERSED      = "\u001b[7m"
	RESET         = "\033[0m"
	BLINK         = "\x1b[5m"
)

const (
	BLACK         = "0"
	BRIGHTBLACK   = "0;1"
	RED           = "1"
	BRIGHTRED     = "1;1"
	GREEN         = "2"
	BRIGHTGREEN   = "2;1"
	YELLOW        = "3"
	BRIGHTYELLOW  = "3;1"
	BLUE          = "4"
	BRIGHTBLUE    = "4;1"
	MAGENTA       = "5"
	BRIGHTMAGENTA = "5;1"
	CYAN          = "6"
	BRIGHTCYAN    = "6;1"
	WHITE         = "7"
	BRIGHTWHITE   = "7;1"
)

var (
	Output *bufio.Writer = bufio.NewWriter(os.Stdout)
)

func getColor(code string) string {
	//return fmt.Sprintf("\033[3%sm", code)
	return fmt.Sprintf("\u001b[3%sm", code)
}

func getBgColor(code string) string {
	//return fmt.Sprintf("\033[4%sm", code)
	return fmt.Sprintf("\u001b[4%sm", code)
}

func Bold(str string) string {
	return fmt.Sprintf("%s%s%s", BOLD, str, RESET)
}

func Underline(str string) string {
	return fmt.Sprintf("%s%s%s", UNDERLINE, str, RESET)
}

func Italic(str string) string {
	return fmt.Sprintf("%s%s%s", ITALIC, str, RESET)
}

func Background(str string, color string) string {
	return fmt.Sprintf("%s%s%s", getBgColor(color), str, RESET)
}

func Color(str string, color string) string {
	return fmt.Sprintf("%s%s%s", getColor(color), str, RESET)
}

func Highlight(str, substr string, color string) string {
	hiSubstr := Color(substr, color)
	return strings.Replace(str, substr, hiSubstr, -1)
}

func MoveTo(str string, x int, y int) (out string) {
	//x, y = GetXY(x, y)

	return fmt.Sprintf("\033[%d;%dH%s", y, x, str)
}

func Reversed(str string) string {
	return fmt.Sprintf("%s%s%s", REVERSED, str, RESET)
}

// TO LOOK AT

func Blink(str string) string {
	return fmt.Sprintf("%s%s%s", BLINK, str, RESET)
}

func StrikeThrough(str string) string {
	return fmt.Sprintf("%s%s%s", STRIKETHROUGH, str, RESET)
}

func test() {
	Clear()
	fmt.Println(Bold("BOLD"))
	fmt.Println(Italic("ITALIC"))
	fmt.Println(Underline("UNDERLINE"))
	fmt.Println(StrikeThrough("STRIKETHROUGH"))
	fmt.Println(Blink("BLINK"))
	fmt.Println(Reversed("REVERSED"))
	fmt.Println(Background("BACKGROUND", RED))
	fmt.Println(Color("COLOR", YELLOW))
	fmt.Println(Highlight("HIGHLIGHT", "GHLI", BRIGHTYELLOW))
	//fmt.Println(MoveTo("TESTING", 40, 0))
	fmt.Println(Color(Bold("BOLD YELLOW"), YELLOW))
	fmt.Println(Background(Color(Bold("BOLD YELLOW WITH BACKGROUND"), YELLOW), RED))
	fmt.Println(Color("BLACK", BLACK))
	fmt.Println(Color("BRIGHTBLACK", BRIGHTBLACK))
	fmt.Println(Color("RED", RED))
	fmt.Println(Color("BRIGHTRED", BRIGHTRED))
	fmt.Println(Color("GREEN", GREEN))
	fmt.Println(Color("BRIGHTGREEN", BRIGHTGREEN))
	fmt.Println(Color("YELLOW", YELLOW))
	fmt.Println(Color("BRIGHTYELLOW", BRIGHTYELLOW))
	fmt.Println(Color("BLUE", BLUE))
	fmt.Println(Color("BRIGHTBLUE", BRIGHTBLUE))
	fmt.Println(Color("MAGENTA", MAGENTA))
	fmt.Println(Color("BRIGHTMAGENTA", BRIGHTMAGENTA))
	fmt.Println(Color("WHITE", WHITE))
	fmt.Println(Color("BRIGHTWHITE", BRIGHTWHITE))
	//fmt.Println(Color("GRAY", GRAY))
}