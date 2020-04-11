FROM golang

RUN mkdir /app && \
    go get github.com/gorilla/mux && \
    go get github.com/dgrijalva/jwt-go && \
    go get golang.org/x/oauth2/google

ADD . /app
WORKDIR /app
RUN go build -o main .
ENTRYPOINT ["/app/main"]