## A server to play the classic Rock-Paper-Scissors game

### Overview

This is a game server that offers the possibility

    1. to play rock, paper, scissors against the computer

    2. by means of a post request to its API

The rock, paper, scissors game:

- Scissors beats paper
- Paper beats rock and well
- Rock beats scissors

## How to run..

If you have Go installed on your computer, you can just get this repository and run this game locally (you can install the go compiler here: https://golang.org/doc/install).

###  the tests

`go test`

###  the compiler

`go build`

###  the program

To run the game server, you can just type:

`go run rpsw-game`

(Observe that it will be listening at port 8080. If you get an error when starting the server because the port is already in use, try changing the constant value of the port in the `api.go` file)

### How to send your move to the api

When testing locally, the API will be listening and serving http at port 8080 of the localhost.

- to send your move, make a POST request at http://localhost:8080/play

You should specify a valid move in JSON format in the body of the request: for instance,

`{
"Throw": "paper"
}`

Valid moves are: "rock", "paper" and "scissors". (The API is case insensitive.)