package main

import (
	"IM/pkg/common/cmd"
	"IM/pkg/program"
)

func main() {
	if err := cmd.NewApiCmd().Exec(); err != nil {
		program.ExitWithError(err)
	}
}
