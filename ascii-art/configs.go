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

var (
	ErrNoArgument    = errors.New("Program needs at least one argument")
	ErrNoReceiving   = errors.New("Program needs one received argument (not including options) or \"reverse\" option only")
	ErrManyReceiving = errors.New("Program needs only one received argument")
	ErrSingleReverse = errors.New("Program needs one received argument to print or \"reverse\" option only")
	ErrArgument      = errors.New("Wrong given argument")
	ErrConflictArgs  = errors.New("Given options are conflicting")
	ErrNotHandleArg  = errors.New("This program do not handle that argument")
	ErrWrongSymbol   = errors.New("There are unexpected symbol")
	ErrInvalidPath   = errors.New("Given path is invalid")
	ErrInvalidFile   = errors.New("Given file name or path is incorrect")
)
