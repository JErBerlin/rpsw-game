## A server to play the Rock-Paper-Scissors game

### Overview

This is a game server that offers the possibility to play _rock, paper, scissors_ against the computer by means of a post request to its API

- Scissors beats paper
- Paper beats rock
- Rock beats scissors

## How to run..

If you have Go installed on your computer, you can just get this repository and run this game locally (you can install the go compiler here: https://golang.org/doc/install).

If you would not like to install a Go compiler, please read further in the section about how to run the application in a docker container. 

###  the tests

`go test`

###  the compiler

`go build`

###  the program

To run the game server, you can just type:

`go run rps-game`

(Observe that it will be listening at port 8080. If you get an error when starting the server because the port is already in use, try changing the constant value of the port in the `api.go` file)

### How to send your move to the api

When testing locally, the API will be listening and serving http at port 8080 of the localhost.

- to send your move, make a POST request at http://localhost:8080/play

You should specify a valid move in JSON format in the body of the request: for instance,

`{
"Throw": "paper"
}`

Valid moves are: "rock", "paper", "scissors" (The API is case insensitive.)

You can use a command-line tool like `curl`:

`curl -X POST -H "Content-Type: application/json" -d '{"Throw": "paper"}' http://localhost:8080/play`

You would just need to replace `"paper"` with whatever move you want to make. The response from the server will appear in your terminal after you send the request

### How to run the app from the Docker image

There is a provided Dockerfile to run the application in a docker container.

To build a new image you can type this from the directory where the Dockerfile is:

`docker build -t rps-game .`

To run the docker container (you can also omit the `-it` option):

`docker run -it -p 8080:8080 --name=rps-game rps-game`

You can stop the container pressing ctrl+C or typing following from a new terminal tab:

`docker stop rps-game`
