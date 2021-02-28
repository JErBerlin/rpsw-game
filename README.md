Rock, Paper, Scissors Game Server

## A server to play the classic Rock-Paper-Scissors game

### Overview

This is a game server that offers the possibility

    1. to play rock, paper, scissors against the computer

    2. by means of a post request to its API

The rock, paper, scissors game:

- Scissors beats paper
- Paper beats rock and well
- Rock beats scissors

### How to use it

### Accessing public and private API for posting, viewing and editing a single message

When testing locally, the API will be listening and serving http at port 8080 of the localhost.

- to send your move, make a POST request at http://localhost:8080/play

You should specify a valid move in JSON format in the body of the request: for instance,

`{
"Throw": "paper"
}`

Valid moves are: "rock", "paper" and "scissors". (The API is case insensitive.)