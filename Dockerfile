FROM golang:1.14
WORKDIR '/src'
COPY main.go ./
COPY go.mod ./
RUN go build -o /usr/bin/hellokube main.go && \
    rm -rf /src
EXPOSE 8090
CMD ["/usr/bin/hellokube"]