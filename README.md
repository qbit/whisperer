# Whisperer

[![Go Report Card](https://goreportcard.com/badge/github.com/qbit/whisperer)](https://goreportcard.com/report/github.com/qbit/whisperer)
[![CircleCI](https://circleci.com/gh/qbit/whisperer.svg?style=svg)](https://circleci.com/gh/qbit/whisperer)
[![GoDoc](https://godoc.org/github.com/qbit/whisperer?status.svg)](https://godoc.org/github.com/qbit/whisperer)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](http://makeapullrequest.com)

Whisperer is a simple Go program that makes HTTP request constantly in order to generate random HTTP/DNS traffic noise.

> Whisperer is a project inspired by [Noisy](https://github.com/1tayH/noisy).

This fork has an embedded urls.txt list and is intended to be run from GoKrazy images.

## Example

```bash
whisperer -v -d 3s
```

## Install

### Go

```bash
go get github.com/qbit/whisperer
```

### Cloning the repository

```bash
# First, clone the repository
git clone https://github.com/qbit/whisperer

# Then navigate into the whisperer directory
cd whisperer

# Run
go run main.go
```

## Options

Whisperer can accept a number of command line arguments:

```text
$ whisperer --help
whisperer makes HTTP request constantly in order to generate random HTTP/DNS traffic noise.

Usage:
  whisperer [flags]

Flags:
  -a, --agent string       user agent (default "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:67.0) Gecko/20100101 Firefox/67.0")
      --debug              prints error messages
  -d, --delay duration     delay between requests (default 1s)
  -g, --goroutines int     number of goroutines (default 1)
  -h, --help               help for whisperer
  -p, --proxy string       proxy URL
  -r, --random             random delay between requests
  -t, --timeout duration   max time to wait for a response before canceling the request (default 3s)
      --urls string        simple .txt file with URL's to visit (default "./urls.txt")
  -v, --verbose            enables verbose mode
```

## URLs file

This file is from which Whisperer will extract the different URLs that will be visiting.

> You can see an example of how this file should be [here](https://github.com/qbit/whisperer/blob/master/cmd/urls.txt).

