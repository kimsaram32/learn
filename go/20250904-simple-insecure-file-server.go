// 20250904
// a simple inefficient file server that leaks basically everything inside your computer
// inspired by exercise 8.2 from The Go Programming Language

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	println("Server started")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	fmt.Printf("Accepted %s\n", conn.RemoteAddr().String())

	input := bufio.NewScanner(conn)
	dir, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(conn, "internal error")
		return
	}

	for input.Scan() {
		args := strings.Split(input.Text(), " ")
		cmd := args[0]

		switch cmd {
		case "cd":
			if len(args) < 2 {
				fmt.Fprintln(conn, "usage: cd [directory]")
				break
			}

			newDir := args[1]
			if newDir[0] != '/' {
				newDir = filepath.Join(dir, newDir)
			}

			stat, err := os.Stat(newDir)
			if err != nil {
				fmt.Fprintln(conn, "failed to open directory")
				break
			}
			if !stat.IsDir() {
				fmt.Fprintln(conn, "not a directory")
				break
			}

			dir = newDir
		case "ls":
			entries, err := os.ReadDir(dir)
			if err != nil {
				fmt.Fprintln(conn, "0")
				break
			}

			for _, entry := range entries {
				fmt.Fprintln(conn, entry.Name())
			}
		case "pwd":
			fmt.Fprintln(conn, dir)
		case "cat":
			if len(args) < 2 {
				fmt.Fprintln(conn, "usage: cat [file]")
				break
			}

			bytes, err := os.ReadFile(filepath.Join(dir, args[1]))
			if err != nil {
				fmt.Fprintln(conn, "failed to read file")
				break
			}
			conn.Write(bytes)
		case "close":
			fmt.Fprintln(conn, "bye")
			return
		}
	}
}
