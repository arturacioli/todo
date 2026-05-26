package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/dustin/go-humanize"
)

func (cli *Cli)HandlerList(args []string){
	showAll := false

	for _, arg := range args {
		if arg == "--all" || arg == "-a" {
			showAll = true
		}
	}

	w := tabwriter.NewWriter(os.Stdout,0,0,2,' ',0)
	// header
	if showAll {
		fmt.Fprintln(w, "ID\tTASK\tCOMPLETED\tADDED_AT")
	} else {
		fmt.Fprintln(w, "ID\tTASK\tADDED_AT")
	}

	for _,task := range cli.Tasks {
		if !showAll && task.Complete {
			continue			
		}
		if showAll {
			fmt.Fprintln(w,
				task.Id+"\t"+
					task.Task+"\t"+
					fmt.Sprintf("%v", task.Complete)+"\t"+
					humanize.Time(task.AddedAt),
			)
		} else {
			fmt.Fprintln(w,
				task.Id+"\t"+
					task.Task+"\t"+
					humanize.Time(task.AddedAt),
			)
		}
	}

	w.Flush()

}
