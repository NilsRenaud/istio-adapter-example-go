FROM golang:buster

EXPOSE 9070


WORKDIR src/go/adapterserver

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["adapterserver"]