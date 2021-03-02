FROM alpine:latest

ENV LANG ru_RU.UTF-8

RUN apk add --no-cache make \
    	    	       texlive texmf-dist-latexextra \
		       texlive-luatex texlive-xetex \
		       texmf-dist-langcyrillic \
		       py3-pygments

COPY include_pkg/ ./tex

WORKDIR /tex
RUN ./build_deps && rm -rfv /tex

RUN adduser -D -s /bin/bash builder

USER builder
WORKDIR /home/builder