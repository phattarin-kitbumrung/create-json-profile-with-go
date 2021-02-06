package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Minizymint struct {
	Profile []Profile `json:"minizymint"`
}

type Profile struct {
	Name              string            `json:"name"`
	Nickname          string            `json:"nickname"`
	Birthday          string            `json:"birthday"`
	Religion          string            `json:"religion"`
	Address           string            `json:"address"`
	Email             string            `json:"email"`
	Mobile            string            `json:"mobile"`
	Github            string            `json:"github"`
	Education         Education         `json:"education"`
	Skills            Skills            `json:"skills"`
	Knowledge         string            `json:"knowledge & technology"`
	Interested        string            `json:"interested"`
	WorkedExperiences WorkedExperiences `json:"worked experiences"`
}

type Education struct {
	BachelorDegree string `json:"bachelor degree"`
	MasterDegree   string `json:"master degree"`
	University     string `json:"university"`
}

type Skills struct {
	OperatingSystems    string `json:"operating systems"`
	ProgrammingLanguage string `json:"programming language"`
	Database            string `json:"database"`
	Tools               string `json:"tools"`
}

type WorkedExperiences struct {
	CooperativeEducation string `json:"cooperative education"`
	WebDeveloper         string `json:"web developer"`
	BackendDeveloper     string `json:"backend developer"`
	Programmer           string `json:"programmer"`
}

func main() {
	fmt.Println("***** Welcome to Phattarin Kitbumrung Profile *****")
	jsonFile, err := os.Open("profile.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var minizymint Minizymint
	json.Unmarshal(byteValue, &minizymint)

	for i := 0; i < len(minizymint.Profile); i++ {
		fmt.Println("Name: " + minizymint.Profile[i].Name)
		fmt.Println("Nickname: " + minizymint.Profile[i].Nickname)
		fmt.Println("Birthday: " + minizymint.Profile[i].Birthday)
		fmt.Println("Religion: " + minizymint.Profile[i].Religion)
		fmt.Println("Address: " + minizymint.Profile[i].Address)
		fmt.Println("Email: " + minizymint.Profile[i].Email)
		fmt.Println("Mobile: " + minizymint.Profile[i].Mobile)
		fmt.Println("Github: " + minizymint.Profile[i].Github)
		fmt.Println("Education >>>")
		fmt.Println("Bachelor Degree: " + minizymint.Profile[i].Education.BachelorDegree)
		fmt.Println("Master Degree: " + minizymint.Profile[i].Education.MasterDegree)
		fmt.Println("University: " + minizymint.Profile[i].Education.University)
		fmt.Println("Skills >>>")
		fmt.Println("Operating Systems: " + minizymint.Profile[i].Skills.OperatingSystems)
		fmt.Println("Programming Language: " + minizymint.Profile[i].Skills.ProgrammingLanguage)
		fmt.Println("Database: " + minizymint.Profile[i].Skills.Database)
		fmt.Println("Tools: " + minizymint.Profile[i].Skills.Tools)
		fmt.Println("Knowledge & Technology: " + minizymint.Profile[i].Knowledge)
		fmt.Println("Interested: " + minizymint.Profile[i].Interested)
		fmt.Println("Worked Experiences >>>")
		fmt.Println("Cooperative Education: " + minizymint.Profile[i].WorkedExperiences.CooperativeEducation)
		fmt.Println("Web Developer: " + minizymint.Profile[i].WorkedExperiences.WebDeveloper)
		fmt.Println("Programmer: " + minizymint.Profile[i].WorkedExperiences.Programmer)
		fmt.Println("Backend Developer: " + minizymint.Profile[i].WorkedExperiences.BackendDeveloper)
	}
}
