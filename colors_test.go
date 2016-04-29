package colors

import (
  "fmt"
  "testing"
)

func TestBold(t *testing.T) {
  bold := Bold("bold")
  fmt.Printf("Testing Bold: %s\n", bold)
  if bold != "\x1b[1m" + "bold" + "\x1b[22m" {
    t.Fail()
  }
}

func TestItalic(t *testing.T) {
  italic := Italic("italic")
  fmt.Printf("Testing Italic: %s\n", italic)
  if italic != "\x1b[3m" + "italic" + "\x1b[23m" {
    t.Fail()
  }
}

func TestUnderline(t *testing.T) {
  underline := Underline("underline")
  fmt.Printf("Testing Underline: %s\n", underline)
  if underline != "\x1b[4m" + "underline" + "\x1b[24m" {
    t.Fail()
  }
}

func TestInverse(t *testing.T) {
  inverse := Inverse("inverse")
  fmt.Printf("Testing Inverse: %s\n", inverse)
  if inverse != "\x1b[7m" + "inverse" + "\x1b[27m" {
    t.Fail()
  }
}

func TestStrikethrough(t *testing.T) {
  strikethrough := Strikethrough("strikethrough")
  fmt.Printf("Testing Strikethrough: %s\n", strikethrough)
  if strikethrough != "\x1b[9m" + "strikethrough" + "\x1b[29m" {
    t.Fail()
  }
}

func TestBlack(t *testing.T) {
  black := Black("black")
  fmt.Printf("Testing Black: %s\n", black)
  if black != "\x1b[30m" + "black" + "\x1b[39m" {
    t.Fail()
  }
}

func TestRed(t *testing.T) {
  red := Red("red")
  fmt.Printf("Testing Red: %s\n", red)
  if red != "\x1b[31m" + "red" + "\x1b[39m" {
    t.Fail()
  }
}

func TestGreen(t *testing.T) {
  green := Green("green")
  fmt.Printf("Testing Green: %s\n", green)
  if green != "\x1b[32m" + "green" + "\x1b[39m" {
    t.Fail()
  }
}

func TestYellow(t *testing.T) {
  yellow := Yellow("yellow")
  fmt.Printf("Testing Yellow: %s\n", yellow)
  if yellow != "\x1b[33m" + "yellow" + "\x1b[39m" {
    t.Fail()
  }
}

func TestBlue(t *testing.T) {
  blue := Blue("blue")
  fmt.Printf("Testing Blue: %s\n", blue)
  if blue != "\x1b[34m" + "blue" + "\x1b[39m" {
    t.Fail()
  }
}

func TestMagenta(t *testing.T) {
  magenta := Magenta("magenta")
  fmt.Printf("Testing Magenta: %s\n", magenta)
  if magenta != "\x1b[35m" + "magenta" + "\x1b[39m" {
    t.Fail()
  }
}

func TestCyan(t *testing.T) {
  cyan := Cyan("cyan")
  fmt.Printf("Testing Cyan: %s\n", cyan)
  if cyan != "\x1b[36m" + "cyan" + "\x1b[39m" {
    t.Fail()
  }
}

func TestWhite(t *testing.T) {
  white := White("white")
  fmt.Printf("Testing White: %s\n", white)
  if white != "\x1b[37m" + "white" + "\x1b[39m" {
    t.Fail()
  }
}

func TestBgBlack(t *testing.T) {
  black := BgBlack("bg black")
  fmt.Printf("Testing BgBlack: %s\n", black)
  if black != "\x1b[40m" + "bg black" + "\x1b[49m" {
    t.Fail()
  }
}

func TestBgRed(t *testing.T) {
  red := BgRed("bg red")
  fmt.Printf("Testing BgRed: %s\n", red)
  if red != "\x1b[41m" + "bg red" + "\x1b[49m" {
    t.Fail()
  }
}

func TestBgGreen(t *testing.T) {
  green := BgGreen("bg green")
  fmt.Printf("Testing BgGreen: %s\n", green)
  if green != "\x1b[42m" + "bg green" + "\x1b[49m" {
    t.Fail()
  }
}

func TestBgYellow(t *testing.T) {
  yellow := BgYellow("bg yellow")
  fmt.Printf("Testing BgYellow: %s\n", yellow)
  if yellow != "\x1b[43m" + "bg yellow" + "\x1b[49m" {
    t.Fail()
  }
}

func TestBgBlue(t *testing.T) {
  blue := BgBlue("bg blue")
  fmt.Printf("Testing BgBlue: %s\n", blue)
  if blue != "\x1b[44m" + "bg blue" + "\x1b[49m" {
    t.Fail()
  }
}

func TestBgMagenta(t *testing.T) {
  magenta := BgMagenta("bg magenta")
  fmt.Printf("Testing BgMagenta: %s\n", magenta)
  if magenta != "\x1b[45m" + "bg magenta" + "\x1b[49m" {
    t.Fail()
  }
}

func TestBgCyan(t *testing.T) {
  cyan := BgCyan("bg cyan")
  fmt.Printf("Testing BgCyan: %s\n", cyan)
  if cyan != "\x1b[46m" + "bg cyan" + "\x1b[49m" {
    t.Fail()
  }
}

func TestBgWhite(t *testing.T) {
  white := BgWhite("bg white")
  fmt.Printf("Testing BgWhite: %s\n", white)
  if white != "\x1b[47m" + "bg white" + "\x1b[49m" {
    t.Fail()
  }
}
