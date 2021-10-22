package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/manifoldco/promptui"
)

type callbackfunc func(interface{}) error

func promptInputString(f callbackfunc) error {
	validate := func(input string) error {
		if input == "" {
			return nil
		}
		_, err := strconv.ParseInt(input, 10, 32)
		if err != nil {
			return errors.New("invalid number")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Input your course ID(default is " + HW_ID + ")",
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return err
	}

	err = f(result)
	if err != nil {
		return err
	}
	return nil
}

var totalJob []string = []string{"Download All Submit File", "Upload Score"}

func promptSelectJob() (int, error) {
	prompt := promptui.Select{
		Label: "Select Job",
		Items: totalJob,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return -1, err
	}

	var id int = -1
	for i, j := range totalJob {
		if result == j {
			id = i
		}
	}

	if id == -1 {
		return id, errors.New("something went wrong")
	}

	return id, nil
}
