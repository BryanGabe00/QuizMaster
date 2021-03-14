#FROM IMAGE:TAG sets the base image of the container, golang:latest is a Linux image with a
#	Go environment pre-installed.
FROM golang:latest

#ENV VARIABLE=VALUE sets environment variables in the container, GO111MODULE=off tells the
#	container to use packages and not modules to fulfill dependencies.
ENV GO111MODULE=off

#WORKDIR /DIR/PATH/ sets the current working directory in the container and affects how subsequent
#	instructions are performed (ADD, COPY, CMD, ENTRYPOINT, RUN).
WORKDIR /go/src

#RUN COMMAND runs a command within the container, this line makes a new directory under /go/src
#	called QuizMaster (this will host our Go project).
RUN mkdir /QuizMaster

#ADD /SRC/PATH /DST/PATH copies files from the host computer to the container directory, this line
#	copies our current directory (the Go project) into /go/src/QuizMaster within the container.
ADD . QuizMaster

#Changes our workdir to /go/src/QuizMaster/httpserver, builds the project to an executable, and
#	exposes port 80 of the container to serve incoming requests.
WORKDIR /go/src/QuizMaster/httpserver
RUN go build -o main .
EXPOSE 80

#Starts the compiled executable within the container environment.
CMD ["/go/src/QuizMaster/httpserver/main"]