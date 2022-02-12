package main

import (
	"log"

	obsws "github.com/christopher-dG/go-obs-websocket"
)

func main() {
	// Connect a client.
	c := obsws.Client{Host: "localhost", Port: 4444}
	if err := c.Connect(); err != nil {
		log.Fatal(err)
	}
	defer c.Disconnect()

	// Send and receive a request asynchronously.
	req := obsws.NewGetSceneListRequest()
	if err := req.Send(c); err != nil {
		log.Fatal(err)
	}
	// This will block until the response comes (potentially forever).
	resp, err := req.Receive()
	if err != nil {
		log.Fatal(err)
	}

	scenes := resp.Scenes
	for _, scene := range scenes {
		print(scene.Name, ": ")
		for _, source := range scene.Sources {
			print(source.Name, "; ")
		}
	}
}
