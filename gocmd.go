package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("/Users/reason/PycharmProjects/common-tool/bin/ffmpeg", "-version")
	buf, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(buf))
}
