FROM ubuntu:latest

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update
RUN apt-get -y install texlive-xetex texlive-latex-extra texlive-fonts-recommended
RUN apt-get -y install git

# Install Google Fonts
RUN git clone https://github.com/google/fonts.git /googlefont
RUN mkdir -p /usr/share/fonts/truetype/google-fonts/
RUN mkdir -p /fonts
RUN mv /googlefont/ofl/* /fonts/

RUN apt-get -q clean

WORKDIR /data
VOLUME ["/data"]
