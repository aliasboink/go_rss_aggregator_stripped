FROM golang:1.21 AS builder
WORKDIR /src
COPY . .
ARG TARGETOS
ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
  go build -o /bin/scraper .

FROM scratch
COPY --from=builder /bin/scraper /bin/scraper
CMD ["/bin/scraper"]