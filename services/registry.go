package services

import (
	"fmt"
	"net/url"
	"time"

	"github.com/nickduskey/dfpm/utils"
)

type Registry struct {
	URL url.URL
}

type User struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	URL   string `json:"url,omitempty"`
}

type Repository struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type Bugs struct {
	URL string `json:"url"`
}

type Distribution struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type License struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type PackageVersionEntry struct {
	Name                 string            `json:"name"`
	Version              string            `json:"version"`
	Description          string            `json:"description"`
	Homepage             string            `json:"homepage"`
	Main                 string            `json:"main"`
	Keywords             []string          `json:"keywords"`
	Licenses             []License         `json:"licenses"`
	Author               User              `json:"author"`
	Bugs                 Bugs              `json:"bugs"`
	Repository           Repository        `json:"repository"`
	Engines              []string          `json:"engines"`
	Directories          map[string]string `json:"directories"`
	NPMUser              User              `json:"_npmUser"`
	ID                   string            `json:"_id"`
	Dependencies         map[string]string `json:"dependencies"`
	DevDependencies      map[string]string `json:"devDependencies"`
	OptionalDependencies map[string]string `json:"OptionalDependencies"`
	EngineSupported      bool              `json:"_engineSupported"`
	NPMVersion           string            `json:"_npmVersion"`
	NodeVersion          string            `json:"_nodeVersion"`
	DefaultsLoaded       bool              `json:"_defaultsLoaded"`
	Dist                 Distribution      `json:"dist"`
	Maintainers          []User            `json:"maintainers"`
}

type PackageEntry struct {
	ID             string                         `json:"_id"`
	Rev            string                         `json:"_rev"`
	Name           string                         `json:"name"`
	Description    string                         `json:"description"`
	DistTags       map[string]string              `json:"dist-tags"`
	Versions       map[string]PackageVersionEntry `json:"versions"`
	Readme         string                         `json:"readme"`
	Maintainers    []User                         `json:"maintainers"`
	Time           map[string]time.Time           `json:"time"`
	Author         User                           `json:"author"`
	Repository     Repository                     `json:"repository"`
	Users          map[string]bool                `json:"users"`
	ReadmeFilename string                         `json:"readmeFilename"`
	Homepage       string                         `json:"homepage"`
	Keywords       []string                       `json:"keywords"`
	Contributors   []User                         `json:"contributors"`
	Bugs           Bugs                           `json:"bugs"`
	License        string                         `json:"license"`
}

func (r *Registry) GetPackageEntry(packageName, version string) (PackageEntry, error) {
	url := fmt.Sprintf("https://registry.npmjs.org/%s", packageName)
	var packageEntry PackageEntry
	err := utils.GetRemoteJSON(url, &packageEntry)
	if err != nil {
		return nil, err
	}
	var version PackageVersionEntry
}
