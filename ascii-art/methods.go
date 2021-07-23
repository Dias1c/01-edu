package asciiart

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

func (c *Configs) Conflict() bool {
	if c.ReverseArg && (c.AlignArg || c.OutputArg || c.ColorArg) {
		return true
	} else if c.OutputArg && (c.AlignArg || c.ReverseArg || c.ColorArg) {
		return true
	}
	return false
}

func (c *Configs) ConflictHint() string {
	if c.ReverseArg && (c.AlignArg || c.OutputArg || c.ColorArg) {
		return "reverse and (align or output or color)"
	} else if c.OutputArg && (c.AlignArg || c.ReverseArg || c.ColorArg) {
		return "output and (align or reverse or color)"
	}
	return ""
}

func (c *Configs) ImportColor(option string) error {
	//todo
	c.ColorArg = true
	c.Color = option
	return nil
}

func (c *Configs) ImportOutput(option string) error {
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
