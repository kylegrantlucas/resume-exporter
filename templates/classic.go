package templates

var Classic = `// Classic Resume Template (Typst)
// Based on Sourabh Bajaj's LaTeX template
// Ported to Typst by Kyle Lucas

#set document(title: "{{.Basics.Name}} - Resume")
#set page(
  paper: "us-letter",
  margin: (x: 0.5in, y: 0.5in),
)
#set text(font: ("Avenir Next", "Helvetica Neue", "Arial"), size: 10pt)
#set par(justify: false, leading: 0.65em)
#show list: set list(spacing: 0.65em)

// Section title with underline
#let section-title(title) = {
  v(4pt)
  text(weight: "bold", size: 1.1em, smallcaps(title))
  v(-2pt)
  line(length: 100%, stroke: 0.5pt + black)
  v(-5pt)
}

// Entry with title/location on top, position/dates below
#let resume-subheading(title, location, position, dates) = {
  v(2pt)
  grid(
    columns: (1fr, auto),
    text(weight: "bold")[#title],
    text()[#location]
  )
  grid(
    columns: (1fr, auto),
    text(style: "italic", size: 0.9em)[#position],
    text(style: "italic", size: 0.9em)[#dates]
  )
  v(2pt)
}

// Company heading (for grouped roles)
#let company-heading(name, location) = {
  v(2pt)
  grid(
    columns: (1fr, auto),
    text(weight: "bold")[#name],
    text()[#location]
  )
}

// Role under a company
#let role-entry(position, dates) = {
  grid(
    columns: (1fr, auto),
    text(style: "italic", size: 0.9em)[#position],
    text(style: "italic", size: 0.9em)[#dates]
  )
}

// Simple subitem for skills/projects
#let resume-subitem(name, description) = {
  [*#name*: #description]
  v(-4pt)
}

// Configure default list style
#set list(marker: [#sym.circle.small], tight: true, spacing: 2pt)

// ============ DOCUMENT STARTS HERE ============

// Header
#grid(
  columns: (1fr, auto),
  [
    #text(size: 1.5em, weight: "bold")[
      #link("{{.Basics.URL}}")[{{.Basics.Name}}]
    ]
    #linebreak()
    #link("{{.Basics.URL}}")[{{.Basics.URL}}]
  ],
  align(right)[
    Email: #link("mailto:{{.Basics.Email}}")[{{ .Basics.Email | escape }}]
    {{if ne .Basics.Phone ""}}\ Phone: {{.Basics.Phone}}{{end}}
  ]
)

// ============ EXPERIENCE ============
#section-title("Experience")

{{ range $companyIdx, $company := .GroupedWorkEntries }}
#company-heading("{{$company.Name}}", "{{$company.Location}}")
{{ range $roleIdx, $role := $company.Roles }}
#role-entry(
  "{{$role.Position}}",
  "{{$role.StartDate | date}} - {{if eq $role.EndDate ""}}Present{{else}}{{$role.EndDate | date}}{{end}}"
)
{{ if $role.Highlights }}
{{ range $itemKey, $itemValue := $role.Highlights }}
- {{$itemValue}}
{{ end }}
{{ end }}
{{ end }}
{{ end }}

// ============ EDUCATION ============
#section-title("Education")

{{ range $key, $value := .Education }}
#resume-subheading(
  "{{$value.Institution}}",
  "{{$value.Location}}",
  "{{if ne $value.StudyType ""}}{{$value.StudyType}} in {{end}}{{$value.Area}}{{if ne $value.Score ""}}; GPA: {{$value.Score}}{{end}}",
  "{{if ne $value.StartDate ""}}{{$value.StartDate | date}} - {{end}}{{if eq $value.EndDate ""}}Present{{else}}{{$value.EndDate | date}}{{end}}"
)
{{ end }}

// ============ PROJECTS ============
{{if .Projects}}
#section-title("Projects")

{{ range $key, $value := .Projects }}
#resume-subitem("{{$value.Name}}", "{{$value.Description}}")
{{ end }}
{{end}}

// ============ VOLUNTEER WORK ============
{{if .Volunteer}}
#section-title("Volunteer Work")

{{ range $key, $value := .Volunteer }}
#resume-subheading(
  "{{$value.Organization}}",
  "",
  "{{$value.Position}}",
  "{{if ne $value.StartDate ""}}{{$value.StartDate | date}} - {{end}}{{if eq $value.EndDate ""}}Present{{else}}{{$value.EndDate | date}}{{end}}"
)
{{ end }}
{{end}}

// ============ AWARDS ============
{{if .Awards}}
#section-title("Awards")

{{ range $key, $value := .Awards }}
#resume-subheading(
  "{{$value.Title}}",
  "",
  "{{$value.Awarder}}",
  "{{$value.Date | date}}"
)
{{ end }}
{{end}}

// ============ SKILLS ============
{{if .Skills}}
#section-title("Skills")

{{ range $key, $value := .Skills }}
- *{{$value.Name}}*: {{ range $itemKey, $itemValue := .Keywords }}{{if $itemKey}}, {{end}}{{$itemValue}}{{ end }}
{{ end }}
{{end}}
`
