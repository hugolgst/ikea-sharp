<h1 align="center">
  <br>
  <img src="https://i.imgur.com/lH4kZ5l.png" alt="IKEA# logo" width="400">
  <br>
  IKEA#
  <br>
</h1>

<h4 align="center">The new groundbreaking programming language</h4>

<p align="center">⚠️ This programming language is a prototype written in Golang, use it at your own risk</p>

<p align="center">
  <a href="#presentation">Presentation</a> —
  <a href="#how-to-use">How To Use</a> —
  <a href="#documentation">Documentation</a> —
  <a href="https://github.com/hugolgst/ikea-sharp/tree/master/examples">Examples</a> —
  <a href="#license">License</a>
</p>

## Presentation
This language is pronounced "Ikea Sharp", and its **only** goal is to disrupt you.

## How To Use
Clone the project and run
```go
go run main.go examples/math.ikea
```
Then follow the instructions

## Documentation

Here is an example to get the user entry:
```ikea
SMÅGLI FUNKÖN SKOGSFIBBLA Hello! SKOGSFIBBLA ÄPPLARÖ FJÄLLBO
SMÅGLI FUNKÖN SKOGSFIBBLA What is your name! SKOGSFIBBLA ÄPPLARÖ FJÄLLBO
TOSTERÖ FUNKÖN SKOGSFIBBLA name SKOGSFIBBLA ÄPPLARÖ FJÄLLBO
FULLSPÄCKAD FUNKÖN SKOGSFIBBLA Hello %s! SKOGSFIBBLA SMÅKALLT FUNKÖN name ÄPPLARÖ ÄPPLARÖ FJÄLLBO
```

Please see the examples [here](https://github.com/hugolgst/ikea-sharp/tree/master/examples)

### Methods
* TILLGÅNG key value : Saves the value in the key
* SMÅKALLT key : Gets the value of the given key
* TOSTERÖ key : Saves the console entry into the given key
* SMÅGLI String : Prints in the console
* FULLSPÄCKAD String... : Printf
* VÅRHOLMEN Int Int : Adds
* SMÅGÖRA Int Int : Subtracts
* ÄNGSLILJA Int Int : Multiply
* BLÖTSNÖ Int Int : Divide
* SNÖYRA Int Int : Random number between min and max
