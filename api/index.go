package handler

import (
	"fmt"
	"net/http"

	"github.com/iawia002/lux/extractors"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	var (
		data []*extractors.Data
		err  error
	)

	data, err = extractors.Extract("https://www.youtube.com/watch?v=Gnbch2osEeo", extractors.Options{})
	fmt.Fprintf(w, "<h1>Hello from Go!</h1>"+data[0].Title+err.Error())

}
