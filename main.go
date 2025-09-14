package main

import (
	"context"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/a-h/templ"
)

func main() {
	err := Generate("data/home.toml", "dst/index.html", HomeData{}, Home)
	if err != nil {
		log.Fatal(err)
	}
	err = Generate("data/media.toml", "dst/writing-speaking-and-press.html", MediaData{}, Media)
	if err != nil {
		log.Fatal(err)
	}
}

func Generate[T any](input, output string, data T, component func(T) templ.Component) error {
	f, err := os.Create(output)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = toml.DecodeFile(input, &data)
	if err != nil {
		return err
	}
	return component(data).Render(context.Background(), f)
}
