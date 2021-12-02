FROM golang:alpine AS builder

# Your's truly and only..
LABEL maintainer="Neil Haria <neil.haria@think360.ai>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the working directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

######## Start a new stage from scratch #######
# Cause you want the resultant image size to as small as possible.. ~67 MBs peace!
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Copy the source file `again` from the current directory to the working directory inside the container
# This is mandatory if you have internal packages / modules (not hosted publicly on github)
# If skipped, `go get` command will search those internal modules in public repositories (github)
# and boom, docker image building will fail :/
COPY --from=builder /app .

# (Optional) Just to check all the files present in the container images
# had to run until I found the underlying issue
# RUN ls

# Expose port 8000 to the outside world
EXPOSE 8000

# Command to run the executable
CMD ["./main"]