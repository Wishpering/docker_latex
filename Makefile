img:
	docker build -t wishpering/latex:debian -f docker/Dockerfile .
deps:
	go build -o scripts/bin/build_deps scripts/build_deps.go
all:
	go build -o scripts/bin/build_deps scripts/build_deps.go
	docker build -t wishpering/latex:debian -f docker/Dockerfile .
clean:
	rm -rvf scripts/bin
