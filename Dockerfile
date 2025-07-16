# syntax=docker/dockerfile:1
FROM golang:1.24.3-alpine

RUN apk add --no-cache ffmpeg curl python3
RUN curl -L "https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp" \
    -o /usr/local/bin/yt-dlp && \
    chmod a+rx /usr/local/bin/yt-dlp


WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o api .

EXPOSE 8080

CMD ["./api"]
