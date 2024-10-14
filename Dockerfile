FROM --platform=$BUILDPLATFORM golang:1.23-alpine3.20 AS builder
RUN go env -w GOMODCACHE=/root/.cache/go-build
# Define target arch variables so we can use them while crosscompiling, will be set automatically
ARG TARGETOS
ARG TARGETARCH
WORKDIR /go/src/

# get dependencies
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/root/.cache/go-build go mod download

# copy code
COPY . .

# Build project
RUN --mount=type=cache,target=/root/.cache/go-build GOOS=$TARGETOS GOARCH=$TARGETARCH CGO_ENABLED=0 go build -ldflags "-s -w" -a -installsuffix cgo -o /radix-prometheus-proxy

# Final stage, ref https://github.com/GoogleContainerTools/distroless/blob/main/base/README.md for distroless
FROM gcr.io/distroless/static

COPY --from=builder /radix-prometheus-proxy /radix-prometheus-proxy

EXPOSE 8000
USER 1000
ENTRYPOINT ["/radix-prometheus-proxy"]
