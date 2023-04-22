package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/macabu/inamedparam"
)

func main() {
	singlechecker.Main(inamedparam.Analyzer)
}
