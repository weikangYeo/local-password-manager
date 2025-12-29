package cli

import (
	"bufio"
	"fmt"
	"local-pwd-manager/model"
	"local-pwd-manager/persistent"
	"local-pwd-manager/utils"
	"os"
	"strconv"
)

func runKvMainView(secretKey string) {
	var action string
	for action != "4" {
		println("Loading KV file...")
		secretKVs, err := persistent.GetAllSecretKVs()
		if err != nil {
			println("Error loading secret kvs: " + err.Error())
			return
		}
		println("================================================")
		println("Vault Menu:")
		println("================================================")
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
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		action := scanner.Text()
		switch action {
		case "0":
			runRevealKVView(secretKVs, secretKey)
		case "1":
			runAddNewKVView(secretKVs, secretKey)
		case "2":
			runEditKVView()
		case "3":
			runDeleteKVView()
		case "4":
			return
		default:
			println("Invalid action, please try again")
		}
	}
}

func runRevealKVView(secretKVs []model.SecretKV, secretKey string) {
	println("Please enter the index of the KV to reveal:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	index, err := strconv.Atoi(scanner.Text())
	if err != nil {
		println("Invalid index, please enter a number only")
		return
	}
	if index < 0 || index >= len(secretKVs) {
		println("Invalid index, please try again")
		return
	}
	secretKV := secretKVs[index]
	unmaskedValue, err := utils.Decrypt(secretKV.Value, secretKey)
	if err != nil {
		println("Error decrypting value: " + err.Error())
		return
	}
	fmt.Printf("Key: %s\n", secretKV.Key)
	fmt.Printf("Value: %s\n", unmaskedValue)
	fmt.Printf("Description: %s\n", secretKV.Description)
	println("\n\n\n")
}

func runAddNewKVView(secretKVs []model.SecretKV, secretKey string) {
	println("Please enter the key:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	key := scanner.Text()
	println("Please enter the value:")
	scanner.Scan()
	value := scanner.Text()
	maskedValue, err := utils.Encrypt(value, secretKey)
	if err != nil {
		println("Error encrypting value: " + err.Error())
		return
	}
	println("Please enter the description:")
	scanner.Scan()
	description := scanner.Text()
	secretKV := model.SecretKV{
		Key:         key,
		Value:       maskedValue,
		Description: description,
	}
	secretKVs = append(secretKVs, secretKV)
	persistent.WriteAllSecretKVs(secretKVs)
}

func runEditKVView(secretKVs []model.SecretKV, secretKey string) {
}

func runDeleteKVView(secretKVs []model.SecretKV, secretKey string) {
}
