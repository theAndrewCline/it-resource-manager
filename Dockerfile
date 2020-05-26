FROM golang:1.14

WORKDIR /go/src/app
COPY . .

RUN apt-get update && apt-get install -y curl

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["it-resource-manager"]