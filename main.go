package main

import (
	"ascap/commands"
	"ascap/database"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type CommandType int

const (
	Song CommandType = iota
	Platform
	Upload
)

func parseCommand(arg string) (CommandType, error) {
	cleaned := strings.ToLower(arg)
	switch cleaned {
	case "song":
		return Song, nil
	case "platform":
		return Platform, nil
	case "upload":
		return Upload, nil
	}
	return Song, errors.New("invalid arg")
}

func main() {
	name := flag.String("name", "", "name")
	flag.Parse()
	fmt.Println("name is ", *name)

	if err := godotenv.Load(); err != nil {
		fmt.Print("env could not be loaded")
		return
	}

	db, err := database.EstablishConnection()
	if err != nil {
		fmt.Println("error establishing connection:", err)
		os.Exit(1)
	}

	args := os.Args
	if len(args) < 2 {
		fmt.Println("please provide a command")
		return
	}
	arg1 := args[1]
	command, err := parseCommand(arg1)
	if err != nil {
		fmt.Println(err)
		return
	}

	switch command {
	case Song:
		commands.ReadSong(db, *name)
		return
	case Platform:
		commands.ReadPlatform(db, *name)
	}

	// upload // requires: file_path + distro type
	// song // requires: song_name
	// platform // requires: name optional(song)
}
