[![GoDoc](https://godoc.org/github.com/andskur/pswd-hashing-tools?status.svg)](https://godoc.org/github.com/andskur/pswd-hashing-tools)
![Supported Go Versions](https://img.shields.io/badge/Go-1.10%2C%201.11-lightgrey.svg)
# Password-hashing tools

Command line utility with which you can play with various crypto algorithms for make hash from given password string and compare its matching after with different hashes values.

### Currently Supported Algorithms:
* [Bcrypt](https://wikipedia.org/wiki/Bcrypt) - based on the Blowfish cipher, and presented at USENIX in 1999
* [Scrypt](https://wikipedia.org/wiki/Scrypt) - created by Colin Percival, originally for the online backup service
* [Argon2](https://wikipedia.org/wiki/Argon2) - winner of the Password Hashing Competition in July 2015

## Installation
### Building from sources (Windows, MacOS, or Linux)

**Requirement:**
* [Go 1.10+](https://golang.org/doc/install)
* [Glide](https://github.com/Masterminds/glide#install)

```
$ go get github.com/andskur/pswd-hashing-tools
$ cd $GOPATH/src/github.com/andskur/pswd-hashing-tools
$ glide install
$ go build
```

## Usage
```
$ ./paswd-hashing-tools -h
Usage: [COMMAND][ARGUMENT][-FLAGS]

Available Commands:
  compare   [password] ['hash']     Compare given string with a given hash
  hash      [password]              Hash given string
  help                              Help about any command

Password and hash arguments are optional, you can type it in stdin after command execution


Flags:
  -a, --algorithm string   Crypto algorithm to use (default "bcrypt")
  -h, --help               help for this command
  -p, --prehash            Enable prehash SHA256 function
  
Available algorithms:
  bcrypt
  scrypt
  argon2

Use " [command] --help" for more information about a command.
```

### Examples
**Hashing password:**
```
$ ./pswd-hashing-tools hash -a scrypt
Using Scrypt hashing algorithm 
Enter password for hash:
qwerty123
Resulting: 16384$8$1$65e43a1b1ba8983e4ae2c92f2a78b298$2f136899f314c9ff1ec57c175fe824eea37c6499facab44dc8b680e52343a357
```
**Compare password with hash :**
```
$ ./pswd-hashing-tools compare qwerty123 '$2a$10$uXvEFzrJyr6dzEmBORB/s.tky7Wcn2Q7AX1Ap32o2SWm6GoTs6wQG' -a bcrypt
Using Bcrypt hashing algorithm 
Hash and password are matching
```

## Built With

* [Cobra](https://github.com/spf13/cobra) - A Commander for modern Go CLI interactions
* [Glide](https://github.com/Masterminds/glide) - Package Management for Golang
* [Simple-scrypt](https://github.com/elithrar/simple-scrypt) - A convenience library for working with Scrypt hashing algorithm

## Authors

* **Andrey Skurlatov** - [andskur](https://github.com/andskur)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details