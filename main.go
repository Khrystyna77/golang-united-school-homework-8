package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Arguments map[string]string

var useIdFlag = flag.String("id", "", "User id")
var useItemFlag = flag.String("item", "", "User data in JSON format")
var useOperationFlag = flag.String("operation", "", "Required operation")
var useFilenameFlag = flag.String("fileName", "", "Path to JSON-file with users' data")

func parseArgs() Arguments {
	flag.Parse()
	result := Arguments{}
	flag.Visit(func(flag *flag.Flag) {
		result[flag.Name] = flag.Value.String()
	})
	return result
}

type User struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func readAllUsers(fileName string) ([]User, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, nil
		}
		return nil, err
	}
	if len(data) == 0 {
		return nil, nil
	}
	var users []User
	err = json.Unmarshal(data, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func writeAllUsersToFile(fileName string, users []User) error {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	userdata, err := json.Marshal(users)
	if err != nil {
		return err
	}
	_, err = file.Write(userdata)
	return err
}

func add(args Arguments, writer io.Writer) error {
	inputJson := args["item"]
	if inputJson == "" {
		return fmt.Errorf("-item flag has to be specified")
	}
	var newUser User
	err := json.Unmarshal([]byte(inputJson), &newUser)
	if err != nil {
		return err
	}
	fileName := args["fileName"]
	if fileName == "" {
		return fmt.Errorf("-fileName flag has to be specified")
	}
	allUsers, err := readAllUsers(fileName)
	if err != nil {
		return err
	}

	id := newUser.Id
	found := -1
	for i, user := range allUsers {
		if user.Id == id {
			found = i
		}
	}
	if found > -1 {
		_, err = fmt.Fprintf(writer, "Item with id %s already exists", newUser.Id)
		if err != nil {
			panic(err)
		}
		return nil
	}

	allUsers = append(allUsers, newUser)
	result := writeAllUsersToFile(fileName, allUsers)
	return result
}

func list(args Arguments, writer io.Writer) error {
	fileName := args["fileName"]
	if fileName == "" {
		return fmt.Errorf("-fileName flag has to be specified")
	}
	allUsers, err := readAllUsers(fileName)
	if err != nil {
		return err
	}
	if len(allUsers) == 0 {
		return nil
	}
	userdata, err := json.Marshal(allUsers)
	if err != nil {
		return err
	}
	_, err = writer.Write(userdata)
	return err
}

func findById(args Arguments, writer io.Writer) error {
	id := args["id"]
	if id == "" {
		return fmt.Errorf("-id flag has to be specified")
	}
	fileName := args["fileName"]
	if fileName == "" {
		return fmt.Errorf("-fileName flag has to be specified")
	}
	allUsers, err := readAllUsers(fileName)
	if err != nil {
		return err
	}

	found := -1
	for i, user := range allUsers {
		if user.Id == id {
			found = i
		}
	}
	if found < 0 {
		_, err = writer.Write([]byte(""))
		return nil
	}

	user, err := json.Marshal(allUsers[found])
	if err != nil {
		return err
	}
	_, err = writer.Write(user)
	return err
}

func remove(args Arguments, writer io.Writer) error {
	id := args["id"]
	if id == "" {
		return fmt.Errorf("-id flag has to be specified")
	}
	fileName := args["fileName"]
	if fileName == "" {
		return fmt.Errorf("-fileName flag has to be specified")
	}
	allUsers, err := readAllUsers(fileName)
	if err != nil {
		return err
	}

	found := -1
	for i, user := range allUsers {
		if user.Id == id {
			found = i
		}
	}
	if found < 0 {
		_, err = fmt.Fprintf(writer, "Item with id %s not found", id)
		if err != nil {
			panic(err)
		}
		return nil
	}

	result := make([]User, 0, len(allUsers)-1)
	for _, user := range allUsers {
		if user.Id != id {
			result = append(result, user)
		}
	}
	return writeAllUsersToFile(fileName, result)
}

func Perform(args Arguments, writer io.Writer) error {
	operation, ok := args["operation"]
	if !ok || operation == "" {
		return fmt.Errorf("-operation flag has to be specified")
	}
	switch operation {
	case "add":
		return add(args, writer)
	case "list":
		return list(args, writer)
	case "findById":
		return findById(args, writer)
	case "remove":
		return remove(args, writer)
	default:
		return fmt.Errorf("Operation %v not allowed!", operation)
	}
}

func main() {
	err := Perform(parseArgs(), os.Stdout)
	if err != nil {
		panic(err)
	}
}
