package graphiste

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func Menu() {
	fmt.Println("▓                                                                                  ▓")
	fmt.Println("▓                                                                                  ▓")
	fmt.Println("▓                                                                                  ▓")
	fmt.Println("▓                       ████████    ██████████    ██████████                       ▓")
	fmt.Println("▓                       ██          ██      ██    ██      ██                       ▓")
	fmt.Println("▓                       ██          ██      ██    ██      ██                       ▓")
	fmt.Println("▓                       ████████    ██████████    ██      ██                       ▓")
	fmt.Println("▓                             ██    ██      ██    ██      ██                       ▓")
	fmt.Println("▓                             ██    ██      ██    ██      ██                       ▓")
	fmt.Println("▓                       ████████    ██      ██    ██████████                       ▓")
	fmt.Println("▓                                                                                  ▓")
	fmt.Println("▓                               ~~SWORD ART ONLINE~~                               ▓")
	fmt.Println("▓                                                                                  ▓")
	fmt.Println("▓                                                                                  ▓")
	fmt.Println("▓                                                                                  ▓")
	time.Sleep(3 * time.Second)
	Clear()
	fmt.Println("                                                                                  ")
	fmt.Println("            █    █  █   █  █  █     ████  █████  █████  ███   █████")
	fmt.Println("            █       ██  █  █ █      █       █    █   █  █  █    █")
	fmt.Println("            █    █  █ █ █  ██       ████    █    █████  ████    █")
	fmt.Println("            █    █  █  ██  █ █         █    █    █   █  █   █   █")
	fmt.Println("            ███  █  █   █  █  █     ████    █    █   █  █    █  █", "\n")
	for i := 1; i < 5; i++ {
		fmt.Println("CHARGEMENT : ", "\n")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔")
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("⚔")
		Clear()
	}
}

func Clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
