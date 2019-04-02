package persist_test

import (
	"fmt"
	"io/ioutil"

	"github.com/axamon/persist"
)

const value string = `{"foo": "bar"}`

var storingFile = "./persiststorefile_tobedeleted"

func ExampleSave() {
	//defer os.Remove(storingFile)

	persist.Save(storingFile, value)

	fileRead, _ := ioutil.ReadFile(storingFile)
	str := string(fileRead)
	fmt.Printf(str)
	// Output:
	// "{\"foo\": \"bar\"}"
}
