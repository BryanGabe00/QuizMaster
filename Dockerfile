FROM golang:latest
ENV GO111MODULE=off
ENV GOPATH=/
RUN mkdir /src
RUN mkdir /src/QuizMaster
ADD . /src/QuizMaster
WORKDIR /src/QuizMaster/httpserver
RUN go build -o main .
EXPOSE 9000
CMD ["/src/QuizMaster/httpserver/main"]