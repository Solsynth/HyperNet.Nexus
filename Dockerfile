# Building Backend
FROM golang:alpine as nexus-server

WORKDIR /source
COPY . .

WORKDIR /source
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -buildvcs -o /dist ./pkg/main.go

# Runtime
FROM golang:alpine

COPY --from=nexus-server /dist /nexus/server

EXPOSE 8444

CMD ["/nexus/server"]
