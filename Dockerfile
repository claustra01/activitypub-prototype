FROM golang:1.22.1-alpine

WORKDIR /app

COPY . .
RUN go mod download
RUN go build -o main .

EXPOSE 8443
ENTRYPOINT [ "./main" ]
