package main

import (
	"fmt"

	"github.com/axamon/persist"
)

const value string = `{"foo": "bar"}`

var storingFile = "./persiststorefile_tobedeleted"

func save() {
	data := value

	persist.Save(storingFile, data)
}

func load() {
	var data string

	persist.Load(storingFile, data)

	fmt.Println(data)
}

func main() {
	save()
	load()

}
