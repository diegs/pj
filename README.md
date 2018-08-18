# pj

`pj` is a simple command line tool that reads JSON from stdin, round-trips it
through the decoder, and pretty-prints it to stdout.

## Installation

`go get -u github.com/diegs/pj`

## Usage

`pj < input.json > output.json`

or

`cat input.json | pj`

## Why not just use jq?

`jq` does not handle numbers larger than 2^53, which truncates 64-bit integers.
