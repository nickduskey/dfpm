package services

import "time"

type NpmsLinks struct {
	Npm        string `json:"npm"`
	Homepage   string `json:"homepage"`
	Repository string `json:"repository"`
	Bugs       string `json:"bugs"`
}

type NpmsAuthor struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type NpmsUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type NpmsPackage struct {
	Name        string     `json:"name"`
	Scope       string     `json:"scope"`
	Version     string     `json:"version"`
	Description string     `json:"description"`
	Keywords    []string   `json:"keywords"`
	Date        time.Time  `json:"date"`
	Links       NpmsLinks  `json:"links"`
	Publisher   NpmsUser   `json:"publisher"`
	Maintainers []NpmsUser `json:"maintainers"`
}

type NpmsFlags struct {
	Deprecated string `json:"deprecated"`
	Unstable   string `json:"unstable"`
	Insecure   string `json:"insecure"`
}

type NpmsScoreDetail struct {
	Quality     float64 `json:"quality"`
	Popularity  float64 `json:"popularity"`
	Maintenance float64 `json:"maintenance"`
}
type NpmsScore struct {
	Final  float64         `json:"final"`
	Detail NpmsScoreDetail `json:"detail"`
}

type NpmsSearchResult struct {
	Package     NpmsPackage `json:"package"`
	Flags       NpmsFlags   `json:"flags"`
	Score       NpmsScore   `json:"score"`
	SearchScore float64     `json:"searchScore"`
}

type NpmsSearchResponse struct {
	Total   int                `json:"total"`
	Results []NpmsSearchResult `json:"results"`
}

type NpmsRelease struct {
	From  time.Time `json:"from"`
	To    time.Time `json:"to"`
	Count int       `json:"count"`
}

type NpmsMetadata struct {
	NpmsPackage
	License      string            `json:"license"`
	Dependencies map[string]string `json:"dependencies"`
	Releases     interface{}       `json:"releases"`
}

type NpmsNpm struct {
	Downloads interface{} `json:"downloads"`
}

type NpmsGithubIssues struct {
	Count        int         `json:"count"`
	OpenCount    int         `json:"openCount"`
	Distribution map[int]int `json:"distribution"`
	IsDisabled   bool        `json:"isDisabled"`
}

type NpmsGithubContributor struct {
	Username     string `json:"username"`
	CommitsCount int    `json:"commitsCount"`
}

type NpmsRange struct {
	From  time.Time `json:"from"`
	To    time.Time `json:"to"`
	Count int       `json:"count"`
}

type NpmsGithubStatus struct {
	Context string `json:"context"`
	State   string `json:"state"`
}

type NpmsGithub struct {
	Homepage         string                  `json:"homepage"`
	StarsCount       int                     `json:"starsCount"`
	ForksCount       int                     `json:"forksCount"`
	SubscribersCount int                     `json:"subscribersCount"`
	Issues           NpmsGithubIssues        `json:"issues"`
	Contributors     []NpmsGithubContributor `json:"contributors"`
	Commits          []NpmsRange             `json:"commits"`
	Statuses         []NpmsGithubStatus      `json:"statuses"`
}

type NpmsFiles struct {
	ReadmeSize   int  `json:"readmeSize"`
	TestsSize    int  `json:"testsSize"`
	HasChangelog bool `json:"hasChangelog"`
}

type NpmsSource struct {
	Files    NpmsFiles `json:"files"`
	Linters  []string  `json:"linters"`
	Coverage float64   `json:"coverage"`
}

type NpmsCollected struct {
	Metadata NpmsMetadata `json:"metadata"`
	Npm      NpmsNpm      `json:"npm"`
	Github   NpmsGithub   `json:"github"`
	Source   NpmsSource   `json:"source"`
}

type NpmsQuality struct {
	Carefulness float64 `json:"carefulness"`
	Tests       float64 `json:"tests"`
	Health      float64 `json:"health"`
	Branding    float64 `json:"branding"`
}

type NpmsPopularity struct {
	CommunityInterest     float64 `json:"communityInterest"`
	DownloadsCount        float64 `json:"downloadsCount"`
	DownloadsAcceleration float64 `json:"downloadsAcceleration"`
	DependentsCount       float64 `json:"dependentsCount"`
}

type NpmsMaintenance struct {
	ReleasesFrequency  float64 `json:"releasesFrequency"`
	CommitsFrequency   float64 `json:"commitsFrequency"`
	OpenIssues         float64 `json:"openIssues"`
	IssuesDistribution float64 `json:"issuesDistribution"`
}

type NpmsEvaluation struct {
	Quality     NpmsQuality     `json:"quality"`
	Popularity  NpmsPopularity  `json:"popularity"`
	Maintenance NpmsMaintenance `json:"maintenance"`
}

type NpmsPackageResponse struct {
	AnalyzedAt time.Time      `json:"analyzedAt"`
	Collected  NpmsCollected  `json:"collected"`
	Evaluation NpmsEvaluation `json:"evaluation"`
	Score      NpmsScore      `json:"score"`
}
