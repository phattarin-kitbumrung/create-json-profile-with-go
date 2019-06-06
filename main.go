package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)


type Minizymint struct {
	Minizymint []minizymint `json:"minizymint"`
}

type minizymint struct {
	Name   string `json:"name"`
	Nickname   string `json:"nickname"`
	Birthday string `json:"birthday"`
	Religion string `json:"religion"`
	Address string `json:"address"`
	Email string `json:"email"`
	Mobile string `json:"mobile"`
	Github string `json:"github"`
	Education Education `json:"education"`
	Skills Skills `json:"skills"`
	Knowledge string `json:"knowledge & technology"`
	Interested string `json:"interested"`
	WorkedExperiences WorkedExperiences `json:"worked experiences"`
}

type Education struct {
	BachelorDegree string `json:"bachelor degree"`
	MasterDegree string `json:"master degree"`
	University string `json:"university"`
}

type Skills struct {
	OperatingSystems string `json:"operating systems"`
	ProgrammingLanguage string `json:"programming language"`
	Database string `json:"database"`
	Tools string `json:"tools"`
}

type WorkedExperiences struct {
	CooperativeEducation string `json:"cooperative education"`
	WebDeveloper string `json:"web developer"`
	BackendDeveloper string `json:"backend developer"`
	Programmer string `json:"programmer"`
}


func main() {
	fmt.Println("***** Welcome to Phattarin Kitbumrung Profile *****")

	// Open our jsonFile
	jsonFile, err := os.Open("profile.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}


	//fmt.Println("Successfully Opened profile.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()


	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var minizymint Minizymint
	json.Unmarshal(byteValue, &minizymint)

	for i := 0; i < len(minizymint.Minizymint); i++ {
		fmt.Println("Name: " + minizymint.Minizymint[i].Name)
		fmt.Println("Nickname: " + minizymint.Minizymint[i].Nickname)
		fmt.Println("Birthday: " + minizymint.Minizymint[i].Birthday)
		fmt.Println("Religion: " + minizymint.Minizymint[i].Religion)
		fmt.Println("Address: " + minizymint.Minizymint[i].Address)
		fmt.Println("Email: " + minizymint.Minizymint[i].Email)
		fmt.Println("Mobile: " + minizymint.Minizymint[i].Mobile)
		fmt.Println("Github: " + minizymint.Minizymint[i].Github)
		fmt.Println("Education >>>")
		fmt.Println("Bachelor Degree: " + minizymint.Minizymint[i].Education.BachelorDegree)
		fmt.Println("Master Degree: " + minizymint.Minizymint[i].Education.MasterDegree)
		fmt.Println("University: " + minizymint.Minizymint[i].Education.University)
		fmt.Println("Skills >>>")
		fmt.Println("Operating Systems: " + minizymint.Minizymint[i].Skills.OperatingSystems)
		fmt.Println("Programming Language: " + minizymint.Minizymint[i].Skills.ProgrammingLanguage)
		fmt.Println("Database: " + minizymint.Minizymint[i].Skills.Database)
		fmt.Println("Tools: " + minizymint.Minizymint[i].Skills.Tools)
		fmt.Println("Knowledge & Technology: " + minizymint.Minizymint[i].Knowledge)
		fmt.Println("Interested: " + minizymint.Minizymint[i].Interested)
		fmt.Println("Worked Experiences >>>")
		fmt.Println("Cooperative Education: " + minizymint.Minizymint[i].WorkedExperiences.CooperativeEducation)
		fmt.Println("Web Developer: " + minizymint.Minizymint[i].WorkedExperiences.WebDeveloper)
		fmt.Println("Backend Developer: " + minizymint.Minizymint[i].WorkedExperiences.BackendDeveloper)
		fmt.Println("Programmer: " + minizymint.Minizymint[i].WorkedExperiences.Programmer)
	}
}