FROM golang

RUN mkdir /app && go get github.com/gorilla/mux

ADD . /app
WORKDIR /app
RUN go build -o main .
ENTRYPOINT ["/app/main"]