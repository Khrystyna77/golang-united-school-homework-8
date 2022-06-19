package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
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

func add(args Arguments, writer io.Writer) error {
	input := args["item"]

	if input == "" {
		return fmt.Errorf("-item flag has to be specified")
	}
	var newUser Userslist
	err := json.Unmarshal([]byte(input), &newUser)
	if err != nil {
		return err
	}
	filej := "users.json"
	fileName1 := args["fileName"]
	if fileName1 == "" {
		return fmt.Errorf("-fileName flag has to be specified")
	}

	var allUsers1 []string
	allUsers1 = append(allUsers1, input)
	//Newuser := []Userslist{}
	f, err := ioutil.ReadFile(filej)
	if err != nil {
		log.Fatal(err)
	}
	data := []Userslist{}
	json.Unmarshal(f, &data)
	//ID := args["id"]
	newStruct := &Userslist{}

	data = append(data, *newStruct)
	dataBytes, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(filej, dataBytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

//don't touch
func Perform(args Arguments, writer io.Writer) error {

	input := args["item"]

	if input == "" {
		return fmt.Errorf("-item flag has to be specified")
	}
	fileName1 := args["fileName"]
	if fileName1 == "" {
		return fmt.Errorf("-fileName flag has to be specified")
	}
	if args["operation"] == "" {
		return fmt.Errorf("-operation flag has to be specified")
	}

	file, err := os.OpenFile("users.json", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}

	// file, err = os.OpenFile("users.json", os.O_RDWR|os.O_CREATE, 0755)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	user := []Userslist{}

	itemBytes, err := json.Marshal(&user)

	if err != nil {
		return fmt.Errorf("not enaught data  %v", err)
	}
	//write data trying
	if _, err := io.WriteString(file, strings.ToLower(string(itemBytes))); err != nil {
		return fmt.Errorf("Write json %v to file %v finished with error: %w\n", string(itemBytes), args["users.json"], err)
	}

	if args["operation"] == "add" {
		return add(args, writer)
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
