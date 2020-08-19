FROM alpine:latest

ENV LANG ru_RU.UTF-8

RUN apk add --no-cache make \
    	    	       texlive texmf-dist-latexextra \
		       texlive-luatex texlive-xetex \
		       texmf-dist-langcyrillic \
		       py3-pygments

WORKDIR /tex
COPY ./scripts/build_deps .
COPY include_pkg .

RUN ./build_deps