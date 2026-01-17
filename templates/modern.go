package templates

var Modern = `// Modern Resume Template (Typst)
// Clean, contemporary two-column layout
// Ported to Typst by Kyle Lucas

#set document(title: "{{.Basics.Name}} - Resume")
#set page(
  paper: "us-letter",
  margin: (x: 0.35in, y: 0.3in),
)

// Color definitions - softer, modern palette
#let primary = rgb("#1a1a1a")
#let headings = rgb("#555555")
#let subheadings = rgb("#222222")
#let accent = rgb("#0066cc")

// Typography setup - clean sans-serif
#set text(font: ("Avenir Next", "Helvetica Neue", "Arial"), size: 8.25pt, fill: primary)
#set par(justify: false, leading: 0.4em)
#show list: set list(spacing: 0.2em)

// Helper functions
#let section-title(title) = {
  text(fill: accent, size: 9pt, weight: "bold", tracking: 0.05em, upper(title))
  v(-4pt)
  line(length: 100%, stroke: 0.5pt + accent.lighten(60%))
  v(1pt)
}

#let subsection-title(title) = {
  text(fill: subheadings, weight: "semibold", size: 8.5pt)[#title]
  linebreak()
}

#let run-subsection(title) = {
  text(fill: subheadings, weight: "bold", size: 9pt)[#title]
}

#let descript(content) = {
  text(fill: headings, size: 7.5pt, style: "italic")[#content]
}

#let location(content) = {
  text(fill: headings, size: 7.5pt)[#content]
}

#let section-sep = v(2pt)

#let custom-bold(content) = {
  text(fill: accent, weight: "medium")[#content]
}

// Configure default list style
#set list(marker: [‣], tight: true, spacing: 0pt, indent: 0.4em)

// ============ DOCUMENT STARTS HERE ============

// Name Header (full width, centered)
#align(center)[
  #text(size: 20pt, font: "Avenir Next", weight: "regular", fill: primary, tracking: 0.02em)[{{.Basics.Name}}]
  #v(-2pt)
  #text(size: 8pt, fill: headings)[
    #link("{{.Basics.URL}}")[{{.Basics.URL}}]
    #h(4pt) | #h(4pt) #link("mailto:{{.Basics.Email}}")[{{ .Basics.Email | escape }}]
    {{if ne .Basics.Phone ""}}#h(4pt) | #h(4pt) {{.Basics.Phone}}{{end}}
  ]
]
#v(-2pt)
#line(length: 100%, stroke: 0.5pt + primary.lighten(70%))
#v(2pt)

// Two-column layout
#grid(
  columns: (1fr, 2.2fr),
  gutter: 16pt,

  // ======== LEFT COLUMN ========
  [
    // Education
    #section-title("Education")
    {{ range $key, $value := .Education }}
    *{{$value.Institution}}*
    #linebreak()
    #text(size: 7.5pt, fill: headings)[{{$value.Area}}{{if ne $value.StudyType ""}} · {{$value.StudyType}}{{end}} · {{if eq $value.EndDate ""}}Present{{else}}{{$value.EndDate | date}}{{end}}]
    #v(2pt)
    {{ end }}

    // Links
    {{if .Basics.Profiles}}
    #section-title("Links")
    {{ range $key, $value := .Basics.Profiles }}
    #text(size: 7.5pt)[{{$value.Network}}:] #link("{{$value.URL}}")[#text(size: 7.5pt, fill: accent)[{{$value.Username}}]]
    {{ end }}
    #section-sep
    {{end}}

    // Skills - more compact
    {{if .Skills}}
    #section-title("Skills")
    {{ range $key, $value := .Skills }}
    #text(size: 8pt, weight: "bold")[{{$value.Name}}]
    #linebreak()
    #text(size: 7pt, fill: headings)[{{ range $itemKey, $itemValue := .Keywords }}{{if $itemKey}}, {{end}}{{$itemValue}}{{ end }}]
    #v(2pt)
    {{ end }}
    {{end}}

    // Projects - compact
    {{if .Projects}}
    #section-title("Projects")
    {{ range $key, $value := .Projects }}
    #text(size: 7.5pt)[*{{$value.Name}}* – {{$value.Description}}]
    #v(1pt)
    {{ end }}
    {{end}}

    // Other
    {{if or .Volunteer .Awards}}
    #section-title("Other")
    {{if .Awards}}{{ range $key, $value := .Awards }}#text(size: 7.5pt)[*{{$value.Title}}* · {{$value.Awarder}}]
    {{ end }}{{end}}
    {{if .Volunteer}}{{ range $key, $value := .Volunteer }}#text(size: 7.5pt)[{{$value.Organization}} · {{$value.Position}}]
    {{ end }}{{end}}
    {{end}}
  ],

  // ======== RIGHT COLUMN ========
  [
    // Experience
    #section-title("Experience")

    {{ range $companyIdx, $company := .GroupedWorkEntries }}
    #grid(
      columns: (1fr, auto),
      text(fill: subheadings, weight: "bold", size: 9pt)[{{$company.Name}}],
      text(size: 7pt, fill: headings)[{{$company.Location}}]
    )
    {{ range $roleIdx, $role := $company.Roles }}
    #v(-3pt)
    #grid(
      columns: (1fr, auto),
      text(fill: headings, size: 8pt, weight: "medium")[{{$role.Position}}],
      text(size: 7pt, fill: headings)[{{$role.StartDate | date}} – {{if eq $role.EndDate ""}}Present{{else}}{{$role.EndDate | date}}{{end}}]
    )
    {{ if $role.Highlights }}
    {{ range $itemKey, $itemValue := $role.Highlights }}
    - #text(size: 8pt)[{{$itemValue}}]
    {{ end }}
    {{ end }}
    {{ end }}
    #v(2pt)
    {{ end }}
  ]
)
`
