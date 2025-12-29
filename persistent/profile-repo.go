package persistent

import (
	"encoding/json"
	"local-pwd-manager/model"
	"local-pwd-manager/utils"
)

func CreateProfile(profile model.Profile) error {
	data, err := json.Marshal(profile)
	if err != nil {
		return err
	}
	err = utils.WriteFile("config.json", data)
	if err != nil {
		return err
	}
	return nil
}

// at the time of writing, only one profile is supported.
func GetProfile() (model.Profile, error) {

	if !utils.IsFileExists("config.json") {
		return model.Profile{}, nil
	}
	config, err := utils.ReadFile("config.json")
	if err != nil {
		return model.Profile{}, err
	}
	var profile model.Profile
	err = json.Unmarshal(config, &profile)
	if err != nil {
		return model.Profile{}, err
	}
	return profile, nil
}
