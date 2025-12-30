package cli

import (
	"fmt"
	"wk-local-pwd-manager/model"
	"wk-local-pwd-manager/persistent"
	"wk-local-pwd-manager/utils"
)

func loadProfile() (model.Profile, error) {
	fmt.Println("Loading default profile...")
	profile, err := persistent.GetProfile()
	if err != nil {
		fmt.Println("Error loading profile: " + err.Error())
		return model.Profile{}, err
	}
	if profile.Name == "" {
		fmt.Println("No profile found, creating new profile...")
		profile, err = runCreateNewProfileView()
		if err != nil {
			fmt.Println("Error creating profile: " + err.Error())
			return model.Profile{}, err
		}
	}
	return profile, err
}

func runCreateNewProfileView() (model.Profile, error) {
	fmt.Println("================================================")
	fmt.Println("Create a new Profile")
	fmt.Println("================================================")
	profileName := utils.Prompt("Please enter the profile name:")
	profileHints := utils.Prompt("Please enter the secret key hints:")
	fmt.Println("You have entered the following profile:")
	fmt.Println("Profile Name: " + profileName)
	fmt.Println("Secret Key Hints: " + profileHints)
	proceed := utils.Prompt("Do you want to proceed? (y/n)")
	if proceed == "y" || proceed == "Y" || proceed == "yes" || proceed == "Yes" {
		fmt.Println("Proceeding with profile creation...")
		profile := model.Profile{
			Name:  profileName,
			Hints: profileHints,
		}
		err := persistent.CreateProfile(profile)
		if err != nil {
			return model.Profile{}, err
		}
		fmt.Println("Profile created successfully")
		return profile, nil
	}
	return runCreateNewProfileView()
}
