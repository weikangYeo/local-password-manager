package cli

import (
	"bufio"
	"fmt"
	"local-pwd-manager/persistent"
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
	println("Enter secret key to")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	secretKey := scanner.Text()
	println("Loading KV file...")
	secretKVs, err := persistent.GetAllSecretKVs()
	if err != nil {
		println("Error loading secret kvs: " + err.Error())
		return
	}
	println("KV file loaded successfully")
	println("------------------------------------------------")
	if len(secretKVs) == 0 {
		println("---Vault is empty---")
	}
	for i, kv := range secretKVs {
		fmt.Printf("%d:\n Key: %s\n Value: %s\n Description: %s\n\n", i, kv.Key, "********", kv.Description)
	}
	println("------------------------------------------------")
	println("Select your action:")
	println("0. Reveal a KV")
	println("1. Add a new KV")
	println("2. Edit a KV")
	println("3. Delete a KV")
	println("4. Exit")
	scanner.Scan()
	action := scanner.Text()
	switch action {
	case "1":
		runAddNewKVView()
	case "2":
		runEditKVView()
	case "3":
		runDeleteKVView()
	case "4":
		return
	}

}
