package asciiart

import (
	"bufio"
	"os"
	"regexp"
)

type font struct {
	Name     string
	Banners  [95]*Banner
	Height   int
	MinWidth int
	MaxWidth int
}

type Banner struct {
	Width int
	View  [8]string
}

func (fnt *font) InitFont() {
	filename := "themes/" + fnt.Name + ".txt"
	file, err := os.Open(filename)
	fnt.MinWidth = MAX_INT
	if err != nil {
		_CloseProgram(err)
	}
	defer file.Close()
	charIdx, lineInd := 0, -1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() && charIdx < len(fnt.Banners) {
		if lineInd == -1 {
			// Initialize Banners
			fnt.Banners[charIdx] = &Banner{}
		} else if lineInd > -1 {
			line := scanner.Text()
			fnt.Banners[charIdx].View[lineInd] = line
			width := len(line)

			fnt.Banners[charIdx].Width = width
			fnt.MinWidth = getMin(fnt.MinWidth, width)
			fnt.MaxWidth = getMax(fnt.MaxWidth, width)
		}
		lineInd++
		if lineInd == 8 {
			lineInd = -1
			charIdx++
		}
	}
	fnt.Height = 8
}

// return args + themename
func getFontNameByParam(args []string) ([]string, string) {
	themeName := "standard"
	files := getFilesFromDir("themes", ".txt")
	cntArgs := len(args)
	for i := 0; i < cntArgs; i++ {
		fName := args[i]
		if files[fName] {
			themeName = fName
			// Removing Params
			tempArgs := append(make([]string, 0), args[:i]...)
			args = append(tempArgs, args[i+1:]...)
			cntArgs, i = cntArgs-1, i-1
		}
	}
	return args, themeName
}

// Getting All Files with extention
func getFilesFromDir(dirPath, extention string) map[string]bool {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		_CloseProgram(err)
	}
	// Getting All Files with extention
	pattern := regexp.MustCompile(`(.{1,})` + extention + "$")
	result := make(map[string]bool, len(files))
	for _, file := range files {
		fName := file.Name()
		isMatched := pattern.MatchString(fName)
		if isMatched {
			fName = pattern.FindStringSubmatch(fName)[1]
			result[fName] = true
		}
	}
	return result
}
