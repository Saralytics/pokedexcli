package main

type Config struct {
	Next     *string
	Previous *string
}

func main() {

	cfg := &Config{}
	startRepl(cfg)

}
