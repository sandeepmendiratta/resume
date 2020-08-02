package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

type Skill struct {
	Title        string
	Specialities []string
}

type Education struct {
	YearStart    int
	YearComplete int
	Description  string
}

type References struct {
	Name        string
	Description string
}

type Experience struct {
	Employer         string
	Role             string
	Location         string
	TimeFrame        string
	Responsibilities []string
}

func main() {
	skills := []*Skill{}
	skillBytes, err := ioutil.ReadFile("skills.json")
	if err != nil {
		fmt.Println("Error reading skills.json:", err)
		fmt.Println(string(skillBytes))
		os.Exit(1)
	}

	err = json.Unmarshal(skillBytes, &skills)
	if err != nil {
		fmt.Println("Error unmarshalling skills:", err)
		fmt.Println(string(skillBytes))
		os.Exit(1)
	}

	skillTmpl, err := template.New("skill").Parse(`
Skills
------

{{ range . }}{{ .Title }}
:{{ range .Specialities }} {{ . }}{{ end }}

{{ end }}`)
	if err != nil {
		fmt.Println("Error parsing skills template:", err)
		os.Exit(1)
	}

	edu := []*Education{}
	eduBytes, err := ioutil.ReadFile("education.json")
	if err != nil {
		fmt.Println("Error reading education.json:", err)
		fmt.Println(string(eduBytes))
		os.Exit(1)
	}

	err = json.Unmarshal(eduBytes, &edu)
	if err != nil {
		fmt.Println("Error unmarshalling education:", err)
		fmt.Println(string(eduBytes))
		os.Exit(1)
	}

	eduTmpl, err := template.New("education").Parse(`
Education
------

{{ range . }}{{ .YearStart }}-{{ .YearComplete }}
: {{ .Description }}

{{ end }}`)
	if err != nil {
		fmt.Println("Error parsing education template:", err)
		os.Exit(1)
	}

	// ref := []*References{}
	// refBytes, err := ioutil.ReadFile("references.json")
	// if err != nil {
	// 	fmt.Println("Error reading references.json:", err)
	// 	fmt.Println(string(refBytes))
	// 	os.Exit(1)
	// }
	//
	// err = json.Unmarshal(refBytes, &ref)
	// if err != nil {
	// 	fmt.Println("Error unmarshalling references:", err)
	// 	fmt.Println(string(refBytes))
	// 	os.Exit(1)
	// }

	// 	refTmpl, err := template.New("references").Parse(`
	// References
	// ------
	//
	// {{ range . }}{{ .Name }}
	// : {{ .Description }}
	//
	// {{ end }}`)
	// 	if err != nil {
	// 		fmt.Println("Error parsing references template:", err)
	// 		os.Exit(1)
	// 	}

	responsbility := []*Experience{}
	respBytes, err := ioutil.ReadFile("experience.json")
	if err != nil {
		fmt.Println("Error reading experience.json:", err)
		// fmt.Println(string(refBytes))
		os.Exit(1)
	}

	err = json.Unmarshal(respBytes, &responsbility)
	if err != nil {
		fmt.Println("Error unmarshalling experience:", err)
		fmt.Println(string(respBytes))
		os.Exit(1)
	}

	respTmpl, err := template.New("experience").Parse(`
Experience
------

{{ range . }}{{ .Employer }}
{{ .Role }}
{{ .Location }}
{{ .TimeFrame }}

{{ range .Responsibilities }} * {{ . }}
{{ end }}
{{ end }}`)
	if err != nil {
		fmt.Println("Error parsing experience template:", err)
		os.Exit(1)
	}

	f, err := os.Create("README.md")
	if err != nil {
		fmt.Println("Error opening README.md:", err)
		os.Exit(1)
	}
	defer f.Close()

	err = copyContents("HEADER.md", f)
	skillTmpl.Execute(f, skills)
	eduTmpl.Execute(f, edu)
	// refTmpl.Execute(f, ref)
	respTmpl.Execute(f, responsbility)
}

func copyContents(origin string, target *os.File) error {
	f, err := os.Open(origin)
	if err != nil {
		fmt.Println("Error opening "+origin+":", err)
		os.Exit(1)
	}
	defer f.Close()

	contents, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	_, err = target.Write(contents)
	if err != nil {
		return err
	}

	return nil
}
