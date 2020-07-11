FROM debian:latest

RUN  apt-get update && \
     apt-get install -y make locales texlive-base \
     	     texlive-latex-recommended texlive-fonts-recommended \
	     texlive-lang-cyrillic texlive-latex-extra \
	     python3-pygments cm-super-minimal && \
     localedef -i ru_RU -c -f UTF-8 -A /usr/share/locale/locale.alias ru_RU.UTF-8 && \
     texhash && update-texmf && \
     rm -rf /var/cache/apt/* && rm -rf /var/lib/apt/lists/*

ENV LANG ru_RU.UTF-8

COPY include_pkg /tex/deps
COPY scripts/bin/* /tex/

RUN /tex/build_deps