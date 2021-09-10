.PHONY: install
install:
	go get -u \
		github.com/dvyukov/go-fuzz/go-fuzz \
		github.com/dvyukov/go-fuzz/go-fuzz-build \
		github.com/dvyukov/go-fuzz/go-fuzz-dep

.PHONY: stdjson
stdjson:
	cd stdjson-fuzz && go-fuzz-build &&  \
	go-fuzz -bin=./json-fuzz.zip -workdir=../workdir

.PHONY: sonic
sonic:
	go-fuzz-build && \
	go-fuzz -bin=./sonic-fuzz.zip -workdir=./workdir
