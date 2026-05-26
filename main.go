package main

import (
	"fmt"
	"os"

	"github.com/arturacioli/todo/internal/storage"
	"github.com/arturacioli/todo/internal/task"
)

type Command struct{
	Handler func(args []string)
}

type Cli struct{
	Commands  map[string]Command
	Arguments []string
	Tasks     []task.Task 
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

	tasks, err := storage.LoadTasks()
	if err != nil{
		fmt.Printf("Error loading/creating csv: %v", err)
		os.Exit(1)
	}

	cli := Cli{
		Commands: make(map[string]Command),
		Arguments: os.Args[2:],
		Tasks: tasks,
	}


	cli.AddCommand("add",cli.HandlerAdd)
	cli.AddCommand("list",cli.HandlerList)
	cli.AddCommand("complete",cli.HandlerComplete)
	cli.AddCommand("delete",cli.HandlerDelete)

	
	command := os.Args[1]
	cli.ExecuteCommand(command)
}


