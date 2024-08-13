FROM golang:1.21.13

RUN mkdir -p /home/app
WORKDIR /home/app

# Copying all the files
COPY . .

RUN go mod download

# Starting our application
CMD ["go", "run", "main.go"]

# Exposing server port
EXPOSE 8080