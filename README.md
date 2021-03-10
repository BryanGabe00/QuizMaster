INFORMATIONAL
=============
Authors: Rich Goluszka, Juan Moncada, Bryan Gabe  
Last Updated: 3/8/2021  
Project: QuizMaster  
Course: Applied Programming Languages (CPSC360-1) with Professor Eric Pogue

Contact
-------
Please direct any and all comments/concerns/inquiries to richardjgoluszka@lewisu.edu

ORIGINALITY
===========
Credit to https://golangdocs.com/golang-read-json-file for an example of parsing JSON data from a 
.json file in Golang.

Credit to tour.golang.org/methods/17 for their example of implementing the Stringer interface, 
allowing fmt.Println to operate on structs.

BUILD / EXECUTE / DEPENDENCY
============================
Required Files
--------------
concept-test  
	-httpclient directory  
		-httpClient.go  
	-httpserver directory  
		-httpServer.go  
	-quizjson directory  
		-bankJSON.go  
		-entryJSON.go  
		-questionJSON.go  
		-reqJSON.go  
	-banks directory  
		-at least one .json file (bank-1.json / bank-2.json / bank-3/json)

_Note: The GitHub repository https://github.com/JuanMoncada/360-RedDragons contains all required_
_files plus README.md and LICENSE files._

Build Instructions
------------------
To compile an executable:
1. Open the command-line or terminal
2. Navigate to .../go/src/360-RedDragons/sprint-4/concept-test
3. Run go install within each subdirectory (httpserver, httpclient, and quizjson)
4. Run go build within the httpserver and httpclient subdirectories (.../concept-test/httpserver 
and .../concept-test/httpclient)
You should now have a `httpserver.exe` executable to run the web server and a `httpclient.exe` 
executable to run the web client.

Execution Instructions
----------------------
1. Build the program _(using above instructions)_
2. Run `httpserver.exe` in a terminal / command-line
3. Run `httpclient.exe` in a separate terminal / command-line
4. Optionally navigate to `localhost:9000` in your browser to view info about the loaded JSON 
question banks

Project Information
===================
HTTP Server
-----------
* The HTTP Server serves data on localhost:9090
* Run this in its own terminal window or as a background process
* POST requests for question sets are served on localhost:9000/req
* POST requests for question banks are served on localhost:9000

HTTP Client
-----------
* The HTTP Client performs a GET request to the HTTP Server
* Run this in your code editor or its own terminal window
* Requests and displays a question set the HTTP server

JSON Interactions
-----------------
* All json-related .go files are contained in the quizjson directory (a go package)
* bankJSON.go represents information about the question banks loaded from .json files
* entryJSON.go represents quiz question entries (individually and as a set)
* questionJSON.go represents quiz questions served to clients (individually and as a set)
* reqJSON.go represents HTTP client requests to the HTTP server

* Each .json file should be formatted to match the format of bank-1.json, bank-2.json, and
bank-3.json in order to ensure the file is correctly read