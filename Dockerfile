FROM golang AS compile-image

WORKDIR /build_dir
COPY ./scripts/build_deps.go .

RUN go build -o build_deps build_deps.go

FROM alpine:latest
COPY --from=compile-image /build_dir/build_deps /tex/

ENV LANG ru_RU.UTF-8

RUN apk add --no-cache libc6-compat make \
    	    	       texlive texmf-dist-latexextra \
		       texlive-luatex texlive-xetex \
		       texmf-dist-langcyrillic \
		       py3-pygments

WORKDIR /tex
COPY include_pkg ./deps/

RUN ./build_deps