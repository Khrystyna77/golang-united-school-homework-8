package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
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

// func add(string) byte {
// 	user1 := []Userslist{}
// 	buff, err := bytes.NewBufferString()
// 	buff.WriteString(&user)
// 	res := os.Stdout
// 	data = append(data[:closingBraceIdx], ins...)

// }

//don't touch
func Perform(args Arguments, writer io.Writer) error {
	// var ItemFlag *string
	// ItemFlag = flag.String("item", "", "Lists of users")
	// flag.Parse()
	//scaner := bufio.NewScanner(args)
	if args["fileName"] == "" {
		return fmt.Errorf("Missing filename")
	}
	if args["item"] == "" {
		return fmt.Errorf("Please type items")
	}

	file, err := os.OpenFile("users.json", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}

	file, err = os.OpenFile("users.json", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}

	user := []Userslist{}

	itemBytes, err := json.Marshal(&user)

	if err != nil {
		return fmt.Errorf("not enaught data  %v", user, err)
	}
	//write data trying
	if _, err := io.WriteString(file, strings.ToLower(string(itemBytes))); err != nil {
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
