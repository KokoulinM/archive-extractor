package main

import (
	"fmt"

	"github.com/gen2brain/go-unarr"
)

func main() {
	var pathFrom string
	var pathTo string

	fmt.Print("Specify the correct path to the archive: ")

	fmt.Scan(&pathFrom)

	fmt.Print("Specify where to unpack the archive: ")

	fmt.Scan(&pathTo)

	a, err := unarr.NewArchive(pathFrom)
	if err != nil {
		panic(err)
	}
	defer a.Close()

	_, err = a.Extract(pathTo)
	if err != nil {
		panic(err)
	}

	fmt.Print("The archive has been successfully unpacked")
}
