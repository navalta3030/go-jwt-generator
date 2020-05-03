FROM golang

RUN mkdir /app && \
    go get github.com/gorilla/mux && \
    go get github.com/dgrijalva/jwt-go

ADD . /app
WORKDIR /app
RUN go build -o main .
ENTRYPOINT ["/app/main"]
