FROM golang:1.17.3-buster as build

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/src/app
COPY src/ go.mod go.sum /go/src/app/
RUN go mod tidy
RUN go build -o /go/bin/app

FROM gcr.io/distroless/base-debian10
COPY --from=build /go/bin/app /
EXPOSE 8080
CMD ["/app"]
