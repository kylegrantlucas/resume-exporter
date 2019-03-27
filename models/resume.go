package models

type Resume struct {
	Basics struct {
		Name     string `json:"name"`
		Label    string `json:"label"`
		Picture  string `json:"picture"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Website  string `json:"website"`
		Summary  string `json:"summary"`
		Location struct {
			Address     string `json:"address"`
			PostalCode  string `json:"postalCode"`
			City        string `json:"city"`
			CountryCode string `json:"countryCode"`
			Region      string `json:"region"`
		} `json:"location"`
		Profiles []struct {
			Network  string `json:"network"`
			Username string `json:"username"`
			URL      string `json:"url"`
		} `json:"profiles"`
	} `json:"basics"`
	Work []struct {
		Company    string   `json:"company"`
		Location   string   `json:"location"`
		Position   string   `json:"position"`
		Website    string   `json:"website"`
		StartDate  string   `json:"startDate"`
		EndDate    string   `json:"endDate"`
		Summary    string   `json:"summary"`
		Highlights []string `json:"highlights"`
	} `json:"work"`
	Volunteer []struct {
		Organization string   `json:"organization"`
		Position     string   `json:"position"`
		Website      string   `json:"website"`
		StartDate    string   `json:"startDate"`
		EndDate      string   `json:"endDate"`
		Summary      string   `json:"summary"`
		Highlights   []string `json:"highlights"`
	} `json:"volunteer"`
	Education []struct {
		Institution string   `json:"institution"`
		Location    string   `json:"location"`
		GPA         string   `json:"gpa"`
		Area        string   `json:"area"`
		StudyType   string   `json:"studyType"`
		StartDate   string   `json:"startDate"`
		EndDate     string   `json:"endDate"`
		Gpa         string   `json:"gpa"`
		Courses     []string `json:"courses"`
	} `json:"education"`
	Awards []struct {
		Title   string `json:"title"`
		Date    string `json:"date"`
		Awarder string `json:"awarder"`
		Summary string `json:"summary"`
	} `json:"awards"`
	Publications []struct {
		Name        string `json:"name"`
		Publisher   string `json:"publisher"`
		ReleaseDate string `json:"releaseDate"`
		Website     string `json:"website"`
		Summary     string `json:"summary"`
	} `json:"publications"`
	Projects []struct {
		Name        string `json:"name"`
		ReleaseDate string `json:"releaseDate"`
		Website     string `json:"website"`
		Summary     string `json:"summary"`
	} `json:"projects"`
	Skills []struct {
		Name     string   `json:"name"`
		Level    string   `json:"level"`
		Keywords []string `json:"keywords"`
	} `json:"skills"`
	Languages []struct {
		Language string `json:"language"`
		Fluency  string `json:"fluency"`
	} `json:"languages"`
	Interests []struct {
		Name     string   `json:"name"`
		Keywords []string `json:"keywords"`
	} `json:"interests"`
	References []struct {
		Name      string `json:"name"`
		Reference string `json:"reference"`
	} `json:"references"`
}
