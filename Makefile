.phony: build-mono
build-mono:
	go build -gcflags '-N -l' -o cmd/monomorphization/monomorphization ./cmd/monomorphization/

.phony: build-bench
build-bench:
	go build -gcflags '-N -l' -o cmd/bench/bench ./cmd/bench/