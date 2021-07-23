package asciiart

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var GBanner [95][8]string

// Import Banner From File to GBanner
func ImportBanner() error {
	filename := "themes/" + GConfigs.Theme + ".txt"
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	charInd, lineInd := 0, -1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if lineInd > -1 {
			GBanner[charInd][lineInd] = scanner.Text()
		}
		lineInd++
		if lineInd == 8 {
			lineInd = -1
			charInd++
		}
	}
	return nil
}

// Get GR of Received String and Return It
func GetGraphicRepresentation(received string) ([][]string, error, string) {
	var gr [][]string
	cnt := strings.Count(received, "\n") + 1
	gr = make([][]string, cnt)
	// Handle Empty String
	if len(received) < 1 {
		return gr, nil, ""
	}
	// Split String in New Lines and Representate It
	splited := strings.Split(received, "\n")
	for i := 0; i < cnt; i++ {
		// If Splited Text is Empty
		if len(splited[i]) < 1 {
			gr[i] = []string{""}
			continue
		}
		// Use strings.Builder
		gr[i] = make([]string, 8)
		var builder [8]strings.Builder
		for _, symbol := range splited[i] {
			if symbol < ' ' || symbol > '~' {
				return gr, ErrWrongSymbol, string(symbol)
			}
			gbInd := int(symbol - ' ')
			for j := 0; j < 8; j++ {
				builder[j].WriteString(GBanner[gbInd][j])
			}
		}
		// Fill gr from builder and Check Ending Space
		endWithSpace := strings.HasSuffix(splited[i], " ")
		for j := 0; j < 8; j++ {
			gr[i][j] = builder[j].String()
			if !endWithSpace {
				gr[i][j] = strings.TrimRight(gr[i][j], " ")
			}
		}
	}
	return gr, nil, ""
}

// Replace "\n" to New Line in String
func GetFormattedString(str string) string {
	return strings.ReplaceAll(str, "\\n", "\n")
}

// Simple Print GR
// to do for align and color
func PrintGraphicRepresentation(gr [][]string) {
	for i := range gr {
		for j := range gr[i] {
			fmt.Println(gr[i][j])
		}
	}
}
