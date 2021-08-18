package asciiartweb

import "fmt"

//
func ShowErrorMessage(msg string) {
	fmt.Printf("%sERROR:%s\t%v\n", "\u001b[31;1m", "\033[0m", msg)
}

//
func ShowWarningMessage(msg string) {
	fmt.Printf("%sWARNING:%s%v\n", "\u001b[33;1m", "\033[0m", msg)
}

//
func ShowLogMessage(msg string) {
	fmt.Printf("%sLOG:%s\t%v\n", "\033[38;2;42;183;202m", "\033[0m", msg)
}

//
func ShowSuccesMessage(msg string) {
	fmt.Printf("%sSUCCES:%s\t%v\n", "\u001b[32;1m", "\033[0m", msg)
}
