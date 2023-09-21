package main

import (
	sao "projet-red/code_game/character"
	graphiste "projet-red/code_game/menu"
)

func main() {
	graphiste.Menu()
	sao.Create_Player()
}
