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

	// cmdsMap does not sort iteration order
	// I want to run with this order: cmd1, cmd2, ...
	// I will sort keys and maitain them on a paralel variable sortedKeys
	for k := range initCmds.cmdsMap {
		initCmds.sortedKeys = append(initCmds.sortedKeys, k)
	}
	sort.Strings(initCmds.sortedKeys)

	if isToRunOnStartup {
		fmt.Println("Running init commands:")
		err = run(&initCmds)
		fmt.Println(err)
	}

	fmt.Println("Press \n" +
		"(r + ENTER) to run\n" +
		"(x + ENTER) to exit")

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
	return nil
}

func run(initCmds *InitCmdsStruct) error {

	for _, key := range initCmds.sortedKeys {

		value := initCmds.cmdsMap[key]
		fmt.Printf("Exec %s = %s\n", key, value)
		//cmd := exec.Command("C:\\Windows\\SysWOW64\\cmd.exe", "/c", "dir", ">>", "d:\\lixo.txt")
		cmd := exec.Command("C:\\Windows\\SysWOW64\\cmd.exe", "/c", value)
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Errorf("%v\n%s", err, output)
		} else {
			printOutput(output)
		}
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
