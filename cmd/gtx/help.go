package main

import (
	"fmt"
	"os"
	"strings"
)

type helpCmd struct{}

var help = helpCmd{}

func (h helpCmd) name() string { return "help" }

func (h helpCmd) short() string { return "Display help text" }

func (h helpCmd) long() string {
	var sb strings.Builder
	sb.WriteString(`usage: gtx COMMANDS

GTX helps create contexts for Go applications to save config
for different environments

Commands:`)

	for _, v := range cmds {
		sb.WriteString(fmt.Sprintf("\n%20s\t%s", bold.Sprint(v.name()), gray.Sprint(v.short())))
	}
	sb.WriteRune('\n')

	return sb.String()
}

func (h helpCmd) run(exit int, args []string) {
	switch len(args) {
	case 0:
		fmt.Print(h.long())
	case 1:
		args[0] = strings.ToLower(args[0])
		for _, v := range cmds {
			if args[0] == v.name() {
				fmt.Print(v.long())
				os.Exit(exit)
			}
		}
		fmt.Printf("unknown command %s", args[0])
		os.Exit(1)
	default:
		fmt.Println("too many args to help")
	}

	os.Exit(exit)
}
