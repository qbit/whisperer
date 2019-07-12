# Whisperer

Whisperer is a simple Go program that makes HTTP request constantly in order to generate random HTTP/DNS traffic noise.

> Whisperer is a project inspired by [Noisy](https://github.com/1tayH/noisy).

## Example

```bash
whisperer -v -d 3s
```

## Install

### Go

```bash
go get github.com/danielkvist/whisperer
```

### Cloning the repository

```bash
# First, clone the repository
git clone https://github.com/danielkvist/whisperer

# Then navigate into the whisperer directory
cd whisperer

# Run
go run main.go
```

## Docker

### Pulling Image

To use whisperer as a Docker container you can pull the image with the following command:

```bash
docker image pull danielkvist/whisperer
```

> Note that the image ```danielkvist/whisperer``` uses the urls file from this repository. So it is not a valid option if you want to customize the URLs that whisperer is going to visit.

### Building Image

```bash
# First, clone the repository
git clone https://github.com/danielkvist/whisperer

# Then navigate into the whisperer directory
cd whisperer

# Modify the urls.txt file if you want
vim urls.txt

# Build the Docker Image from the Dockerfile inside the repository
docker image build -t whisperer .

# Run
docker container run whisperer
```

## Options

Whisperer can accept a number of command line arguments:

```text
$ whisperer --help
whisperer makes HTTP request constantly in order to generate random HTTP/DNS traffic noise.

Usage:
  whisperer [flags]

Flags:
  -a, --agent string     user agent (default "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:67.0) Gecko/20100101 Firefox/67.0")
  -d, --delay duration   delay between requests
  -g, --goroutines int   number of goroutines (default 1)
  -h, --help             help for whisperer
      --urls string      simple .txt file with URL's to visit (default "./urls.txt")
  -v, --verbose          enables verbose mode
```

## URLs file

This file is from which Whisperer will extract the different URLs that will be visiting.

> You can see an example of how this file should be [here](https://github.com/danielkvist/whisperer/blob/master/urls.txt).

## Help is always welcome!

If you know about anything else I can improve or add please, don't hesitate to let me know!
