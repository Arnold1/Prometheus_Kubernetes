FROM golang:1.9.2

RUN apt-get update && apt-get install -y git

RUN mkdir -p /home/go-mux-example
WORKDIR /home/go-mux-example

COPY main.go /home/go-mux-example
COPY handlers.go /home/go-mux-example

RUN go get -v github.com/gorilla/handlers
RUN go get -v github.com/gorilla/mux
RUN go get -v github.com/prometheus/client_golang/prometheus
RUN go get -v github.com/sirupsen/logrus

RUN chmod 0755 /home/go-mux-example
RUN CGO_ENABLED=0 GOOS=linux go build -o mux

EXPOSE 8080
CMD ["/home/go-mux-example/mux"]