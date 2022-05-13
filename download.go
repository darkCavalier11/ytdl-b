package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("yt-dlp", "-f 251", "https://www.youtube.com/watch?v=Nq3x1AkwgpY")
	stdout, _ := cmd.Output()
	fmt.Println(stdout)
}
