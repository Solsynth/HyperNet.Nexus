# Building Backend
FROM golang:alpine as nexus-server

WORKDIR /source
COPY . .

WORKDIR /source
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -buildvcs -o /dist ./pkg/main.go

# Runtime
FROM golang:alpine

RUN apk add postgresql-client

COPY --from=nexus-server /dist /nexus/server
COPY ./templates /templates

EXPOSE 8444

CMD ["/nexus/server"]
