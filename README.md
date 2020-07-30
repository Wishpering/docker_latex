# What is it?

Docker image based on alpine:latest for building LaTeX docs in docker container.\
Includes such packages as make, pygments.

## Installation

Make image by yourself

```bash
git clone https://github.com/Wishpering/docker_latex
docker build .
```

Or just run
```bash
docker pull wishpering/latex:latest
```

## Usage

```bash
docker run --rm -it \
           --name latex_builder \
           -v $(PWD):/project \ 
           $IMAGE_NAME:latest \ 
           sh -c "cd /project && (make/pdflatex/etc)"
```
