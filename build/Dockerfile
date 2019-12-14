FROM golang:latest

COPY ./entrypoint.sh /entrypoint.sh

RUN go get -u github.com/uma-co82/AtCoder-golang && \
    chmod 777 /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]