package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

type configuration struct {
	configurationRoot *configurationRoot `json:"configuration"`
}

type configurationRoot struct {
	Credentials apiCredentialsList `json:"credentials"`
}

type apiCredentialsList []apiCredentials

type apiCredentials struct {
	Key    string `json:"key"`
	Url    string `json:"url"`
	Id     string `json:"client_id"`
	Secret string `json:"client_secret"`
}

func NewConfiguration(clientId string, clientSecret string) *configuration {
	var credentials = make(apiCredentialsList, 1)

	credentials[0] = apiCredentials{
		Key:    "default", // These credentials are the ones used by default
		Url:    "default", // These credentials are for the production API
		Id:     clientId,
		Secret: clientSecret,
	}

	config := &configurationRoot{
		Credentials: credentials,
	}

	configuration := &configuration{
		configurationRoot: config,
	}

	return configuration
}

func (config *configuration) Save() error {
	json, _ := json.Marshal(config)
	fmt.Println(string(json))

	CreateCH360DirIfNotExists()

	filename := filepath.Join(GetCH360Dir(), "config.json")

	err := ioutil.WriteFile(filename, json, 0644) //TODO: Permissions?
	return err
}

func CreateCH360DirIfNotExists() {
	dir := GetCH360Dir()

	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			// file does not exist
			os.Mkdir(dir, 0644) //TODO: Permissions?
		} else {
			// other error
			//TODO: Return error
		}
	}
}

func GetCH360Dir() string {
	return filepath.Join(UserHomeDir(), ".ch360")
}

func UserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}
