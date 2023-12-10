# Dockerfile definition for Backend application service.

# From which image we want to build. This is basically our environment.
FROM golang:1.21.5-alpine as Build

WORKDIR /src/app

# This will copy all the files in our repo to the inside the container at root location.
COPY . .

# Build our binary at root location.
RUN go build -o main cmd/main.go

####################################################################
# This is the actual image that we will be using in production.
FROM alpine:latest

# We need to copy the binary & config from the build image to the production image.
COPY --from=Build /src/app/main /src/app/cmd/main
COPY --from=Build /src/app/configs /src/app/configs

WORKDIR /src/app/cmd

# This is the port that our application will be listening on.
EXPOSE 8082

# This is the command that will be executed when the container is started.
ENTRYPOINT ["./main"]