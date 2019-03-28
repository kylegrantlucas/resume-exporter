package models

type Resume struct {
	Basics struct {
		Name     string `json:"name,omitempty"`
		Label    string `json:"label,omitempty"`
		Picture  string `json:"picture,omitempty"`
		Email    string `json:"email,omitempty"`
		Phone    string `json:"phone,omitempty"`
		Website  string `json:"website,omitempty"`
		Summary  string `json:"summary,omitempty"`
		Location struct {
			Address     string `json:"address,omitempty"`
			PostalCode  string `json:"postalCode,omitempty"`
			City        string `json:"city,omitempty"`
			CountryCode string `json:"countryCode,omitempty"`
			Region      string `json:"region,omitempty"`
		} `json:"location,omitempty"`
		Profiles []struct {
			Network  string `json:"network,omitempty"`
			Username string `json:"username,omitempty"`
			URL      string `json:"url,omitempty"`
		} `json:"profiles,omitempty"`
	} `json:"basics,omitempty"`
	Work []struct {
		Company    string   `json:"company,omitempty"`
		Location   string   `json:"location,omitempty"`
		Position   string   `json:"position,omitempty"`
		Website    string   `json:"website,omitempty"`
		StartDate  string   `json:"startDate,omitempty"`
		EndDate    string   `json:"endDate,omitempty"`
		Summary    string   `json:"summary,omitempty"`
		Highlights []string `json:"highlights,omitempty"`
	} `json:"work,omitempty"`
	Volunteer []struct {
		Organization string   `json:"organization,omitempty"`
		Position     string   `json:"position,omitempty"`
		Website      string   `json:"website,omitempty"`
		StartDate    string   `json:"startDate,omitempty"`
		EndDate      string   `json:"endDate,omitempty"`
		Summary      string   `json:"summary,omitempty"`
		Highlights   []string `json:"highlights,omitempty"`
	} `json:"volunteer,omitempty"`
	Education []struct {
		Institution string   `json:"institution,omitempty"`
		Location    string   `json:"location,omitempty"`
		GPA         string   `json:"gpa,omitempty"`
		Area        string   `json:"area,omitempty"`
		StudyType   string   `json:"studyType,omitempty"`
		StartDate   string   `json:"startDate,omitempty"`
		EndDate     string   `json:"endDate,omitempty"`
		Courses     []string `json:"courses,omitempty"`
	} `json:"education,omitempty"`
	Awards []struct {
		Title   string `json:"title,omitempty"`
		Date    string `json:"date,omitempty"`
		Awarder string `json:"awarder,omitempty"`
		Summary string `json:"summary,omitempty"`
	} `json:"awards,omitempty"`
	Publications []struct {
		Name        string `json:"name,omitempty"`
		Publisher   string `json:"publisher,omitempty"`
		ReleaseDate string `json:"releaseDate,omitempty"`
		Website     string `json:"website,omitempty"`
		Summary     string `json:"summary,omitempty"`
	} `json:"publications,omitempty"`
	Projects []struct {
		Name        string `json:"name,omitempty"`
		ReleaseDate string `json:"releaseDate,omitempty"`
		Website     string `json:"website,omitempty"`
		Summary     string `json:"summary,omitempty"`
	} `json:"projects,omitempty"`
	Skills []struct {
		Name     string   `json:"name,omitempty"`
		Level    string   `json:"level,omitempty"`
		Keywords []string `json:"keywords,omitempty"`
	} `json:"skills,omitempty"`
	Languages []struct {
		Language string `json:"language,omitempty"`
		Fluency  string `json:"fluency,omitempty"`
	} `json:"languages,omitempty"`
	Interests []struct {
		Name     string   `json:"name,omitempty"`
		Keywords []string `json:"keywords,omitempty"`
	} `json:"interests,omitempty"`
	References []struct {
		Name      string `json:"name,omitempty"`
		Reference string `json:"reference,omitempty"`
	} `json:"references,omitempty"`
}
