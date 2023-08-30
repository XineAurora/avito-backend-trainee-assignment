FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./
RUN go mod download
RUN go build -o user-segmentation ./cmd/user-segmentation/user-segmentation.go

CMD ["./user-segmentation"]
