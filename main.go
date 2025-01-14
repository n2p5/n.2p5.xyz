package main

import (
	"context"
	"os"
)

func main() {
	err := GenerateHome("data/home.toml", "dst/index.html")
	if err != nil {
		panic(err)
	}
	err = GenerateMedia("dst/writing-speaking-and-press.html")
	if err != nil {
		panic(err)
	}
}

func GenerateHome(input, output string) error {
	f, err := os.Create(output)
	if err != nil {
		return err
	}

	d := HomeData{
		Title: "Nathan Toups",
		Description: `My name is Nathan Toups. I'm not convinced that I could (or should) try
      to sum myself up in a few words, but I can certainly share a bit of
      propaganda:`,
		Items: []string{
			`I co-host a podcast called 
        <a href="https://bookoverflow.io/">Book Overflow</a>.`,
		},
	}

	return Home(d).Render(context.TODO(), f)
}

func GenerateMedia(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	d := MediaData{
		Title:       "Nathan Toups - Writing, Speaking, and Press",
		Description: `Selected Writing, Speaking, and Press`,
		Items: []MediaItem{
			{
				Title:   "Nathan Toups - Selected Writing, Speaking, and Press",
				Details: `Blah`,
			},
		},
	}

	return Media(d).Render(context.TODO(), f)
}
