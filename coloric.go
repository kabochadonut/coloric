package coloric

import (
	"fmt"
	"strconv"
	"strings"
)

func (c *Color) sequence() string {
	format := make([]string, len(c.Params))
	for i, v := range c.Params {
		format[i] = strconv.Itoa(int(v))
	}
	return strings.Join(format, ";")
}

func (c *Color) wrap(s string) string {
	return c.format() + s + c.unformat()
}

func (c *Color) format() string {
	return fmt.Sprintf("%s[%sm", escape, c.sequence())
}

func (c *Color) unformat() string {
	format := make([]string, len(c.Params))
	for i, v := range c.Params {
		format[i] = strconv.Itoa(int(Reset))
		ra, ok := mapResetAttributes[v]
		if ok {
			format[i] = strconv.Itoa(int(ra))
		}
	}
	return fmt.Sprintf("%s[%sm", escape, strings.Join(format, ";"))
}

func (c *Color) SetFormat() string {
	return c.format()
}

func (c *Color) UnsetFormat() string {
	return c.unformat()
}

func ResetFormat() string {
	return fmt.Sprintf("%s[%d", escape, Reset)
}

func New(value ...Attribute) *Color {
	c := &Color{
		Params: make([]Attribute, 0),
	}
	c.Add(value...)
	return c
}

func (c *Color) Add(value ...Attribute) *Color {
	c.Params = append(c.Params, value...)
	return c
}

func (c *Color) Fg(value Attribute) *Color {
	c.Add(38, 5, value)
	return c
}

func (c *Color) Bg(value Attribute) *Color {
	c.Add(48, 5, value)
	return c
}

func (c *Color) SprintFunc() func(a ...any) string {
	return func(a ...any) string {
		return c.wrap(fmt.Sprint(a...))
	}
}

func (c *Color) SprintfFunc() func(format string, a ...any) string {
	return func(format string, a ...any) string {
		return c.wrap(fmt.Sprintf(format, a...))
	}
}

func (c *Color) SprintlnFunc() func(a ...any) string {
	return func(a ...any) string {
		return c.wrap(fmt.Sprintln(a...))
	}
}

func FormatString(formats ...string) string {
	return strings.Join(formats, "")
}

func (c *Color) String(message string) string {
	return c.SprintFunc()(message)
}
