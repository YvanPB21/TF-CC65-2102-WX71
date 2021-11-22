FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /docker-api-tf

EXPOSE 8071 8090 6943 6944 8072

CMD [ "/docker-api-tf" ]