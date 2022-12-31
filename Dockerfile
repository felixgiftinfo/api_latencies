FROM golang:alpine as builder
LABEL maintainer="Unique <felixgiftinfo@gmail.com>"
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

RUN mkdir /app
WORKDIR /app
COPY files/ /app/
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o /build-main .
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /build-main .

######Multi stage section
FROM alpine:latest as server
WORKDIR /
COPY --from=builder /build-main /build-main
COPY --from=builder /app/files /files
RUN chmod a+x /build-main
# ENV PORT=8000
EXPOSE 8000
CMD [ "/build-main" ]