FROM golang:1.21 AS builder
WORKDIR /src
COPY . .
ENV CGO_ENABLED=0
RUN go build -o /bin/scraper .

FROM scratch
COPY --from=builder /bin/scraper /bin/scraper
CMD ["/bin/scraper"]