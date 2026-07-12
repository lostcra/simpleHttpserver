# builder is first large image for downloading all dependencies and building services into executable file
FROM golang:1.26.5-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./
RUN go build -o app cmd/httpserver

# lightweight container for running executable file from builder
FROM alpine
WORKDIR /simpleHttpserver/
COPY --from=builder /app/app .
CMD [ "./app" ]
