package main

import (
	"fmt"
	"time"

	"github.com/argadhana/learn-go-mongo/config"
	"github.com/argadhana/learn-go-mongo/src/modules/profile/model"
	"github.com/argadhana/learn-go-mongo/src/modules/profile/repository"
)

func main() {
	db, err := config.GetMongoDB()

	if err != nil {
		fmt.Println(err)
	}
	profilesRepository := repository.NewProfileRepositoryMongo(db, "profiles")

	// saveProfile(profilesRepository)
	// updateProfile(profilesRepository)
	// deleteProfile(profilesRepository)
	// getProfile("U2", profilesRepository)
	getProfiles(profilesRepository)
}

func saveProfile(profileRepository repository.ProfileRepository) {
	var p model.Profile

	p.ID = "U2"
	p.FirstName = "nayya"
	p.LastName = "kamila"
	p.Email = "nayyakamila@arga.com"
	p.Password = "Bismillah!"
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	err := profileRepository.Save(&p)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Profile Saved")
	}

}

func updateProfile(profileRepository repository.ProfileRepository) {
	var p model.Profile

	p.ID = "U1"
	p.FirstName = "arga"
	p.LastName = "dhana"
	p.Email = "arga@arga.com"
	p.Password = "Wisanggeni"
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	err := profileRepository.Update("U1", &p)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Profile Updated")
	}

}

func deleteProfile(profileRepository repository.ProfileRepository) {
	err := profileRepository.Delete("U1")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Profile deleted")
	}

}

func getProfile(id string, profileRepository repository.ProfileRepository) {
	profile, err := profileRepository.FindByID(id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(profile.ID)
	fmt.Println(profile.FirstName)
	fmt.Println(profile.LastName)
	fmt.Println(profile.Email)

}

func getProfiles(profileRepository repository.ProfileRepository) {
	profiles, err := profileRepository.FindAll()
	if err != nil {
		fmt.Println(err)
	}

	for _, profile := range profiles {
		fmt.Println("===================")
		fmt.Println(profile.ID)
		fmt.Println(profile.FirstName)
		fmt.Println(profile.LastName)
		fmt.Println(profile.Email)
		fmt.Println("===================")

	}
}
