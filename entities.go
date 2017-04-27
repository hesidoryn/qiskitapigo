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

type LastestCodes struct {
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
	Playgrounds   []Playground `json:"playground"`
	NumberColumns int          `json:"numberColumns"`
	NumberLines   int          `json:"numberLines"`
	NumberGates   int          `json:"numberGates"`
	HasMeasures   bool         `json:"hasMeasures"`
	Topology      string       `json:"topology"`
	HasBloch      bool         `json:"hasBloch"`
}

type Playground struct {
	Line  int    `json:"line"`
	Name  string `json:"name"`
	Gates []Gate `json:"gates"`
}

type Gate struct {
	Position int    `json:"position"`
	Name     string `json:"name"`
	Qasm     string `json:"qasm"`
}
