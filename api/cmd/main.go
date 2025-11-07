package main

import (
	"summarizer/internal"
)

func main() {
	app := internal.App()
	app.Listen(":3000")
}
