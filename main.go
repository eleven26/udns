package main

import "log"

func main() {
	client, err := NewFromDefaultPaths()
	if err != nil {
		log.Fatal(err)
	}

	err = client.Replace()
	if err != nil {
		log.Fatal(err)
	}
}
