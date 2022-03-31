FROM golang:alpine AS base_app

RUN apk update && apk add --no-cache git

WORKDIR /src

COPY . .

RUN go mod download

RUN GOOS=linux go build -ldflags="-w -s" -o app ./app/.

#############################################################################

FROM golang:alpine AS market

WORKDIR /

COPY --from=base_app ./src/app ./

ENTRYPOINT ["/app"]