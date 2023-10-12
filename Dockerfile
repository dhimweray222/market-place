FROM golang:1.19
WORKDIR /app
COPY go.mod go.sum ./
COPY . .
RUN go mod download
RUN go build -o app app.go

CMD ["./app"]


