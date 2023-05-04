FROM golang

WORKDIR /app

COPY go.mod go.sum ./
COPY *.go ./
COPY apis/* ./apis/
COPY structs/* ./structs/
COPY templates/* ./templates/
COPY gin/* ./gin/
COPY db/* ./db/
COPY *.yml ./
COPY dbcontext/* ./dbcontext/

RUN go build -o /gc

EXPOSE 8080

CMD ["/gc"]