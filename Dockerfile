FROM ubuntu:latest

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update
RUN apt-get -y install fontconfig texlive-xetex texlive-latex-extra texlive-fonts-recommended texlive-fonts-extra fonts-lato
# RUN apt-get -y install git

# Install Google Fonts
# RUN git clone https://github.com/google/fonts.git /googlefont
# RUN mkdir -p /usr/share/fonts/truetype/google-fonts/
# RUN mkdir -p /fonts
# RUN mv /googlefont/ofl/* /fonts/
RUN  apt-get purge -f -y make-doc \
  texlive-fonts-extra-doc \
  texlive-fonts-recommended-doc \
  texlive-humanities-doc \
  texlive-latex-base-doc \
  texlive-latex-extra-doc \
  texlive-latex-recommended-doc \
  texlive-metapost-doc \
  texlive-pictures-doc \
  texlive-pstricks-doc \
  texlive-science-doc &&\
  fc-cache -fv &&\
  apt-get clean &&\
  apt-get autoclean -y &&\
  apt-get autoremove -y &&\
  apt-get clean &&\
  rm -rf /tmp/* /var/tmp/* &&\
  rm -rf /var/lib/apt/lists/* &&\
  rm -f /etc/ssh/ssh_host_*

WORKDIR /data
VOLUME ["/data"]
