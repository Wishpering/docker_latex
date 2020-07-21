FROM alpine:latest

ENV LANG ru_RU.UTF-8

RUN apk add --no-cache libc6-compat make \
    	    	       texlive texmf-dist-latexextra \
		       texlive-luatex texlive-xetex \
		       texmf-dist-langcyrillic \
		       py3-pygments

WORKDIR /tex
COPY include_pkg ./deps/
COPY scripts/bin/* ./

RUN ./build_deps