FROM golang:1.19.0

WORKDIR /usr/src/app

COPY . .
RUN go mod tidy
RUN go build cmd/app/main.go

RUN ls -alp .

ENTRYPOINT [ "./main" ]