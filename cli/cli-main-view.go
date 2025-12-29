package cli

import (
	"bufio"
	"fmt"
	"os"
)

func RunCLI() {

	println("================================================")
	println("Welcome to Password Manager CLI Mode")
	println("================================================")
	profile, err := loadProfile()
	if err != nil {
		return
	}
	// todo: to verify secret key, use the profile.KeySignature to verify.
	fmt.Printf("Enter secret key to proceed\n, hints:\n%s\n\n> ", profile.Hints)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	secretKey := scanner.Text()
	// the sole purpose of profile is to provide hint to user to remember the secret key, and verify signature.
	runKvMainView(secretKey)
}
