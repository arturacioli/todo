package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
)

func (cli *Cli)handlerList(args []string){

	file,err := os.OpenFile(cli.Cfg.PathToCsv, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil{
		fmt.Printf("Error opening file %v", err)	
		return
	}
	
	for{
		file, err := file.Read()
		if errors.Is(err, io.EOF){
			break	
		}
	


	}

}
