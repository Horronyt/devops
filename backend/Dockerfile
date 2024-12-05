FROM golang:1.19
	WORKDIR /app
	COPY . .
	RUN go mod init my-go-web-app
	RUN go get github.com/lib/pq
	RUN go build -o api .
	CMD ["./api"]
