FROM golang:1.18-alpine

WORKDIR /redis_docker

# Copy everything from this project into the filesystem of the container.
COPY . .

# Obtain the package needed
RUN go get
RUN go mod tidy

# Compile the binary exe for our app.
RUN go build -o main .
# Start the application.
CMD ["./main"]