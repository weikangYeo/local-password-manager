package persistent

import (
	"encoding/json"
	"wk-local-pwd-manager/model"
	"wk-local-pwd-manager/utils"
)

func GetAllSecretKVs() ([]model.SecretKV, error) {
	if !utils.IsFileExists("kv.json") {
		return []model.SecretKV{}, nil
	}
	config, err := utils.ReadFile("kv.json")
	if err != nil {
		return []model.SecretKV{}, err
	}
	var secretKVs []model.SecretKV
	err = json.Unmarshal(config, &secretKVs)
	if err != nil {
		return []model.SecretKV{}, err
	}
	return secretKVs, nil
}

// for simplicity sake (scope), always write all to file whenver there is a change.
func WriteAllSecretKVs(secretKVs []model.SecretKV) error {
	data, err := json.Marshal(secretKVs)
	if err != nil {
		return err
	}
	err = utils.WriteFile("kv.json", data)
	if err != nil {
		return err
	}
	return nil
}
