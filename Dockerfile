FROM golang:1.15.5

LABEL name="media server"

WORKDIR /go/src/mediamanager
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o ./out/app .
EXPOSE 8080
CMD ["./out/app"]
