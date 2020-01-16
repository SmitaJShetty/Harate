FROM golang:latest
LABEL maintainer="smita.narasimha@gmail.com"

ENV APP_LISTENER_ADDR=:8080
ENV REQ_COUNT=600
ENV TIME_LIMIT_IN_MINUTES=1

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN cd cmd && go build -o ../build/vayu
EXPOSE 8080

CMD ["./build/vayu"]