PACKAGES=$(shell go list ./... | grep -v '/vendor/')
MARKDOWNS=$(shell find example/ -name "*md" -type f)

all: get_vendor_deps install test

build:
	go build ./cmd/...

install:
	go install ./cmd/...

test: test_unit test_example

test_unit:
	go test $(PACKAGES)

test_example: get_shunit2
	shelldown ${MARKDOWNS} 
	for script in ${MARKDOWNS} ; do \
		echo "\n\n\nRunning test for script: $$script.sh" ; \
		bash $$script.sh ; \
	done

get_vendor_deps:
	go get -u -v github.com/Masterminds/glide
	glide install

get_shunit2:
	wget "https://raw.githubusercontent.com/kward/shunit2/master/source/2.1/src/shunit2" \
		-q -O example/shunit2

.PHONY: all build install test get_vendor_deps
