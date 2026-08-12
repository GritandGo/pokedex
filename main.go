package main


func main() {
	cfg := &config{
		commandConfig: getCommands(),
	}

	startRepl(cfg)
}