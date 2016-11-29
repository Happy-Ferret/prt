package main

import (
	"fmt"
	"os"

	"github.com/chiyouhen/getopt"
	"github.com/onodera-punpun/prt/commands"
)

func main() {
	// Define opts
	shortopts := ""
	longopts := []string{}

	// Read out opts
	_, vals, _ := getopt.Getopt(os.Args, shortopts, longopts)

	if len(vals) == 1 {
		fmt.Fprintln(os.Stderr, "Missing command, use help for a list of commands!")
		os.Exit(1)
	}

	switch vals[1] {
	case "help":
		fmt.Println("Usage: prt command [arguments]")
		fmt.Println("")
		fmt.Println("commands:")
		fmt.Println("  depends                 list dependencies recursivly")
		fmt.Println("  diff                    list outdated packages")
		fmt.Println("  build                   build and install ports")
		fmt.Println("  info                    print port information")
		fmt.Println("  list                    list ports")
		fmt.Println("  loc                     prints port locations")
		fmt.Println("  patch                   patch ports")
		fmt.Println("  provide                 search ports for files")
		fmt.Println("  remove                  remove installed ports")
		fmt.Println("  pull                    pull in ports")
		fmt.Println("  sysup                   update outdated packages")
		fmt.Println("  help                    print help and exit")
		os.Exit(0)
	case "depends":
		commands.Depends(os.Args[1:])
		os.Exit(0)
		//	case "build":
		//		commands.Build(os.Args[1:])
		//		os.Exit(0)
	case "info":
		commands.Info(os.Args[1:])
		os.Exit(0)
		//	case "list":
		//		command.List(os.Args[1:])
		//		os.Exit(0)
	case "loc":
		commands.Loc(os.Args[1:])
		os.Exit(0)
		//	case "patch":
		//		command.Patch(os.Args[1:])
		//		os.Exit(0)
		//	case "provide":
		//		command.Provide(os.Args[1:])
		//		os.Exit(0)
		//	case "remove":
		//		command.Remove(os.Args[1:])
		//		os.Exit(0)
	case "pull":
		commands.Pull(os.Args[1:])
		os.Exit(0)
		//	case "sysup":
		//		command.Sysup(os.Args[1:])
		//		os.Exit(0)
	default:
		fmt.Fprintln(os.Stderr, "Invalid command, use help for a list of commands!")
		os.Exit(1)
	}
}
