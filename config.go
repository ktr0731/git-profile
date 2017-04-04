package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
)

const (
	profileFile = "profiles"
)

var configPath string
var profiles *Profiles

type Profile struct {
	Title string
	Name  string
	Email string

	using   bool
	deleted bool
}

type Profiles []Profile

func (p *Profiles) Add(profile *Profile) error {
	for _, prof := range *p {
		if prof.Title == profile.Title {
			return errors.New("duplicated title, please use unique title")
		}
	}
	*p = append(*p, *profile)

	return nil
}

func (p Profiles) Get(title string) (Profile, error) {
	for _, prof := range p {
		if prof.Title == title {
			return prof, nil
		}
	}

	return Profile{}, errors.New("profile not found")
}

func (p Profiles) Remove(title string) error {
	for i, profile := range p {
		if profile.Title == title {
			p[i].deleted = true
			return nil
		}
	}

	return errors.New("passed profile not found")
}

func (p Profiles) Save() error {
	tmp := make(Profiles, 0, len(p))
	for _, profile := range p {
		if !profile.deleted {
			tmp = append(tmp, profile)
		}
	}

	b, err := json.MarshalIndent(&tmp, "", "  ")
	if err != nil {
		return errors.Wrap(err, "failed to marshal profiles JSON")
	}

	if err := ioutil.WriteFile(filepath.Join(configPath, profileFile), b, 0644); err != nil {
		return errors.Wrap(err, "cannot write profiles JSON to file")
	}

	return nil
}

func init() {
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	configPath = filepath.Join(home, ".config", "git-config")

	os.MkdirAll(configPath, 0755)
	f, err := os.OpenFile(filepath.Join(configPath, profileFile), os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	if len(b) == 0 {
		profiles = &Profiles{}
		return
	}

	err = json.Unmarshal(b, &profiles)
	if err != nil {
		panic(err)
	}
}
