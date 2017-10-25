### Multi-stage build
FROM golang:1.8.3-alpine3.6 as build

RUN apk --no-cache add git curl openssh

COPY keys/id_rsa /root/.ssh/id_rsa
RUN chmod 700 /root/.ssh/id_rsa && \
    echo -e "Host github.com\n\tStrictHostKeyChecking no\n" >> /root/.ssh/config && \
    git config --global url."ssh://git@github.com:".insteadOf "https://github.com"

RUN go get -u github.com/goadesign/goa/... && \
    go get -u gopkg.in/mgo.v2 && \
    go get -u golang.org/x/crypto/bcrypt && \
    go get -u github.com/JormungandrK/microservice-security/... && \
    go get -u github.com/JormungandrK/microservice-tools

COPY . /go/src/github.com/JormungandrK/user-microservice
RUN go install github.com/JormungandrK/user-microservice


### Main
FROM alpine:3.6

COPY --from=build /go/bin/user-microservice /usr/local/bin/user-microservice
EXPOSE 8080

ENV API_GATEWAY_URL="http://localhost:8001"

CMD ["/usr/local/bin/user-microservice"]
