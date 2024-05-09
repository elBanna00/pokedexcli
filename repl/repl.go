package repl

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
)

func Startuprepl() {
	color.Green("Welcome To Pokedex! \nType 'help' for more information")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()

		input := scanner.Text()

		cleanInput := Cleanstr(input)

		if len(cleanInput) == 0 {
			continue
		}

		command := cleanInput[0]

		switch command {
		case "help":
			fmt.Println("Welcome to pokedex help menu")
			fmt.Println("Here are all the available commands")
			color.Yellow("- help")
			color.Yellow("- exit")
			fmt.Println("")
		case "exit":
			os.Exit(0)
		default:
			color.Red("Invalid Command %v", command)
		}
	}
}

type cliCommand struct {
	name        string
	description string
}

func Getcommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
		},
	}
}
func Cleanstr(str string) []string {

	lowerCaseStr := strings.ToLower(str)
	inputWords := strings.Fields(lowerCaseStr)
	return inputWords
}
