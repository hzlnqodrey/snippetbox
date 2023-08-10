# Stage 1: Use the official Go image as the base image
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

# Set working directory
WORKDIR /app

# Copy the Go application binary from the builder stage
COPY --from=builder /app/webapp .

# Copy UI files
COPY ui/ ./ui/

# Copy MySQL configuration from pkg directory
COPY pkg/ ./pkg/

# Copy the initialization script
COPY pkg/init.sql /docker-entrypoint-initdb.d/

# Set environment variables for MySQL (using build arguments) [look at deploy.sh docker build cmd]
ARG MYSQL_ROOT_PASSWORD
ARG MYSQL_DATABASE
ARG MYSQL_USER
ARG MYSQL_PASSWORD

ENV MYSQL_ROOT_PASSWORD=$MYSQL_ROOT_PASSWORD
ENV MYSQL_DATABASE=$MYSQL_DATABASE
ENV MYSQL_USER=$MYSQL_USER
ENV MYSQL_PASSWORD=$MYSQL_PASSWORD

# Install MySQL client
RUN apk add --no-cache mysql-client

# Expose the application port
EXPOSE 4000

# Start the application
CMD ["./webapp"]