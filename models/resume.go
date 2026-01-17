package models

// Resume represents the JSONResume v1.0.0 schema with backward compatibility
// for legacy v0.x field names.
type Resume struct {
	Basics struct {
		Name    string `json:"name,omitempty"`
		Label   string `json:"label,omitempty"`
		Image   string `json:"image,omitempty"`   // v1.0.0
		Picture string `json:"picture,omitempty"` // legacy v0.x
		Email   string `json:"email,omitempty"`
		Phone   string `json:"phone,omitempty"`
		URL     string `json:"url,omitempty"`     // v1.0.0
		Website string `json:"website,omitempty"` // legacy v0.x
		Summary string `json:"summary,omitempty"`
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
		Name       string   `json:"name,omitempty"`    // v1.0.0
		Company    string   `json:"company,omitempty"` // legacy v0.x
		Location   string   `json:"location,omitempty"`
		Position   string   `json:"position,omitempty"`
		URL        string   `json:"url,omitempty"`     // v1.0.0
		Website    string   `json:"website,omitempty"` // legacy v0.x
		StartDate  string   `json:"startDate,omitempty"`
		EndDate    string   `json:"endDate,omitempty"`
		Summary    string   `json:"summary,omitempty"`
		Highlights []string `json:"highlights,omitempty"`
	} `json:"work,omitempty"`
	Volunteer []struct {
		Organization string   `json:"organization,omitempty"`
		Position     string   `json:"position,omitempty"`
		URL          string   `json:"url,omitempty"`     // v1.0.0
		Website      string   `json:"website,omitempty"` // legacy v0.x
		StartDate    string   `json:"startDate,omitempty"`
		EndDate      string   `json:"endDate,omitempty"`
		Summary      string   `json:"summary,omitempty"`
		Highlights   []string `json:"highlights,omitempty"`
	} `json:"volunteer,omitempty"`
	Education []struct {
		Institution string   `json:"institution,omitempty"`
		URL         string   `json:"url,omitempty"` // v1.0.0
		Location    string   `json:"location,omitempty"`
		Area        string   `json:"area,omitempty"`
		StudyType   string   `json:"studyType,omitempty"`
		StartDate   string   `json:"startDate,omitempty"`
		EndDate     string   `json:"endDate,omitempty"`
		Score       string   `json:"score,omitempty"` // v1.0.0
		GPA         string   `json:"gpa,omitempty"`   // legacy v0.x
		Courses     []string `json:"courses,omitempty"`
	} `json:"education,omitempty"`
	Awards []struct {
		Title   string `json:"title,omitempty"`
		Date    string `json:"date,omitempty"`
		Awarder string `json:"awarder,omitempty"`
		Summary string `json:"summary,omitempty"`
	} `json:"awards,omitempty"`
	Certificates []struct {
		Name   string `json:"name,omitempty"`
		Date   string `json:"date,omitempty"`
		Issuer string `json:"issuer,omitempty"`
		URL    string `json:"url,omitempty"`
	} `json:"certificates,omitempty"`
	Publications []struct {
		Name        string `json:"name,omitempty"`
		Publisher   string `json:"publisher,omitempty"`
		ReleaseDate string `json:"releaseDate,omitempty"`
		URL         string `json:"url,omitempty"`     // v1.0.0
		Website     string `json:"website,omitempty"` // legacy v0.x
		Summary     string `json:"summary,omitempty"`
	} `json:"publications,omitempty"`
	Projects []struct {
		Name        string   `json:"name,omitempty"`
		Description string   `json:"description,omitempty"` // v1.0.0
		Summary     string   `json:"summary,omitempty"`     // legacy v0.x
		Highlights  []string `json:"highlights,omitempty"`  // v1.0.0
		Keywords    []string `json:"keywords,omitempty"`    // v1.0.0
		StartDate   string   `json:"startDate,omitempty"`   // v1.0.0
		EndDate     string   `json:"endDate,omitempty"`     // v1.0.0
		URL         string   `json:"url,omitempty"`         // v1.0.0
		Website     string   `json:"website,omitempty"`     // legacy v0.x
		Roles       []string `json:"roles,omitempty"`       // v1.0.0
		Entity      string   `json:"entity,omitempty"`      // v1.0.0
		Type        string   `json:"type,omitempty"`        // v1.0.0
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
	Meta struct {
		Canonical    string `json:"canonical,omitempty"`
		Version      string `json:"version,omitempty"`
		LastModified string `json:"lastModified,omitempty"`
	} `json:"meta,omitempty"`
}

// WorkRole represents a single role/position at a company
type WorkRole struct {
	Position   string
	StartDate  string
	EndDate    string
	Summary    string
	Highlights []string
}

// GroupedWork represents a company with potentially multiple roles
type GroupedWork struct {
	Name     string
	Location string
	URL      string
	Roles    []WorkRole
}

// GroupedWorkEntries returns work entries grouped by company name.
// Consecutive entries with the same company name are merged into one entry with multiple roles.
func (r *Resume) GroupedWorkEntries() []GroupedWork {
	if len(r.Work) == 0 {
		return nil
	}

	var grouped []GroupedWork
	var current *GroupedWork

	for _, w := range r.Work {
		companyName := w.Name
		if companyName == "" {
			companyName = w.Company
		}

		role := WorkRole{
			Position:   w.Position,
			StartDate:  w.StartDate,
			EndDate:    w.EndDate,
			Summary:    w.Summary,
			Highlights: w.Highlights,
		}

		// Check if this is the same company as the previous entry
		if current != nil && current.Name == companyName {
			current.Roles = append(current.Roles, role)
		} else {
			// New company - save previous and start new
			if current != nil {
				grouped = append(grouped, *current)
			}
			current = &GroupedWork{
				Name:     companyName,
				Location: w.Location,
				URL:      w.URL,
				Roles:    []WorkRole{role},
			}
		}
	}

	// Don't forget the last one
	if current != nil {
		grouped = append(grouped, *current)
	}

	return grouped
}

// Normalize copies legacy field values to v1.0.0 fields when the new fields are empty.
// This ensures backward compatibility with older JSONResume files.
func (r *Resume) Normalize() {
	// Basics
	if r.Basics.URL == "" && r.Basics.Website != "" {
		r.Basics.URL = r.Basics.Website
	}
	if r.Basics.Image == "" && r.Basics.Picture != "" {
		r.Basics.Image = r.Basics.Picture
	}

	// Work
	for i := range r.Work {
		if r.Work[i].Name == "" && r.Work[i].Company != "" {
			r.Work[i].Name = r.Work[i].Company
		}
		if r.Work[i].URL == "" && r.Work[i].Website != "" {
			r.Work[i].URL = r.Work[i].Website
		}
	}

	// Volunteer
	for i := range r.Volunteer {
		if r.Volunteer[i].URL == "" && r.Volunteer[i].Website != "" {
			r.Volunteer[i].URL = r.Volunteer[i].Website
		}
	}

	// Education
	for i := range r.Education {
		if r.Education[i].Score == "" && r.Education[i].GPA != "" {
			r.Education[i].Score = r.Education[i].GPA
		}
	}

	// Publications
	for i := range r.Publications {
		if r.Publications[i].URL == "" && r.Publications[i].Website != "" {
			r.Publications[i].URL = r.Publications[i].Website
		}
	}

	// Projects
	for i := range r.Projects {
		if r.Projects[i].Description == "" && r.Projects[i].Summary != "" {
			r.Projects[i].Description = r.Projects[i].Summary
		}
		if r.Projects[i].URL == "" && r.Projects[i].Website != "" {
			r.Projects[i].URL = r.Projects[i].Website
		}
	}
}
