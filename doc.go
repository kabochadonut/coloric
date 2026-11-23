/*
Package coloric is an ANSI color package that streamlines the styling of text in the terminal.
Coloric is designed to be light and extensible.

Some interesting functions to decorate text

			// Rainbow using the terminals theme
			func Rainbow(s string) string {
				ss := make([]string, len(s))
				for i, char := range s {
					cc := coloric.New(coloric.FgRed + coloric.Attribute(i%6))
					ss[i] = string(char) + cc.SetFormat()
				}
				return strings.Join(ss, "") + coloric.ResetFormat()
			}

			// Rainbow using ANSI colors
			func RainbowAnsi(s string) string {
				colors := []int{
					196,
					208,
					226,
					118,
					27,
					57,
					165,
				}
				ss := make([]string, len(s))
				for i, char := range s {
					cc := coloric.New().Fg(coloric.Attribute(colors[i%len(colors)]))
					ss[i] = cc.SetFormat() + string(char)
				}
				return strings.Join(ss, "") + coloric.ResetFormat()
			}

			// Covered until highlighted
			func Spoiler(a coloric.Attribute) *coloric.Color {
				c := &coloric.Color{
					Params: []coloric.Attribute{38, 5, a, 48, 5, a},
				}
				return c
			}

			// Covered even when highlighted but not when copy/pasted
			func Redacted() *coloric.Color {
				c := &coloric.Color{
					Params: []coloric.Attribute{7, 8},
				}
				return c
			}

		// Create a map of colors where the key is the name and the value is the Color object formatted to a string
		colors := map[string]string{}

		colors["Success"] = coloric.New(
			coloric.FgGreen,
			coloric.Bold,
			coloric.Reverse,
		).String(" SUCCESS ")

		colors["Error"] = coloric.New(
			coloric.FgHiRed,
			coloric.Bold,
			coloric.BlinkFast,
			coloric.Reverse,
		).String(" ERROR ")

		// Generate static formatted strings for reuse anywhere in your project
		result := "const (\n"
		for key, val := range colors {
			colorStr := fmt.Sprintf("%+q", val)
			result += fmt.Sprintf("\t%s = %s\n", key, colorStr)
		}
		result += ")\n"
		fmt.Println(result)
		// Output:
		const (
	        Error = "\x1b[91;1;6;7m ERROR \x1b[0;22;25;27m"
	        Success = "\x1b[32;1;7m SUCCESS \x1b[0;22;27m"
		)

		// Print all 256 ANSI colors in a grid
		for i := range 256 {
			c := coloric.New().Fg(coloric.Attribute(i)).String(fmt.Sprintf("%03d ", i))
			fmt.Printf(c)
			if (i+1)%8 == 0 {
				fmt.Printf("\n")
			}
		}
*/
package coloric
