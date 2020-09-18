package services

import (
	"encoding/json"
	"io/ioutil"
)

type FileService struct{}

type PackageJSON struct {
	Name                 string      `json:"name"`
	Version              string      `json:"version"`
	Description          string      `json:"description"`
	Keywords             []string    `json:"keywords"`
	Homepage             interface{} `json:"homepage"`
	Bugs                 interface{} `json:"bugs"`
	License              string      `json:"license"`
	Licenses             interface{} `json:"licenses"`
	Author               interface{} `json:"author"`
	Contributors         interface{} `json:"contributors"`
	Maintainers          interface{} `json:"maintainers"`
	Files                []string    `json:"files"`
	Main                 string      `json:"main"`
	Bin                  interface{} `json:"bin"`
	Types                string      `json:"types"`
	Typings              string      `json:"typings"`
	Man                  []string    `json:"man"`
	Directories          interface{} `json:"directories"`
	Repository           interface{} `json:"repository"`
	Scripts              interface{} `json:"scripts"`
	Config               interface{} `json:"config"`
	Dependencies         interface{} `json:"dependencies"`
	DevDependencies      interface{} `json:"devDependencies"`
	OptionalDependencies interface{} `json:"optionalDependencies"`
	PeerDependencies     interface{} `json:"peerDependencies"`
	Resolutions          interface{} `json:"resolutions"`
	Engines              interface{} `json:"engines"`
	Resolutions          interface{} `json:"resolutions"`
	Engines              interface{} `json:"engines"`
	EngineStrict         bool        `json:"engineStrict"`
	OS                   []string    `json:"os"`
	CPU                  []string    `json:"cpu"`
	PreferGlobal         bool        `json:"preferGlobal"`
	Private              bool        `json:"private"`
	PublishConfig        interface{} `json:"publishConfig"`
	Dist                 interface{} `json:"dist"`
	Readme               string      `json:"readme"`
	Module               string      `json:"module"`
	Esnext               interface{} `json:"esnext"`
	Workspaces           []string    `json:"workspaces"`
}

func (f *FileService) WriteJSONToFile(filePath string, data interface{}) error {
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filePath, file, 0644)
	return err
}

func (f *FileService) ReadPackageJSON() (PackageJSON, error) {

}
