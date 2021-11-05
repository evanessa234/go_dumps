package main

import (
    "fmt"
    "os/exec"
	"os"
)

func main() {
    app := "curl"
    api := "wttr.in/"
	arg := os.Args[1]
	arg_for_cmd := api+arg
    cmd := exec.Command(app, arg_for_cmd)
    stdout, err := cmd.Output()

    if err != nil {
        fmt.Println(err.Error())
        return
    }

    // Print the output
    fmt.Println(string(stdout))
}

// to run use below command:
// `go run weather_command_line_app.go <city name>`

