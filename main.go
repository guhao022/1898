package main

import "1898/utils/env"

func main() {
	env.Load()
	HttpRun("9900")
}
