package templates

var Modern = `
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
% Modern, based on Deedy
%
% Original author:
% Debarghya Das (http://debarghyadas.com)
% Modified and Templated by Kyle Lucas
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%

\documentclass[]{modern}


\begin{document}

%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%
%     LAST UPDATED DATE
%
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%\lastupdated

%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%
%     TITLE NAME
%
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%


\namesection{[[.Basics.Name]]}{}{%
  \urlstyle{same}\href{[[.Basics.Website]]}{[[.Basics.Website]]}
| \href{mailto:[[.Basics.Email]]}{[[.Basics.Email]]}
[[if ne .Basics.Name ""]]| [[.Basics.Phone]]%[[end]]
}

%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%
%     COLUMN ONE
%
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%

\begin{minipage}[t]{0.33\textwidth}

%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%     EDUCATION
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%

\section{Education}
[[ range $key, $value := .Education ]]
\subsection{[[$value.Institution]]}
\descript{[[$value.Area]][[if ne $value.StudyType ""]] | [[$value.StudyType]][[end]]}
\location{[[$value.StartDate]] - [[if eq $value.EndDate ""]]Present[[else]][[$value.EndDate]][[end]][[if ne $value.GPA ""]] | [[$value.GPA]][[end]]}[[end]]
\sectionsep

%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%     LINKS
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%

\section{Links}[[ range $key, $value := .Basics.Profiles ]]
[[$value.Network]]: \href{[[$value.URL]]}{\custombold{[[$value.Username]]}} \\[[end]]
\sectionsep

%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%     SKILLS
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%

\section{Skills}[[ range $key, $value := .Skills ]]
\subsection{[[$value.Name]]}
\cvskill{[[ range $itemKey, $itemValue := .Keywords ]][[if $itemKey]], [[end]][[$itemValue]][[end]]}{4}
[[end]]

\sectionsep

%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%     Projects
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%

\section{Projects}[[ range $key, $value := .Projects ]]
\location{[[$value.Name]] | [[$value.ReleaseDate]]}
\vspace{\topsep} % Hacky fix for awkward extra vertical space
\begin{tightemize}
\item [[$value.Summary]]
\end{tightemize}[[end]]

\sectionsep

%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%
%     COLUMN TWO
%
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%

\end{minipage}
\hfill
\begin{minipage}[t]{0.66\textwidth}

%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%     EXPERIENCE
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%

\section{Experience}

[[ range $key, $value := .Work ]]
\runsubsection{[[$value.Company]]}
\descript{| [[$value.Position]]}
\location{[[$value.StartDate]] - [[if eq $value.EndDate ""]]Present[[else]][[$value.EndDate]][[end]] | [[$value.Location]]}
\vspace{\topsep} % Hacky fix for awkward extra vertical space
\begin{tightemize}[[ range $itemKey, $itemValue := $value.Highlights ]]
\item [[$itemValue]][[end]]
\end{tightemize}
\sectionsep
[[end]]

%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
%     AWARDS
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%

% NOTE: These are not mine. They are remnants from the template.
%  \section{Awards}
%  \begin{tabular}{rll}
%  2014	     & top 52/2500  & KPCB Engineering Fellow\\
%  2014	     & 2\textsuperscript{nd} most points  & Google Code Jam, Qualification Round\\
%  2014	     & 1\textsuperscript{st}/50  & Microsoft Coding Competition, Cornell\\
%  2013	     & National  & Jump Trading Challenge Finalist\\
%  2013     & 7\textsuperscript{th}/120 & CS 3410 Cache Race Bot Tournament  \\
%  2012     & 2\textsuperscript{nd}/150 & CS 3110 Biannual Intra-Class Bot Tournament \\
%  2011     & National & Indian National Mathematics Olympiad (INMO) Finalist \\
%  2010     & National & Comp. Soc. of India's National Programming Contest\\
%  \end{tabular}
%  \sectionsep
%
\end{minipage}
\end{document} \documentclass[]{article}
`

var ModernClass = `
% Intro Options
\ProvidesClass{modern}[2014/04/30 CV class]
\NeedsTeXFormat{LaTeX2e}
\DeclareOption{print}{\def\@cv@print{}}
\DeclareOption*{%
  \PassOptionsToClass{\CurrentOption}{article}
}
\ProcessOptions\relax
\LoadClass{extarticle}

% Package Imports
\usepackage[hmargin=1.25cm, vmargin=0.7cm]{geometry}
\usepackage[usenames,dvipsnames]{xcolor}
\usepackage{hyperref}
\usepackage{titlesec}
\usepackage[absolute]{textpos}
\usepackage[english]{babel}
\usepackage[english]{isodate}
\usepackage{fontspec,xltxtra,xunicode}

\defaultfontfeatures{
  Path = /usr/share/fonts/truetype/font-awesome/ }

\usepackage{fontawesome}
% Color definitions
\definecolor{date}{HTML}{666666}
\definecolor{primary}{HTML}{2b2b2b}
\definecolor{headings}{HTML}{6A6A6A}
\definecolor{subheadings}{HTML}{333333}

% Set main fonts
\defaultfontfeatures{Mapping=tex-text}
\setmainfont[Color=primary, Path = /usr/share/fonts/truetype/lato/]{Lato-Light}
\setsansfont[Scale=MatchLowercase,Mapping=tex-text, Path = /fonts/raleway/]{Raleway-ExtraLight}
\newcommand{\custombold}[1]{\color{subheadings}\fontspec[Path = /usr/share/fonts/truetype/lato/]{Lato-Regular}\selectfont #1 \normalfont}

% Date command
\setlength{\TPHorizModule}{1mm}
\setlength{\TPVertModule}{1mm}
\textblockorigin{0mm}{5mm} % start everyth
\newcommand{\lastupdated}{\begin{textblock}{60}(165,0)
\color{date}\fontspec[Path = /fonts/raleway/]{Raleway-ExtraLight}\fontsize{8pt}{10pt}\selectfont
Last Updated on
\today
\end{textblock}}

% Name command
\newcommand{\namesection}[3]{
	\centering{
		\sffamily
		\fontspec[Path = /usr/share/fonts/truetype/lato/]{Lato-Light}\fontsize{40pt}{10cm}\selectfont #1
		\fontspec[Path = /usr/share/fonts/truetype/lato/]{Lato-Light}\selectfont #2
	} \\
	\vspace{5pt}
	\centering{ \color{headings}\fontspec[Path = /fonts/raleway/]{Raleway-Medium}\fontsize{11pt}{14pt}\selectfont #3}
	\noindent\makebox[\linewidth]{\rule{\paperwidth}{0.4pt}}
	\vspace{-15pt}
}
\titlespacing{\section}{0pt}{0pt}{0pt}

% Headings command
\titleformat{\section}{\color{headings}
\scshape\fontspec[Path = /usr/share/fonts/truetype/lato/]{Lato-Light}\fontsize{16pt}{24pt}\selectfont \raggedright\uppercase}{} {0em}{}

% Subeadings command
\titleformat{\subsection}{\color{subheadings}
\fontspec{Lato-Bold}\fontsize{12pt}{12pt}\selectfont\bfseries\uppercase}{}{0em}{}
\titlespacing{\subsection}{0pt}{\parskip}{-\parskip}
\titlespacing{\subsubsection}{0pt}{\parskip}{-\parskip}
\newcommand{\runsubsection}[1]{\color{subheadings}
\fontspec[Path = /usr/share/fonts/truetype/lato/]{Lato-Bold}\fontsize{12pt}{12pt}\selectfont\bfseries\uppercase {#1} \normalfont}

% Descriptors command
\newcommand{\descript}[1]{\color{subheadings}\raggedright\scshape\fontspec[Path = /fonts/raleway/]{Raleway-Medium}\fontsize{11pt}{13pt}\selectfont {#1 \\} \normalfont}

% Location command
\newcommand{\location}[1]{\color{headings}\raggedright\fontspec[Path = /fonts/raleway/]{Raleway-Medium}\fontsize{10pt}{12pt}\selectfont {#1\\} \normalfont}

% Section seperators command
\newcommand{\sectionsep}[0]{\vspace{8pt}}

% Bullet Lists with fewer gaps command
\newenvironment{tightemize}{\vspace{-\topsep}\begin{itemize}\itemsep1pt \parskip0pt \parsep0pt}{\end{itemize}\vspace{-\topsep}}

\colorlet{emphasis}{black}
\colorlet{accent}{black}
\newcommand{\ratingmarker}{\faCircle}

\newcommand{\cvskill}[2]{%
  \textcolor{emphasis}{\textbf{#1}}\hfill
  \foreach \x in {1,...,5}{%
    \space{\ifnumgreater{\x}{#2}{\color{body!30}}{\color{accent}}\ratingmarker}}\par%
  }
`
