package asciiart

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var GBanner [95][8]string

// import banner from file
func ImportBanner() {
	file, err := os.Open("themes/" + GConfigs.FileTheme)
	if err != nil {
		CloseProgram(err)
	}
	defer file.Close()
	//todo
	charInd, lineInd := 0, -1
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString(byte(10))
		if err != nil {
			if err == io.EOF {
				break
			} else {
				CloseProgram(err)
			}
		}
		if lineInd > -1 {
			GBanner[charInd][lineInd] = str[:len(str)-1]
		}
		lineInd++
		if lineInd == 8 {
			lineInd = -1
			charInd++
		}
	}
}

//Printing Str With Banner
func PrintWithBanner(str string) {
	symbols := []rune(str)
	count := len(symbols)
	checkPoint := 0
	for idx, lineInd := 0, 0; idx <= count && lineInd < 8; idx++ {
		if idx == count || symbols[idx] == '\n' {
			fmt.Println()
			if lineInd == 7 && idx < count && symbols[idx] == '\n' {
				checkPoint, idx = idx+1, idx+1
				for ; idx < count && symbols[idx] == '\n'; checkPoint, idx = idx+1, idx+1 {
					fmt.Print("\n")
				}
				lineInd = 0
			} else if lineInd < 7 {
				lineInd++
				idx = checkPoint
			}
			if idx == count {
				break
			}
		}
		_PrintBannerLine(symbols[idx], lineInd)
	}
}

func _PrintBannerLine(symbol rune, lineInd int) {
	charInd := int(symbol - ' ')
	if lineInd == 8 || symbol < 32 || 126 < symbol {
		return
	} else if -1 < lineInd && lineInd < 8 {
		fmt.Print(GBanner[charInd][lineInd])
	}
}
