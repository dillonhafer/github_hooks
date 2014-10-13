package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

func updateMinecraft() error {
	return exec.Command("/home/dhafer/./mcupdate.sh").Run()
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Command Must be a POST\n")
			return
		}
		err := updateMinecraft()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	})
	err := http.ListenAndServe(":42000", nil)
	fmt.Fprintln(os.Stderr, err)
}
