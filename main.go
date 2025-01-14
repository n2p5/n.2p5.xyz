package main

import (
	"context"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

func main() {
	err := GenerateHome("data/home.toml", "dst/index.html")
	if err != nil {
		log.Fatal(err)
	}
	err = GenerateMedia("dst/writing-speaking-and-press.html")
	if err != nil {
		log.Fatal(err)
	}
}

func GenerateHome(input, output string) error {
	f, err := os.Create(output)
	if err != nil {
		return err
	}

	var d HomeData
	toml.DecodeFile(input, &d)
	return Home(d).Render(context.TODO(), f)
}

func GenerateMedia(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	var d MediaData
	toml.DecodeFile("data/media.toml", &d)
	return Media(d).Render(context.TODO(), f)
}
