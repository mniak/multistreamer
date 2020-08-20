package main

import "os"

func init() {
	STREAMING_KEY = os.Getenv("STREAMING_KEY")
}
