# colors
Simple go ANSI colors library

# Usage

Simply wrap in a call to the function for the desired color:

```Go
greenText := colors.Green("This is green!")
```

Available colors:
  * Black
  * Red
  * Green
  * Yellow
  * Blue
  * Magenta
  * Cyan
  * White

All colors are also available as background colors:

```Go
greenBackground := colors.BgGreen("The background is now green!")
```

Additional styles available:
  * Italic
  * Bold
  * Strikethrough
  * Underline
