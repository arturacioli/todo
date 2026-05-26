package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/arturacioli/todo/internal/storage"
	"github.com/arturacioli/todo/internal/task"
)



func (cli *Cli)HandlerAdd(args [] string){

	nextId, err := cli.GetNextId()
	if err != nil{
		fmt.Printf("Erro ao gerar id %v\n",err)
		return
	}

	row := task.Task{
		Id: strconv.Itoa(nextId),
		Task: args[0],
		Complete: false,
		AddedAt: time.Now(),
	}

	cli.Tasks = append(cli.Tasks, row)
	err = storage.WriteTasks(cli.Tasks)
	if err != nil {
		fmt.Printf("Erro ao salvar tasks: %v\n", err)
		return
	}

}

func (cli *Cli)GetNextId() (int,error){
	lastID := 0
	for _, t := range cli.Tasks {
		id, err := strconv.Atoi(t.Id)
		if err != nil {
			continue
		}

		if id > lastID {
			lastID = id
		}
	}

	return lastID + 1, nil
}
