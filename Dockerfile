# Go basic image
FROM golang:1.18-alpine

# workdir
WORKDIR /app

# copy go.mod and go.sum
COPY go.mod go.sum ./

# download dependency
RUN go mod download

# copy source code
COPY . .

# build go application
RUN go build -o main ./cmd/todo-list

# expose port
EXPOSE 8080

# run application
CMD ["./main"]
