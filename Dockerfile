FROM golang:alpine

WORKDIR /app

COPY go.mod go.sum ./
COPY *.go ./
COPY apis/* ./apis/
COPY structs/* ./structs/
COPY templates/* ./templates/
COPY gin/* ./gin/
COPY db/* ./db/

RUN CGO_ENABLED=1 go build -o /gc

EXPOSE 8080

CMD ["/gc"]