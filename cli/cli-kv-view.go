package cli

import (
	"fmt"
	"strconv"
	"wk-local-pwd-manager/model"
	"wk-local-pwd-manager/persistent"
	"wk-local-pwd-manager/utils"
)

func runKvMainView(secretKey string) {
	for {
		fmt.Println("Loading KV file...")
		secretKVs, err := persistent.GetAllSecretKVs()
		if err != nil {
			fmt.Println("Error loading secret kvs: " + err.Error())
			return
		}
		fmt.Println("================================================")
		fmt.Println("Vault")
		fmt.Println("================================================")
		fmt.Println("------------------------------------------------")
		if len(secretKVs) == 0 {
			fmt.Println("---Vault is empty---")
		}
		for i, kv := range secretKVs {
			fmt.Printf("%d:\nKey: %s\nValue: %s\nDescription: %s\n\n", i, kv.Key, "********", kv.Description)
		}
		fmt.Println("------------------------------------------------")
		fmt.Println("Select your action:")
		fmt.Println("0. Reveal a KV")
		fmt.Println("1. Add a new KV")
		fmt.Println("2. Edit a KV")
		fmt.Println("3. Delete a KV")
		fmt.Println("4. Exit")
		action := utils.GetStdIn()
		switch action {
		case "0":
			runRevealKvView(secretKVs, secretKey)
		case "1":
			runAddNewKvView(secretKVs, secretKey)
		case "2":
			runEditKvView(secretKVs, secretKey)
		case "3":
			runDeleteKvView(secretKVs, secretKey)
		case "4":
			return
		default:
			fmt.Println("Invalid action, please try again")
		}
	}
}

func runRevealKvView(secretKVs []model.SecretKV, secretKey string) {
	fmt.Println("Please enter the index of the KV to reveal:")
	index, err := strconv.Atoi(utils.GetStdIn())
	if err != nil {
		fmt.Println("Error parsing index: " + err.Error())
		return
	}
	if index < 0 || index >= len(secretKVs) {
		fmt.Println("Invalid index, please try again")
		return
	}
	secretKV := secretKVs[index]
	unmaskedValue, err := utils.Decrypt(secretKV.Value, secretKey)
	if err != nil {
		fmt.Println("Error decrypting value: " + err.Error())
		return
	}
	fmt.Printf("Key: %s\n", secretKV.Key)
	fmt.Printf("Value: %s\n", unmaskedValue)
	fmt.Printf("Description: %s\n", secretKV.Description)
	utils.Prompt("Press any key to continue...")
}

func runAddNewKvView(secretKVs []model.SecretKV, secretKey string) {
	key := utils.Prompt("Please enter the key:")
	value := utils.Prompt("Please enter the value:")
	maskedValue, err := utils.Encrypt(value, secretKey)
	if err != nil {
		fmt.Println("Error encrypting value: " + err.Error())
		return
	}
	description := utils.Prompt("Please enter the description:")
	secretKV := model.SecretKV{
		Key:         key,
		Value:       maskedValue,
		Description: description,
	}
	secretKVs = append(secretKVs, secretKV)
	err = persistent.WriteAllSecretKVs(secretKVs)
	if err != nil {
		fmt.Println("Error writing secret kvs: " + err.Error())
		return
	}
}

func runEditKvView(secretKVs []model.SecretKV, secretKey string) {
	fmt.Println("Please enter the index of the KV to edit:")
	index, err := strconv.Atoi(utils.GetStdIn())
	if err != nil {
		fmt.Println("Error parsing index: " + err.Error())
		return
	}
	if index < 0 || index >= len(secretKVs) {
		fmt.Println("Invalid index, please try again")
		return
	}
	secretKV := secretKVs[index]
	unmaskedValue, err := utils.Decrypt(secretKV.Value, secretKey)
	if err != nil {
		fmt.Println("Error decrypting value: " + err.Error())
		return
	}
	fmt.Printf("Key: %s\n", secretKV.Key)
	fmt.Printf("Value: %s\n", unmaskedValue)
	fmt.Printf("Description: %s\n", secretKV.Description)
	fmt.Println("------------------------------------------------")
	key := utils.Prompt("Enter new Key:")
	value := utils.Prompt("Enter new Value:")
	maskedValue, err := utils.Encrypt(value, secretKey)
	if err != nil {
		fmt.Println("Error encrypting value: " + err.Error())
		return
	}
	description := utils.Prompt("Enter new Description:")
	secretKVs[index].Key = key
	secretKVs[index].Value = maskedValue
	secretKVs[index].Description = description
	err = persistent.WriteAllSecretKVs(secretKVs)
	if err != nil {
		fmt.Println("Error writing secret kvs: " + err.Error())
		return
	}
}

func runDeleteKvView(secretKVs []model.SecretKV, secretKey string) {
	fmt.Println("Please enter the index of the KV to delete:")
	index, err := strconv.Atoi(utils.GetStdIn())
	if err != nil {
		fmt.Println("Error parsing index: " + err.Error())
		return
	}
	if index < 0 || index >= len(secretKVs) {
		fmt.Println("Invalid index, please try again")
		return
	}
	secretKV := secretKVs[index]
	fmt.Printf("Key: %s\n", secretKV.Key)
	fmt.Printf("Description: %s\n", secretKV.Description)
	fmt.Println("------------------------------------------------")
	proceed := utils.Prompt("Are you sure you want to delete this KV? (y/n)")
	if proceed == "y" || proceed == "Y" {
		secretKVs = append(secretKVs[:index], secretKVs[index+1:]...)
		err = persistent.WriteAllSecretKVs(secretKVs)
		if err != nil {
			fmt.Println("Error writing secret kvs: " + err.Error())
			return
		}
	}

}
