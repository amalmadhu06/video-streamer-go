#use a light weight base image for Go
FROM golang:1.20-alpine

#install FFmpeg dependencies
RUN apk add --no-cache ffmpeg

#setup working directory
WORKDIR /app

#copy go.mod and go.sum
COPY go.mod go.sum ./

#download go dependencies
RUN go mod download

#copy rest of the code
COPY . .

#build the go application
RUN go build -o ./cmd/main ./cmd

#set the entrypoint for the application
CMD ["./cmd/main"]