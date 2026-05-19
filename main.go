package main

import (
	"fmt"
	"os"
)

type Cfg struct{
	PathToCsv string	
}

type Command struct{
	Handler func(args []string)
}

type Cli struct{
	Commands  map[string]Command
	Arguments []string
	Cfg       *Cfg
}

func (cli *Cli)AddCommand(name string, handler func(args []string)){
	newCmd := Command{
		Handler: handler,
	}

	cli.Commands[name] = newCmd
}

func (cli *Cli)ExecuteCommand(commandName string) error{
	if _,ok := cli.Commands[commandName]; !ok{
		return fmt.Errorf("Command not mapped")
	}

	cli.Commands[commandName].Handler(cli.Arguments)
	return nil
}

func main(){
	if len(os.Args) < 2 {
		fmt.Println("Usage: <command> [args]")
		os.Exit(1)
	}
	cfg := Cfg{
		PathToCsv: "todo.csv",
	}
	cli := Cli{
		Commands: make(map[string]Command),
		Arguments: os.Args[2:],
		Cfg: &cfg,
	}


	cli.AddCommand("add",cli.HandlerAdd)

	
	command := os.Args[1]
	cli.ExecuteCommand(command)
}


