FROM golang:alpine
LABEL maintainer="Michael Thanh <mthanh@circleci.com>"

# Working directory
WORKDIR /app

# Copy over dependency files and download deps
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY cmd ./cmd

# Build the executable
RUN go build -o user cmd/user/main.go

# Run executable
CMD ["./user"]

