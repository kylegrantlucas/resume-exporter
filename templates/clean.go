package templates

var Clean = `
%%%%%%%%%%%%%%%%%
% This is an example CV created using altacv.cls (v1.1.5, 1 December 2018) written by
% LianTze Lim (liantze@gmail.com), based on the
% Cv created by BusinessInsider at http://www.businessinsider.my/a-sample-resume-for-marissa-mayer-2016-7/?r=US&IR=T
%
%% It may be distributed and/or modified under the
%% conditions of the LaTeX Project Public License, either version 1.3
%% of this license or (at your option) any later version.
%% The latest version of this license is in
%%    http://www.latex-project.org/lppl.txt
%% and version 1.3 or later is part of all distributions of LaTeX
%% version 2003/12/01 or later.
%%%%%%%%%%%%%%%%

%% If you are using \orcid or academicons
%% icons, make sure you have the academicons
%% option here, and compile with XeLaTeX
%% or LuaLaTeX.
% \documentclass[10pt,a4paper,academicons]{altacv}

%% Use the "normalphoto" option if you want a normal photo instead of cropped to a circle
% \documentclass[10pt,a4paper,normalphoto]{altacv}

\documentclass[10pt,a4paper,ragged2e]{altacv}

%% AltaCV uses the fontawesome and academicon fonts
%% and packages.
%% See texdoc.net/pkg/fontawecome and http://texdoc.net/pkg/academicons for full list of symbols. You MUST compile with XeLaTeX or LuaLaTeX if you want to use academicons.

% Change the page layout if you need to
\geometry{left=2cm,right=10cm,marginparwidth=6.8cm,marginparsep=1.2cm,top=1.25cm,bottom=1.25cm}

% Change the font if you want to, depending on whether
% you're using pdflatex or xelatex/lualatex
\ifxetexorluatex
  % If using xelatex or lualatex:
  \setmainfont{Carlito}
\else
  % If using pdflatex:
  \usepackage[utf8]{inputenc}
  \usepackage[T1]{fontenc}
  \usepackage[default]{lato}
\fi

% Change the colours if you want to
\definecolor{VividPurple}{HTML}{000000}
\definecolor{SlateGrey}{HTML}{2E2E2E}
\definecolor{LightGrey}{HTML}{2E2E2E}
\colorlet{heading}{VividPurple}
\colorlet{accent}{VividPurple}
\colorlet{emphasis}{SlateGrey}
\colorlet{body}{LightGrey}

% Change the bullets for itemize and rating marker
% for \cvskill if you want to
\renewcommand{\itemmarker}{{\small\textbullet}}
\renewcommand{\ratingmarker}{\faCircle}

%% sample.bib contains your publications
\addbibresource{sample.bib}

\begin{document}
\name{Ronak Dedhiya}
\tagline{Machine Learning Engineer}
% Cropped to square from https://en.wikipedia.org/wiki/Marissa_Mayer#/media/File:Marissa_Mayer_May_2014_(cropped).jpg, CC-BY 2.0
%\photo{3.3cm}{profile.jpg}
\personalinfo{%
  % Not all of these are required!
  % You can add your own with \printinfo{symbol}{detail}
  \email{dedhiyaronak@gmail.com}
%   \phone{000-00-0000}
%  \mailaddress{Address, Street, 00000 County}
  \location{Mumbai, India}
%  \homepage{marissamayr.tumblr.com/}
%  \twitter{@marissamayer}
  \linkedin{linkedin.com/in/ronak-dedhiya}
   \github{github.com/ronakdedhiya} % I'm just making this up though.
%   \orcid{orcid.org/0000-0000-0000-0000} % Obviously making this up too. If you want to use this field (and also other academicons symbols), add "academicons" option to \documentclass{altacv}
}

%% Make the header extend all the way to the right, if you want.
\begin{fullwidth}
\makecvheader
\end{fullwidth}

%% Depending on your tastes, you may want to make fonts of itemize environments slightly smaller
\AtBeginEnvironment{itemize}{\small}

%% Provide the file name containing the sidebar contents as an optional parameter to \cvsection.
%% You can always just use \marginpar{...} if you do
%% not need to align the top of the contents to any
%% \cvsection title in the "main" bar.
\cvsection[page1sidebar]{Experience}

\cvevent{Machine Learning Engineer}{AitoeLabs ( AiSight Video Analytics Pvt. Ltd. )}{June 2018 -- Present}{Mumbai,India}
\begin{itemize}
\item Research and development of Computer Vision algorithms centered around the video surveillance domain in collaboration with NCETIS, IIT Bombay.
\smallskip
\item Working with team on design, development and integration of real-time video analytic solution.
\smallskip
\item Leading the team on the training and evaluation of Deep Neural Network models.
\smallskip
\item Experience with performance analysis, optimizations and benchmark evaluations.
\end{itemize}

\divider

\cvevent{Software Engineer}{Atos India Pvt. Ltd.}{Feb 2017 -- June 2018}{Pune,India}
\begin{itemize}
\item Developed multiple Proofs-of-Concepts in order to build Machine Learning capabilities in the team.
\smallskip
\item Incorporated various Machine Learning features in Atos's internal tools, in order to utilize the performance gain offered by Machine Learning algorithms.
\item Deploying machine learning solution on cloud.
\end{itemize}

%\divider

\cvsection{Achievements}
\smallskip
\begin{itemize}
\item Placed in top 14\% in RSNA Pneumonia Detection Kaggle challenge
\smallskip
\item Placed in top 3\% in Digit Recognition Challenge using One-Shot Learning
\smallskip
\item Winner at Samsung SMS Classification Hackathon
\smallskip
\item Placed in top 5\% in "Predict the Happiness" Hackerearth Challenge
\smallskip
\item Secured AIR 1232 in GATE 2017
\end{itemize}

\cvsection{SKILLS}

\cvskill{C++, Python, Keras, Unix}{5}
%\divider

\cvskill{Tensorflow, Pytorch, Darknet}{4}
% \divider
%\cvskill{German}{3}


\cvsection{Education / Courses}
\cvevent{Deep Learning Specialization}{Coursera}{ June 2017 -- Aug 2017}{}
%\divider
\cvevent{Bachelor of Technology}{Vivekanand Education Society Institute of Technology}{ June 2012 -- May 2016}{}
% \divider

% \cvevent{Product Engineer}{Google}{23 June 1999 -- 2001}{Palo Alto, CA}

% \begin{itemize}
% \item Joined the company as employe \#20 and female employee \#1
% \item Developed targeted advertisement in order to use user's search queries and show them related ads
% \end{itemize}

%\cvsection{A Day of My Life}

% Adapted from @Jake's answer from http://tex.stackexchange.com/a/82729/226
% \wheelchart{outer radius}{inner radius}{
% comma-separated list of value/text width/color/detail}
% Some ad-hoc tweaking to adjust the labels so that they don't overlap
% \wheelchart{1.5cm}{0.5cm}{%
%   10/10em/accent!30/Sleeping \& dreaming about work,
%   25/9em/accent!60/Public resolving issues with Yahoo!\ investors,
%   5/13em/accent!10/\footnotesize\\[1ex]New York \& San Francisco Ballet Jawbone board member,
%   20/15em/accent!40/Spending time with family,
%   5/8em/accent!20/\footnotesize Business development for Yahoo!\ after the Verizon acquisition,
%   30/9em/accent/Showing Yahoo!\ employees that their work has meaning,
%   5/8em/accent!20/Baking cupcakes
% }

\clearpage

% \cvsection[page2sidebar]{Publications}

\nocite{*}

% \printbibliography[heading=pubtype,title={\printinfo{\faBook}{Books}},type=book]

% \divider

% \printbibliography[heading=pubtype,title={\printinfo{\faFileTextO}{Journal Articles}}, type=article]

% \divider

% \printbibliography[heading=pubtype,title={\printinfo{\faGroup}{Conference Proceedings}},type=inproceedings]

% %% If the NEXT page doesn't start with a \cvsection but you'd
% %% still like to add a sidebar, then use this command on THIS
% %% page to add it. The optional argument lets you pull up the
% %% sidebar a bit so that it looks aligned with the top of the
% %% main column.
% % \addnextpagesidebar[-1ex]{page3sidebar}


\end{document}`
