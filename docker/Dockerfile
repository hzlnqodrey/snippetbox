# Use the official Go image as the base image
FROM golang:1.18.5-alpine3.15 AS builder

# Set the working directory inside the container
WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the GO APPLICATION
RUN CGO_ENABLED=0 GOOS=linux go build -a  -installsuffix cgo -o snippetboxapp ./cmd/web

# CGO_ENABLED [This environment variable disables the use of the cgo (C compiler) during the build process. It's often used in statically compiled Go binaries to ensure that the binary is fully self-contained and doesn't rely on external C libraries.]
# GOOS=linux [This environment variable specifies the target operating system for the build. In this case, it's set to linux to build a Linux executable.]
# -a [This flag tells the Go build tool to rebuild all the packages, even if they are already up to date. It ensures that all dependencies are included and up to date.]
# -installsuffix cgo [This flag specifies an install suffix for the package directory. In this case, it's set to cgo, which means that the build process should avoid using cgo, as mentioned earlier.]
# -o webapp [This flag specifies the output binary name (webapp in this case).]

# Stage 2: Build the final image
FROM alpine:latest

