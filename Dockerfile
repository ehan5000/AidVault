FROM golang:1.21
WORKDIR /app
COPY . .
RUN go mod init aidvault && go get -d ./...
RUN go build -o aidvault
CMD ["./aidvault"]
