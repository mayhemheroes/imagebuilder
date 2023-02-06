FROM golang:1.18 as builder

COPY . /imagebuilder
WORKDIR /imagebuilder/mayhem

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN go get github.com/dvyukov/go-fuzz/go-fuzz-dep
RUN go get github.com/AdaLogics/go-fuzz-headers
RUN apt update && apt install -y clang

ENV GOFLAGS -mod=mod

RUN go-fuzz-build -libfuzzer -o fuzzImagebuilder.a && \
    clang -fsanitize=fuzzer fuzzImagebuilder.a -o fuzzImagebuilder.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /imagebuilder/mayhem/fuzzImagebuilder.libfuzzer /