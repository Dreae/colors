package colors

import "bytes"

// Returns a string wrapped in ANSI start/stop codes for bold
func Bold(in string) string {
  return stylize("\x1b[1m", in, "\x1b[22m")
}

// Returns a string wrapped in ANSI start/stop codes for italic
func Italic(in string) string {
  return stylize("\x1b[3m", in, "\x1b[23m")
}

// Returns a string wrapped in ANSI start/stop codes for underline
func Underline(in string) string {
  return stylize("\x1b[4m", in, "\x1b[24m")
}

// Returns a string wrapped in ANSI start/stop codes to reverse the foreground and background colors
func Inverse(in string) string {
  return stylize("\x1b[7m", in, "\x1b[27m")
}

// Returns a string wrapped in ANSI start/stop codes for strikethrough text
func Strikethrough(in string) string {
  return stylize("\x1b[9m", in, "\x1b[29m")
}

// Returns a string wrapped in ANSI start/stop codes for foreground black
func Black(in string) string {
  return stylize("\x1b[30m", in, "\x1b[39m")
}

// Returns a string wrapped in ANSI start/stop codes for foreground red
func Red(in string) string {
  return stylize("\x1b[31m", in, "\x1b[39m")
}

// Returns a string wrapped in ANSI start/stop codes for foreground green
func Green(in string) string {
  return stylize("\x1b[32m", in, "\x1b[39m")
}

// Returns a string wrapped in ANSI start/stop codes for foreground yellow
func Yellow(in string) string {
  return stylize("\x1b[33m", in, "\x1b[39m")
}

// Returns a string wrapped in ANSI start/stop codes for foreground blue
func Blue(in string) string {
  return stylize("\x1b[34m", in, "\x1b[39m")
}

// Returns a string wrapped in ANSI start/stop codes for foreground magenta
func Magenta(in string) string {
  return stylize("\x1b[35m", in, "\x1b[39m")
}

// Returns a string wrapped in ANSI start/stop codes for foreground cyan
func Cyan(in string) string {
  return stylize("\x1b[36m", in, "\x1b[39m")
}

// Returns a string wrapped in ANSI start/stop codes for foreground white
func White(in string) string {
  return stylize("\x1b[37m", in, "\x1b[39m")
}

// Returns a string wrapped in ANSI start/stop codes for background black
func BgBlack(in string) string {
  return stylize("\x1b[40m", in, "\x1b[49m")
}

// Returns a string wrapped in ANSI start/stop codes for background red
func BgRed(in string) string {
  return stylize("\x1b[41m", in, "\x1b[49m")
}

// Returns a string wrapped in ANSI start/stop codes for background green
func BgGreen(in string) string {
  return stylize("\x1b[42m", in, "\x1b[49m")
}

// Returns a string wrapped in ANSI start/stop codes for background yellow
func BgYellow(in string) string {
  return stylize("\x1b[43m", in, "\x1b[49m")
}

// Returns a string wrapped in ANSI start/stop codes for background blue
func BgBlue(in string) string {
  return stylize("\x1b[44m", in, "\x1b[49m")
}

// Returns a string wrapped in ANSI start/stop codes for background magenta
func BgMagenta(in string) string {
  return stylize("\x1b[45m", in, "\x1b[49m")
}

// Returns a string wrapped in ANSI start/stop codes for background cyan
func BgCyan(in string) string {
  return stylize("\x1b[46m", in, "\x1b[49m")
}

// Returns a string wrapped in ANSI start/stop codes for background white
func BgWhite(in string) string {
  return stylize("\x1b[47m", in, "\x1b[49m")
}

// Wraps string in provided ANSI codes
func stylize(style string, text string, closing string) string {
  var buffer bytes.Buffer
  buffer.WriteString(style)
  buffer.WriteString(text)
  buffer.WriteString(closing)

  return buffer.String()
}
