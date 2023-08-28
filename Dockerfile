FROM golang:latest

COPY ./ ./
RUN go mod download
RUN go build ./cmd/user-segmentation/user-segmentation.go

CMD ["./user-segmentation"]
