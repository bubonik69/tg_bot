package fileParsing

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)
type ExampleRequest struct{
	Examples []string `json : examples`
	Responses []string `json : responses`
}
type Intents struct{
	Intents map[string]ExampleRequest `json:"intents"`
}


func JsonToData(path string)(data Intents){
	content, err := ioutil.ReadFile(path)
	CheckErr(err)
	if json.Valid(content){
		err= json.Unmarshal(content,&data)
		CheckErr(err)
		for key,value:= range data.Intents{
			fmt.Println(key,value.Examples)
		}
	}
	return
}

func CheckErr(err error){
	if err!=nil{
		log.Println(err)
		//log.Fatal(err)
		return
	}
}