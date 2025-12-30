package utils

import (
	"bufio"
	"fmt"
	"os"
)

func GetStdIn() string {
	fmt.Print("PwdMgr> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func Prompt(message string) string {
	fmt.Println(message)
	fmt.Print("PwdMgr> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
