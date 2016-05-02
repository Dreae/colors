/*
Colors provides a simple way to print colored messages to the console.

Simply wrap a string in a call to the function for the desired color:
  greenText := colors.Green("This is green!")

This will return strings that have been wrapped in ANSI codes to start and stop
the desired color. Printing these strings to the console will result the
appropriate colors on supported terminals.

All colors are also available as background colors:
  greenBackground := colors.BgGreen("The background is now green!")

*/
package colors
