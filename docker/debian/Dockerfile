FROM debian:latest

RUN  apt-get update && apt-get install -y make locales texlive-latex-recommended texlive-fonts-recommended texlive-lang-cyrillic texlive-latex-extra && \
     localedef -i ru_RU -c -f UTF-8 -A /usr/share/locale/locale.alias ru_RU.UTF-8 && \
     rm -rf /var/cache/apt/* && rm -rf /var/lib/apt/lists/*

ENV LANG ru_RU.utf8

COPY deps /tex/deps
COPY scripts/bin/* /tex/

RUN /tex/build_deps