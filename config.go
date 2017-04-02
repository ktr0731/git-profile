package main

import (
	"encoding/json"
	"errors"
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
var profiles Profiles

type Profile struct {
	Title, Name, Email string
}

type Profiles []Profile

func (p *Profile) Save() error {
	profiles = append(profiles, *p)
	b, err := json.Marshal(profiles)
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

	err = json.Unmarshal(b, &profiles)
	if err != nil {
		panic(err)
	}
}

func NewProfile(title, name, email string) *Profile {
	return &Profile{
		Title: title,
		Name:  name,
		Email: email,
	}
}
