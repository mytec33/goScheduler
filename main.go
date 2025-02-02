package main

func main() {
	var configuration config.Configuration

	configuration.ReadConfiguration("./configuration/config.json")
}
