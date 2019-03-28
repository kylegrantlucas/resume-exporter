package templates

var Classic = `
%-------------------------
% Resume in Latex
% Author : Sourabh Bajaj (Templating by Kyle Lucas)
% License : MIT
%------------------------

\documentclass[letterpaper,11pt]{article}

\usepackage{latexsym}
\usepackage[empty]{fullpage}
\usepackage{titlesec}
\usepackage{marvosym}
\usepackage[usenames,dvipsnames]{color}
\usepackage{verbatim}
\usepackage{enumitem}
\usepackage[hidelinks]{hyperref}
\usepackage{fancyhdr}
\usepackage[english]{babel}

\pagestyle{fancy}
\fancyhf{} %----clear all header and footer fields
\fancyfoot{}
\renewcommand{\headrulewidth}{0pt}
\renewcommand{\footrulewidth}{0pt}

% Adjust margins
\addtolength{\oddsidemargin}{-0.5in}
\addtolength{\evensidemargin}{-0.5in}
\addtolength{\textwidth}{1in}
\addtolength{\topmargin}{-.5in}
\addtolength{\textheight}{1.0in}

\urlstyle{same}

\raggedbottom
\raggedright
\setlength{\tabcolsep}{0in}

% Sections formatting
\titleformat{\section}{
  \vspace{-4pt}\scshape\raggedright\large
}{}{0em}{}[\color{black}\titlerule \vspace{-5pt}]

%-------------------------
% Custom commands
\newcommand{\resumeItem}[2]{
  \item\small{
    \textbf{#1}{: #2 \vspace{-2pt}}
  }
}

\newcommand{\companyItem}[1]{
  \item\small{
    {#1 \vspace{-2pt}}
  }
}


\newcommand{\resumeSubheading}[4]{
  \vspace{-1pt}\item
    \begin{tabular*}{0.97\textwidth}[t]{l@{\extracolsep{\fill}}r}
      \textbf{#1} & #2 \\
      \textit{\small#3} & \textit{\small #4} \\
    \end{tabular*}\vspace{-5pt}
}

\newcommand{\resumeSubItem}[2]{\resumeItem{#1}{#2}\vspace{-4pt}}
\renewcommand{\labelitemii}{$\circ$}

\newcommand{\resumeSubHeadingListStart}{\begin{itemize}[leftmargin=*]}
\newcommand{\resumeSubHeadingListEnd}{\end{itemize}}
\newcommand{\resumeItemListStart}{\begin{itemize}}
\newcommand{\resumeItemListEnd}{\end{itemize}\vspace{-5pt}}

%-------------------------------------------
%----------- CV STARTS HERE-----------------


\begin{document}

%----------HEADING-----------------
\begin{tabular*}{\textwidth}{l@{\extracolsep{\fill}}r}
  \textbf{\href{[[.Basics.Website]]}{\Large [[.Basics.Name]]}} & Email : \href{mailto:[[.Basics.Email]]}{[[.Basics.Email]]}\\
  \href{[[.Basics.Website]]}{[[.Basics.Website]]} [[if ne .Basics.Phone ""]]& Phone: [[.Basics.Phone]][[end]] \\
\end{tabular*}


%-----------EDUCATION-----------------
\section{Education}
  \resumeSubHeadingListStart[[ range $key, $value := .Education ]]
    \resumeSubheading
      {[[$value.Institution]]}{[[$value.Location]]}
      {[[if ne $value.StudyType ""]][[$value.StudyType]] in [[end]][[$value.Area]][[if ne $value.GPA ""]]; GPA: [[$value.GPA]][[end]]}{[[$value.StartDate]] - [[if eq $value.EndDate ""]]Present[[else]][[$value.EndDate]][[end]]}
[[end]]
	\resumeSubHeadingListEnd


%-----------EXPERIENCE-----------------
\section{Experience}
  \resumeSubHeadingListStart[[ range $key, $value := .Work ]]
	  \resumeSubheading
	    {[[$value.Company]]}[[if ne $value.Location ""]]{[[$value.Location]]}[[end]]
	    {[[$value.Position]]}{[[$value.StartDate]] - [[if eq $value.EndDate ""]]Present[[else]][[$value.EndDate]][[end]]}
	    \resumeItemListStart[[ range $itemKey, $itemValue := $value.Highlights ]]
	      \companyItem
	        {[[$itemValue]]}[[ end ]]
			\resumeItemListEnd
[[ end ]]
  \resumeSubHeadingListEnd

%-----------PROJECTS-----------------
\section{Projects}
  \resumeSubHeadingListStart[[ range $key, $value := .Projects ]]
    \resumeSubItem{[[$value.Name]]}
      {[[$value.Summary]]}[[end]]
  \resumeSubHeadingListEnd

%
%--------PROGRAMMING SKILLS------------
\section{Skills}
  \resumeSubHeadingListStart[[ range $key, $value := .Skills ]]
    \item{
      \textbf{[[$value.Name]]}{: [[ range $itemKey, $itemValue := .Keywords ]][[if $itemKey]], [[end]][[$itemValue]][[end]]}
      \hfill
    }[[end]]
  \resumeSubHeadingListEnd


%-------------------------------------------
\end{document}
`
