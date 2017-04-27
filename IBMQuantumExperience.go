package qiskitapigo

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

const URL = "https://quantumexperience.ng.bluemix.net/api"

func IBMQuantumExperience(token string) Api {
	creds := Credentials{}
	creds.ApiToken = token

	p := map[string]string{"apiToken": token}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(p)
	resp, err := http.Post(URL+"/users/loginWithToken", "application/json", b)
	if err != nil {
		log.Fatal(err)
	}

	json.NewDecoder(resp.Body).Decode(&creds)
	return Api{creds}
}

func (api *Api) Get_Code() int {
	return 1
}

func (api *Api) Get_Last_Codes() interface{} {
	lastestCodes := LastestCodes{}

	resp, err := http.Get(URL + "/users/" + api.Creds.UserID + "/codes/lastest?access_token=" + api.Creds.ID)
	if err != nil {
		log.Fatal(err)
	}

	json.NewDecoder(resp.Body).Decode(&lastestCodes)
	return lastestCodes
}

func (api *Api) Get_Execution() interface{} {
	return 3
}

func (api *Api) Get_Result_From_Execution() interface{} {
	return 4
}

func (api *Api) Run_Experiment() interface{} {
	return 5
}

func (api *Api) Run_Jobes() interface{} {
	return 6
}

func (api *Api) Get_Job() interface{} {
	return 7
}
