package asciiart

import "errors"

type Configs struct {
	ThemeArg   bool
	Theme      string
	ColorArg   bool
	Color      string
	OutputArg  bool
	Output     string
	AlignArg   bool
	Align      string
	ReverseArg bool
	Reverse    string
}

func (c *Configs) ImportTheme(option string) bool {
	switch option {
	case "standard":
		c.Theme = option
		c.ThemeArg = true
		return true
	case "thinkertoy":
		c.Theme = option
		c.ThemeArg = true
		return true
	case "shadow":
		c.Theme = option
		c.ThemeArg = true
		return true
	default:
		return false
	}
}

func (c *Configs) ImportColor(option string) error {
	//todo
	c.ColorArg = true
	c.Color = option
	return nil
}

func (c *Configs) ImportOutput(option string) error {
	//todo
	c.OutputArg = true
	c.Output = option
	return nil
}

func (c *Configs) ImportAlign(option string) error {
	//todo
	c.AlignArg = true
	c.Align = option
	return nil
}

func (c *Configs) ImportReverse(option string) error {
	//todo
	c.ReverseArg = true
	c.Reverse = option
	return nil
}

var (
	ErrNoArgument    = errors.New("Program needs at least one argument")
	ErrNoReceiving   = errors.New("Program needs one received argument (not including options) or \"reverse\" option only")
	ErrManyReceiving = errors.New("Program needs only one received argument")
	ErrSingleReverse = errors.New("Program needs one received argument to print or \"reverse\" option only")
	ErrArgument      = errors.New("Wrong given argument")
	ErrNotHandleArg  = errors.New("This program do not handle that argument")
)
