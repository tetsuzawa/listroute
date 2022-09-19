package main

import (
	"github.com/tetsuzawa/listroute"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(listroute.Analyzer) }
