package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/fatih/color"

	"github.com/iawia002/lux/app"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	var args = []string{"-i", "https://www.youtube.com/watch?v=6g4dkBF5anU"}
	if err := app.New().Run(args); err != nil {
		fmt.Fprintf(
			color.Output,
			"Run %s failed: %s\n",
			color.CyanString("%s", app.Name), color.RedString("%v", err),
		)
		os.Exit(1)
	}

	fmt.Fprintf(w, "<h1>Hello from Go!</h1>")

}
