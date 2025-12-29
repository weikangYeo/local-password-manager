package cli

import (
	"bufio"
	"fmt"
	"local-pwd-manager/model"
	"local-pwd-manager/persistent"
	"os"
)

func loadProfile() (model.Profile, error) {
	println("Loading default profile...")
	profile, err := persistent.GetProfile()
	if err != nil {
		println("Error loading profile: " + err.Error())
		return model.Profile{}, err
	}
	if profile.Name == "" {
		println("No profile found, creating new profile...")
		profile, err = runCreateNewProfileView()
		if err != nil {
			println("Error creating profile: " + err.Error())
			return model.Profile{}, err
		}
	}
	println("Profile loaded successfully")
	println("------------------------------------------------")
	fmt.Printf("Profile Name: %s\n", profile.Name)
	println("------------------------------------------------")
	return model.Profile{}, err
}

func runCreateNewProfileView() (model.Profile, error) {
	println("================================================")
	println("Create a new Profile")
	println("================================================")
	println("Please enter the profile name:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	profileName := scanner.Text()
	println("Please enter the secret key hints:")
	scanner.Scan()
	profileHints := scanner.Text()

	println("You have entered the following profile:")
	println("Profile Name: " + profileName)
	println("Secret Key Hints: " + profileHints)
	println("Do you want to proceed? (y/n)")
	scanner.Scan()
	proceed := scanner.Text()
	if proceed == "y" || proceed == "Y" || proceed == "yes" || proceed == "Yes" {
		println("Proceeding with profile creation...")
		profile := model.Profile{
			Name:  profileName,
			Hints: profileHints,
		}
		err := persistent.CreateProfile(profile)
		if err != nil {
			return model.Profile{}, err
		}
		println("Profile created successfully")
		return profile, nil
	}
	return runCreateNewProfileView()
}
