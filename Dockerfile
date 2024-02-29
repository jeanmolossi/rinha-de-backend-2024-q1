FROM golang:1.22-alpine as BUILDER

RUN apk update && apk upgrade && \
    apk add --no-cache git

WORKDIR /src/app

COPY go.sum ./
COPY go.mod ./

RUN go mod download

COPY ./cmd ./cmd
COPY ./internal ./internal

ENV GOOS=linux
ENV GOARCH=amd64

RUN go build -ldflags="-w -s" -o rinha ./cmd/http

FROM scratch as RUNNER

COPY --from=BUILDER /src/app/rinha /bin/rinha

EXPOSE 8080

CMD ["/bin/rinha"]