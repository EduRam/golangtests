package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/vaughan0/go-ini"
)

var (
	isToRunOnStartup bool
	initCmds         InitCmdsStruct
)

type InitCmdsStruct struct {
	cmdsMap    map[string]string
	sortedKeys []string
	delay      int
}

func main() {
	err := start()
	fmt.Fprintln(os.Stderr, err)
	os.Exit(0)
}

func start() error {

	var configFile = flag.String("configfile", "sysexec.ini", "Name of '.ini' file with commands to execute.")
	flag.Parse()
	file, err := ini.LoadFile(*configFile)
	if err != nil {
		return err
	}

	var value, ok = file.Get("general", "run")
	if !ok {
		return errors.New("Section [general] or Key 'run=' not found. Exit!")
	}

	isToRunOnStartup, err = strconv.ParseBool(value)
	fmt.Println("isToRunOnStartup=", isToRunOnStartup)

	if err != nil {
		return err
	}

	initCmds.cmdsMap = make(map[string]string)
	for key, value := range file["init"] {

		fmt.Printf("[init] %s => %s\n", key, value)
		if key == "delay" {
			i, err := strconv.Atoi(value)
			if err != nil {
				return err
			}
			initCmds.delay = i
		} else if strings.HasPrefix(key, "cmd") {
			initCmds.cmdsMap[key] = value
		}
	}

	for k := range initCmds.cmdsMap {
		initCmds.sortedKeys = append(initCmds.sortedKeys, k)
	}
	sort.Strings(initCmds.sortedKeys)

	err = run(&initCmds)
	fmt.Println(err)

	/*
		fmt.Println("Press \n" +
			"(r) to run\n" +
			"(x) to exit")

		var cmd int
		fmt.Scanf("%c", &cmd)

		if err != nil {
			return err
		}

		switch cmd {
		case 'x':
			return nil
		case 'r':
			return run(&initCmds)

		}
	*/
	return nil
}

func run(initCmds *InitCmdsStruct) error {

	fmt.Println("run")

	for key, value := range initCmds.cmdsMap {

		fmt.Printf("key %s value %s\n", key, value)
		//cmd := exec.Command("C:\\Windows\\SysWOW64\\cmd.exe", "/c", "dir", ">>", "d:\\lixo.txt")
		cmd := exec.Command("C:\\Windows\\SysWOW64\\cmd.exe", "/c", value)
		output, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("%v\n%s", err, output)
		}
		printOutput(output)
		fmt.Println("sleep")
		time.Sleep(time.Duration(initCmds.delay) * time.Second)

	}
	return nil

}

func printOutput(outs []byte) {
	if len(outs) > 0 {
		fmt.Printf("==> Output: %s\n", string(outs))
	}
}
