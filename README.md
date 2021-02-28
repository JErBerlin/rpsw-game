## A server to play the Rock-Paper-Scissors-Well game

### Overview

This is a game server that offers the possibility to play _rock, paper, scissors, well_ against the computer by means of a post request to its API

The _rock, paper, scissors, well_ game is a variation of the classical _rock, paper, scissors_ game:

- Scissors beats paper
- Paper beats rock and well
- Rock beats scissors
- Well beats rock and scissors

You can read more about this game here:

- [Circulant games](https://link.springer.com/article/10.1007/s11238-014-9478-4)
- [Rock, Paper, Scissors, Well](https://math.stackexchange.com/questions/410558/rock-paper-scissors-well)

## How to run..

If you have Go installed on your computer, you can just get this repository and run this game locally (you can install the go compiler here: https://golang.org/doc/install).

If you would not like to install a Go compiler, please read further in the section about how to run the application in a docker container. 

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

Valid moves are: "rock", "paper", "scissors" and "well". (The API is case insensitive.)

### How to run the app from the Docker image

There is a provided Dockerfile to run the application in a docker container.

To build a new image you can type this from the directory where the Dockerfile is:

`docker build -t rpsw-game .`

To run the docker container (you can also omit the `-it` option):

`docker run -it -p 8080:8080 --name=rpsw-game rpsw-game`

You can stop the container pressing ctrl+C or typing following from a new terminal tab:

`docker stop rpsw-game`
