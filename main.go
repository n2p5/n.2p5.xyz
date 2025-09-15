package main

import (
	"context"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/a-h/templ"
)

func main() {
	g := []func() error{
		generator("data/home.toml", "dst/index.html", HomeData{}, Home),
		generator("data/media.toml", "dst/writing-speaking-and-press.html", MediaData{}, Media),
	}

	for _, gen := range g {
		if err := gen(); err != nil {
			log.Fatal(err)
		}
	}
}

// generator is a generic function that takes an input TOML file, an output file path,
// a data structure to decode the TOML into, and a templ.Component function that
// generates the HTML. It returns a function that performs the generation when called.
func generator[T any](input, output string, data T, component func(T) templ.Component) func() error {
	return func() error {
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
}
