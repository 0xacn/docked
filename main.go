package main

import (
  "os"
  "fmt"
  "github.com/fatih/color"
)
var (
	version = "dev"
	helpCmd = `whack - a better process killer
						
	Usage
		$ whack <pid> <name> <port>
	Options
		--force								Force kill a process
		--verbose							Show process arguments
	Examples
	$ whack 1445
	$ whack discord
	$ whack :3000
	To terminate a specific port, add a colon. e.g: :3000.`
)

func main() {
  if len(os.Args) > 4 {
    handleErrStr("Too many arguments >3")
    fmt.Println(helpCmd)
    return
  }
}

func handleErr(err error) {
  handleErrStr(err.Error())
}

func handleErrStr(str string) {
  _, _ = fmt.Fprintln(os.Stderr, color.RedString("error: ")+str)
}
