package cli

import (
	"fmt"
	"wk-local-pwd-manager/utils"
)

func RunCLI() {

	fmt.Println("================================================")
	fmt.Println("Welcome to Password Manager CLI Mode")
	fmt.Println("================================================")
	profile, err := loadProfile()
	if err != nil {
		return
	}
	fmt.Println("Profile loaded successfully")
	fmt.Printf("Welcome %s!\n", profile.Name)

	// todo: to verify secret key, use the profile.KeySignature to verify.
	fmt.Println("Enter secret key to proceed")
	fmt.Println("------------------------------------------------")
	fmt.Printf("Hints: %s\n", profile.Hints)
	fmt.Println("------------------------------------------------")
	secretKey := utils.GetStdIn()
	// the sole purpose of "profile" is to provide hint to user to remember the secret key, and verify signature.
	runKvMainView(secretKey)
}
