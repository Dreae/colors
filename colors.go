package colors

import "bytes"

func Bold(in string) string {
  return stylize("\x1b[1m", in, "\x1b[22m")
}

func Italic(in string) string {
  return stylize("\x1b[3m", in, "\x1b[23m")
}

func Underline(in string) string {
  return stylize("\x1b[4m", in, "\x1b[24m")
}

func Inverse(in string) string {
  return stylize("\x1b[7m", in, "\x1b[27m")
}

func Strikethrough(in string) string {
  return stylize("\x1b[9m", in, "\x1b[29m")
}

func Black(in string) string {
  return stylize("\x1b[30m", in, "\x1b[39m")
}

func Red(in string) string {
  return stylize("\x1b[31m", in, "\x1b[39m")
}

func Green(in string) string {
  return stylize("\x1b[32m", in, "\x1b[39m")
}

func Yellow(in string) string {
  return stylize("\x1b[33m", in, "\x1b[39m")
}

func Blue(in string) string {
  return stylize("\x1b[34m", in, "\x1b[39m")
}

func Magenta(in string) string {
  return stylize("\x1b[35m", in, "\x1b[39m")
}

func Cyan(in string) string {
  return stylize("\x1b[36m", in, "\x1b[39m")
}

func White(in string) string {
  return stylize("\x1b[37m", in, "\x1b[39m")
}

func BgBlack(in string) string {
  return stylize("\x1b[40m", in, "\x1b[49m")
}

func BgRed(in string) string {
  return stylize("\x1b[41m", in, "\x1b[49m")
}

func BgGreen(in string) string {
  return stylize("\x1b[42m", in, "\x1b[49m")
}

func BgYellow(in string) string {
  return stylize("\x1b[43m", in, "\x1b[49m")
}

func BgBlue(in string) string {
  return stylize("\x1b[44m", in, "\x1b[49m")
}

func BgMagenta(in string) string {
  return stylize("\x1b[45m", in, "\x1b[49m")
}

func BgCyan(in string) string {
  return stylize("\x1b[46m", in, "\x1b[49m")
}

func BgWhite(in string) string {
  return stylize("\x1b[47m", in, "\x1b[49m")
}

func stylize(style string, text string, closing string) string {
  var buffer bytes.Buffer
  buffer.WriteString(style)
  buffer.WriteString(text)
  buffer.WriteString(closing)

  return buffer.String()
}
