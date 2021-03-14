INFORMATIONAL
=============
Authors: Rich Goluszka, Juan Moncada, Bryan Gabe  
Last Updated: 3/14/2021  
Project: QuizMaster  
Course: Applied Programming Languages (CPSC360-1) with Professor Eric Pogue

Contact
-------
Please direct any and all comments/concerns/inquiries to richardjgoluszka@lewisu.edu

ORIGINALITY
===========
Credit to https://youtu.be/lIbdPrUpGz4?t=300 for their example of writing a Dockerfile for Golang.

Credit to https://golangdocs.com/golang-read-json-file for their example of parsing JSON data 
from a .json file in Golang.

Credit to tour.golang.org/methods/17 for their example of implementing the Stringer interface, 
allowing fmt.Println to operate on structs.

BUILD / EXECUTE / DEPENDENCY
============================
Required Files
--------------
QuizMaster  
	-httpclient directory  
		--httpClient.go  
	-httpserver directory  
		--httpServer.go  
	-quizjson directory  
		--bankJSON.go  
		--entryJSON.go  
		--questionJSON.go  
		--reqJSON.go  
	-banks directory  
		--at least one .json file in the same format as bank-1.json  
	-Dockerfile

_Note: The GitHub repository https://github.com/BryanGabe00/QuizMaster contains all required_
_files plus README.md and LICENSE files._

Build & Execution Instructions
------------------------------
_IMPORTANT: The HTTP Client program uses a deployed Azure container as the HTTP Server by default._
_If you desire to run the HTTP Server locally, you WILL have to change lines 18 and 40 in_
_httpClient.go BEFORE building the executable._

Build HTTP Client and Server (Executable):
1. Open a command-line or terminal
2. Navigate to .../go/src/QuizMaster
3. Run `go build` within the httpserver and httpclient subdirectories (.../QuizMaster/httpserver 
and .../QuizMaster/httpclient)
You should now have `httpserver.exe` and `httpclient.exe` executables to run the web server 
and web client.

Build HTTP Server (Docker):
1. Open a command-line or terminal
2. Navigate to .../go/src/QuizMaster
3. Run `docker build -t quizmaster-server .` to build a container image for the HTTP web server.
You should now have a docker container image to run the web server.

Execute HTTP Client:
1. Build the executable _(using above instructions)_
2. Run the `httpclient.exe` executable

Execute HTTP Server (Executable):
1. Build the executable _(using above instructions)_
2. Run the `httpserver.exe` executable

Execute HTTP Server (Docker):
1. Build the container image _(using above instructions)_
2. Open a terminal and run `docker run -dp 80:80 quizmaster-server` to launch a container using
the web server image

Project Information
===================
HTTP Server
-----------
* The HTTP Server serves data on port 80
* A Docker container running the server has been deployed and is used by default
* For custom HTTP Servers, run in its own terminal window or as a background process
* POST requests for question sets are served on localhost:9000/req
* POST requests for question banks are served on localhost:9000

HTTP Client
-----------
* The HTTP Client performs a POST request to the HTTP Server (default server is Azure container)
* Run this in your code editor or its own terminal window
* Requests and displays a question set from the HTTP server

JSON Interactions
-----------------
* All json-related .go files are contained in the quizjson directory (a go package)
* bankJSON.go represents information about the question banks loaded from .json files
* entryJSON.go represents quiz question entries (individually and as a set)
* questionJSON.go represents quiz questions served to clients (individually and as a set)
* reqJSON.go represents HTTP client requests to the HTTP server
* Each .json file should be formatted to match the format of bank-1.json in order to ensure the
file is read by the server correctly