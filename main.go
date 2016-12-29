package main

import (
	"fmt"
	"os"

	"github.com/onodera-punpun/prt/commands"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "Missing command, use help for a list of commands!")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "help":
		fmt.Println("Usage: prt command [arguments]")
		fmt.Println("")
		fmt.Println("commands:")
		fmt.Println("  depends                 list dependencies recursivly")
		fmt.Println("  diff                    list outdated packages")
		fmt.Println("  info                    print port information")
		fmt.Println("  install                 build and install packages")
		fmt.Println("  list                    list ports")
		fmt.Println("  loc                     print port locations")
		//fmt.Println("  patch                   patch ports")
		fmt.Println("  prov                    search ports for files")
		//fmt.Println("  remove                  remove installed ports")
		fmt.Println("  pull                    pull in ports")
		fmt.Println("  sysup                   update outdated packages")
		fmt.Println("  help                    print help and exit")
		os.Exit(0)
	case "depends":
		commands.Depends(os.Args[2:])
		os.Exit(0)
	case "diff":
		commands.Diff(os.Args[2:])
		os.Exit(0)
	case "info":
		commands.Info(os.Args[2:])
		os.Exit(0)
	case "install":
		commands.Install(os.Args[2:])
		os.Exit(0)
	case "list":
		commands.List(os.Args[2:])
		os.Exit(0)
	case "loc":
		commands.Loc(os.Args[2:])
		os.Exit(0)
		//	case "patch":
		//		command.Patch(os.Args[1:])
		//		os.Exit(0)
	case "prov":
		commands.Prov(os.Args[2:])
		os.Exit(0)
		//	case "remove":
		//		command.Remove(os.Args[1:])
		//		os.Exit(0)
	case "pull":
		commands.Pull(os.Args[2:])
		os.Exit(0)
	case "sysup":
		commands.Sysup(os.Args[2:])
		os.Exit(0)
	default:
		fmt.Fprintln(os.Stderr, "Invalid command, use help for a list of commands!")
		os.Exit(1)
	}
}
