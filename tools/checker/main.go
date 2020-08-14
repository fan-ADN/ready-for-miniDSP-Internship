package main

import (
	"os"

	"github.com/fan-ADN/ready-for-miniDSP-Internship/tools/checker/cmd"
)

func main() {
	ret := cmd.Execute(os.Args[1:])
	os.Exit(ret)
}
