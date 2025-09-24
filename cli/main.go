package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mathcale/what-the-cron/internal/pkg/cron"
	usecase "github.com/mathcale/what-the-cron/internal/usecase/cron"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: what-the-cron <cron-expression>")
		os.Exit(1)
	}

	uc := usecase.NewCronUseCase(cron.NewAdapter())

	expr := strings.Join(os.Args[1:], " ")

	result, err := uc.Execute(expr)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(result.Description)
	fmt.Printf("Next execution: %s\n", result.FormattedNextExecution())
}
