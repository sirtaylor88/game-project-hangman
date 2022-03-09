## Description

The player has to guess the full word by suggesting letters within a certain number of guesses.

### Build for Windows
##### Win 64 bit

    GOOS=windows GOARCH=amd64 go build -o bin/hangman-win64.exe main.go

##### Win 32 bit

    GOOS=windows GOARCH=386 go build -o bin/hangman-win32.exe main.go

### Build for macOS

    GOOS=darwin GOARCH=amd64 go build -o bin/hangman-mac64 main.go

### Build for Linux

    GOOS=linux GOARCH=amd64 go build -o bin/hangman-linux64 main.go
