FROM golang:1.21 AS builder
WORKDIR /src
COPY . .
ARG TARGETOS
ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
  go build -o /bin/scraper .

FROM alpine:3.19.1
COPY --from=builder /bin/scraper /bin/scraper
CMD ["/bin/scraper"]