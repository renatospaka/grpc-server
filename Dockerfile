# FROM golang:1.16.5-stretch
FROM golang:1.18.0-stretch

## UPDATE THE OS
RUN apt-get update && \
    apt-get install build-essential librdkafka-dev -y  && \
    apt-get install -y tzdata  && \
    go install github.com/spf13/cobra-cli@latest


WORKDIR /go/src

## SET ENVIRONMENT
# ENV PATH="/go/bin:${PATH}"
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
ENV TZ America/Sao_Paulo

## START A PROJECT
RUN go mod init server

## COPY NECESSARY FILES
COPY go.mod .
RUN go mod download && \
    go mod tidy

## KEEP THE CONTAINER RUNNiNG
CMD ["tail", "-f", "/dev/null"]