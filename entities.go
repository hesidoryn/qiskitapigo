package qiskitapigo

import "time"

type Credentials struct {
	ID       string    `json:"id"`
	TTL      int       `json:"ttl"`
	Created  time.Time `json:"created"`
	UserID   string    `json:"userId"`
	ApiToken string
}

type Api struct {
	Creds Credentials
}

type CodesResult struct {
	Total int    `json:"total"`
	Count int    `json:"count"`
	Codes []Code `json:"codes"`
}

type Code struct {
	Type         string    `json:"type"`
	Active       bool      `json:"acive"`
	VersionID    int       `json:"versionId"`
	IdCode       string    `json:"idCode"`
	Name         string    `json:"name"`
	Qasm         string    `json:"qasm"`
	CodeType     string    `json:"codeType"`
	CreationDate time.Time `json:"creationDate"`
	Deleted      bool      `json:"deleted"`
	OrderDate    int       `json:"orderDate"`
	UserDeleted  bool      `json:"userDeleted"`
	JSONQasm     JsonQASM  `json:"jsonQASM"`
	DisplayUrls  struct {
		Png string `json:"png"`
	} `json:"displayUrls"`
	ID     string `json:"id"`
	UserID string `json:"userID"`
}

type JsonQASM struct {
	Playgrounds     []Playground  `json:"playground"`
	NumberColumns   int           `json:"numberColumns"`
	NumberLines     int           `json:"numberLines"`
	NumberGates     int           `json:"numberGates"`
	HasMeasures     bool          `json:"hasMeasures"`
	Topology        string        `json:"topology"`
	HasBloch        bool          `json:"hasBloch"`
	GateDefinitions []interface{} `json:"gateDefinitions"`
}

type Playground struct {
	Line  int    `json:"line"`
	Name  string `json:"name"`
	Gates []Gate `json:"gates"`
}

type Gate struct {
	Position    int    `json:"position"`
	Name        string `json:"name"`
	Qasm        string `json:"qasm"`
	MeasureCreg struct {
		Line int `json:"line"`
		Bit  int `json:"bit"`
	} `json:"measureCreg"`
	IsOperation bool `json:"isOperation"`
}

type Execution struct {
	Result struct {
		Date time.Time `json:"date"`
		Data struct {
			P struct {
				Qubits []int     `json:"qubits"`
				Labels []string  `json:"labels"`
				Values []float64 `json:"values"`
			} `json:"p"`
		} `json:"data"`

		Qasm               string  `json:"qasm"`
		SerialNumberDevice string  `json:"serialNumberDevice"`
		Time               float64 `json:"time"` // ?????????????
	} `json:"result"`
	StartDate        time.Time `json:"startDate"`
	ModificationDate int       `json:"modificationDate"` // ?????????????
	Time             float64   `json:"time"`             // ????????????????????
	EndDate          time.Time `json:"endDate"`
	TypeCredits      string    `json:"typeCredits"`
	Status           struct {
		Id string `json:"id"`
	} `json:"status"`
	DeviceRunType string `json:"deviceRunType"`
	IP            struct {
		IP        string `json:"ip"`
		Country   string `json:"country"`
		Continent string `json:"continent"`
	} `json:"ip"`
	Calibration struct {
		Date              time.Time  `json:"date"`
		Device            string     `json:"device"`
		FridgeTemperature float64    `json:"fridge_temperature"`
		Properties        []Property `json:"properties"`
	} `json:"calibration"`
	Shots           int      `json:"shots"`
	ParamsCustomize struct{} `json:"paramsCustomize"`
	Deleted         bool     `json:"deleted"`
	UserDleted      bool     `json:"userDeleted"`
	ID              string   `json:"id"`
	CodeID          string   `json:"codeId"`
	DeviceID        string   `json:"deviceId"`
	UserID          string   `json:"userId"`
}

type Property struct {
	Values struct {
		E_G01 float64 `json:"e_g{01}"`
		E_G02 float64 `json:"e_g{02}"`
		E_G12 float64 `json:"e_g{12}"`
		E_G32 float64 `json:"e_g{32}"`
		E_G34 float64 `json:"e_g{34}"`
		E_G42 float64 `json:"e_g{42}"`
		T_2   float64 `json:"t_2"`
		E_R   float64 `json:"e_r"`
		T_1   float64 `json:"t_1"`
		E_G   float64 `json:"E_G"`
		F     float64 `json:"f"`
	} `json:"values"`
	Key string `json:"key"`
}
