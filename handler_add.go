package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

type todo struct{
	task string
	complete bool
	addedAt time.Time
}

func (cli *Cli)HandlerAdd(args [] string){

	nextId, err := getNextId(cli.Cfg.PathToCsv)
	if err != nil{
		fmt.Printf("Erro ao gerar id %v\n",err)
		return
	}
	row := []string{
		strconv.Itoa(nextId),
		args[0],
		strconv.FormatBool(false),
		time.Now().String(),
	}

	file,err := os.OpenFile(cli.Cfg.PathToCsv, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil{
		fmt.Printf("Ocorreu um erro: %v\n", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	err = writer.Write(row)
	if err != nil{
		fmt.Printf("Erro ao escrever no CSV: %v\n",err)
		return
	}

	writer.Flush()
	fmt.Printf("Task salva com sucesso!")		
}

func getNextId(filePath string) (int,error){
	file, err := os.Open(filePath)	

	if errors.Is(err, os.ErrNotExist){
		return 1, nil
	}
	if err != nil{
		return 0, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lastId := 0

	for {
		record, err := reader.Read()
		if errors.Is(err, io.EOF){
			break
		}
		if err != nil{
			return 0, err
		}

		if len(record) > 0 {
			id, err := strconv.Atoi(record[0])
			if err == nil && id > lastId{
				lastId = id
			}
		}


	}
	return lastId + 1, nil
}
