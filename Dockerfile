# builder
FROM golang:1.19 as builder
WORKDIR /builder

ENV PORT=3000

COPY go.mod go.sum /builder/
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/app ./cmd

#runner
FROM gcr.io/distroless/static-debian11

COPY --from=builder /go/bin/app /
CMD ["/app"]

EXPOSE $PORT