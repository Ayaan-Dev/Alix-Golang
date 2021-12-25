package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token  string
	Prefix string

	config *configStruct
)

type configStruct struct {
	Token  string `json:"Token"`
	Prefix string `json:"Prefix"`
}

func ReadConfig() error {
	fmt.Println("Reading Config File!")

	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println(string(file))
	err = json.Unmarshal(file, config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	Token = config.Token
	Prefix = config.Prefix

	return nil
}
