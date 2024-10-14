FROM --platform=$BUILDPLATFORM golang:1.23-alpine3.20 AS builder
# Define target arch variables so we can use them while crosscompiling, will be set automatically
ARG TARGETARCH
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=${TARGETARCH}

WORKDIR /go/src/

# get dependencies
COPY go.mod go.sum ./
RUN go mod download

# copy code
COPY . .

# Build project
RUN go build -ldflags "-s -w" -a -installsuffix cgo -o /radix-prometheus-proxy

# Final stage, ref https://github.com/GoogleContainerTools/distroless/blob/main/base/README.md for distroless
FROM gcr.io/distroless/static

COPY --from=builder /radix-prometheus-proxy /radix-prometheus-proxy

EXPOSE 8000
USER 1000
ENTRYPOINT ["/radix-prometheus-proxy"]
