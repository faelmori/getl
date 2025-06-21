package main

import (
	"fmt"
	"os"

	gl "github.com/rafa-mori/getl/logger"
)

func main() {
	if err := RegX().Execute(nil); err != nil {
		gl.Log("Error",fmt.Sprintf("Error: %v", err))
		os.Exit(1)
	}
}
