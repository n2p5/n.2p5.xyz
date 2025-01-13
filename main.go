package main

import (
	"context"
	"os"
)

func main() {
	err := GenerateHome("dst/index.html")
	if err != nil {
		panic(err)
	}
	err = GenerateMedia("dst/writing-speaking-and-press.html")
	if err != nil {
		panic(err)
	}
}

func GenerateHome(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	return Home().Render(context.TODO(), f)
}

func GenerateMedia(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	return Media().Render(context.TODO(), f)
}
