package qiskitapigo

//import "net/http"

// const CONFIG_BASE = {
//   URL: "https://quantumexperience.ng.bluemix.net/api"
// }

type Api struct{}

type Credentials struct {
	ID string `json:"id"`
	TTL int `json:"ttl"`
	Created time.Time `json:"created"`
	UserID string `json:"userId"`
}

func IBMQuantumExperience() interface{} {
    //resp, _ := http.Post(CONFIG_BASE.URL + "/users/loginWithToken", "application/json", &buf)
    return Api{}
}

func (api *Api) Get_Code() interface{} {
    return 1
}

func (api *Api) Get_Last_Codes() interface{} {
    return 2
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