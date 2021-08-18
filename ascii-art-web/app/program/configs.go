package asciiartweb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type gconfigs struct {
	HostURL string `json:"HostURL"`
}

// PathGConfigs - path to GlobalConfigs json
const PathGConfigs = "app/configs.json"

// SetJsonConfigs -----------------------------
func SetJsonConfigs() {
	if gConfigs == nil {
		gConfigs = &gconfigs{}
	}
	file, err := os.Open(PathGConfigs)
	if err != nil {
		if err == os.ErrNotExist {
			CreateJsonConfigs()
			return
		}
	}
	defer file.Close()

	jsonData, _ := ioutil.ReadAll(file)
	isValid := json.Valid(jsonData)
	if !isValid {
		CreateJsonConfigs()
		return
	}
	if json.Unmarshal(jsonData, gConfigs) != nil {
		CreateJsonConfigs()
		return
	}
	ShowWarningMessage("Setted Json Configs for GConfigs")
}

// CreateJsonConfigs ----------------------------
func CreateJsonConfigs() {
	f, err := os.Create(PathGConfigs)
	if err != nil {
		ShowErrorMessage(err.Error())
		os.Exit(1)
	}
	defer f.Close()

	jsonData, err := json.Marshal(gConfigs)
	f.Write(jsonData)
	ShowWarningMessage(fmt.Sprintf("Created '%v' for GlobalConfigs", PathGConfigs))
}
