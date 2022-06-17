package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

type Userslist struct {
	ID    uint   `json:"id"`
	email string `json:"email"`
	age   uint   `json:"age`
}
type Arguments map[string]string

func parseArgs() Arguments {
	var OperationFlag *string
	var ItemFlag *string
	var IdFlag *string
	var FileNameFlag *string

	//var args Arguments = make(map[string]string)

	IdFlag = flag.String("id", "", "User's id")
	FileNameFlag = flag.String("fileName", "users.json", "Path to JSON-file")
	OperationFlag = flag.String("operation", "", "action to be done")
	ItemFlag = flag.String("item", "", "Lists of users")
	flag.Parse()

	return map[string]string{"operation": *OperationFlag, "item": *ItemFlag, "id": *IdFlag, "fileName": *FileNameFlag}
}

//don't touch
func Perform(args Arguments, writer io.Writer) error {
	//scaner := bufio.NewScanner(args)

	f, err := os.OpenFile("users.json", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	args["item"] = "[{}]"
	user := []Userslist{}

	file, err := os.OpenFile("users.json", os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	itemBytes, err := json.Marshal(&user)
	if _, err := io.WriteString(file, string(itemBytes)); err != nil {
		return fmt.Errorf("Write json %v to file %v finished with error: %w\n", string(itemBytes), args["users.json"], err)
	}

	return nil
}

// 	return 0

// }

func main() {

	err := Perform(parseArgs(), os.Stdout)
	if err != nil {
		panic(err)
	}
	//don't touch

}
