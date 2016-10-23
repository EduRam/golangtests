package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func main() {
	log.Println("Start")

	var test_path = "/dev/bus1"

	//Clang
	// r = stat(test_path, &st);
	// assert((st.st_mode & S_IFMT) == S_IFCHR);

	fileInfo, err := os.Stat(test_path)
	if err != nil {
		log.Fatal(err)
	}

	if (fileInfo.Mode() & os.ModeCharDevice) != os.ModeCharDevice {
		log.Fatal("Is not character device")
	}

	//Clang
	// r = open(test_path, O_RDWR | O_CLOEXEC | O_NONBLOCK | O_NOCTTY);
	file, err := os.OpenFile(test_path, syscall.O_RDWR|syscall.O_CLOEXEC|syscall.O_NONBLOCK|syscall.O_NOCTTY, os.ModePerm)

	if err != nil {
		log.Fatal(err)
	}

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(countOpenFiles())

	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	log.Println("End test")
}

func countOpenFiles() int {
	out, err := exec.Command("/bin/sh", "-c", fmt.Sprintf("lsof -p %v", os.Getpid())).Output()
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(out), "\n")
	return len(lines) - 1
}
