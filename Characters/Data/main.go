package main

import (
	b "bufio"
	j "encoding/json"
	f "fmt"
	o "os"
)
type Characters struct {
	Char Character `json:"Characters"`
}

type Information struct {
	Age int `json:"age"`
	Role string `json:"role"`
	Faction string `json:"faction"`
}

type Character struct {
	Name string `json:"name"`
	Info Information `json:"info"`
}


func main(){
	var (
		name, role, faction string
		age int
		char Character
	)
	in := b.NewReader(o.Stdin)
	f.Println("Enter Character Name: ")
	name, _ = in.ReadString('\n')
	f.Println("Enter Character Age: ")
	_, _ = f.Scan(&age)
	f.Println("Enter Character Role: ")
	role, _ = in.ReadString('\n')
	f.Println("Enter Character Faction: ")
	faction, _ = in.ReadString('\n')

	info := Information{
		Age:     age,
		Role:    role,
		Faction: faction,
	}
	char = Character{
		Name: name,
		Info: info,
	}
	Chars := Characters{Char: char}
	// Open our jsonFile
	jsonFile, err := o.Open("character.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		f.Println(err)
	}
	f.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	enc, _ := j.MarshalIndent(Chars, "", "")
	jsonFile.
}

