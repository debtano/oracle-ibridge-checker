package check

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	identityAgentEndpoint = ".identity.oraclecloud.com/admin/v1/IdentityAgents"
)

// Token struct for token
type Token struct {
	Token  string `json:"access_token"`
	Expire int    `json:"expires_in"`
}

// IdentityAgents struct contain agent structs for identity agents
type IdentityAgents struct {
	TotalResources int `json:"totalResults"`
	Resources      []*Agent
}

//Agent struct contain agents information
type Agent struct {
	Hostname string `json:"hostName"`
	Type     string `json:"type"`
	Status   string `json:"status"`
	Version  string `json:"version"`
	Sources  []*Instance
}

//Instance struct contain instantiated agents
type Instance struct {
	SyncState string `json:"currentSyncState"`
	Active    bool   `json:"active"`
	Domain    string `json:"display"`
	SourceID  string `json:"value"`
}

// GetTokenBody shoud obtain the Body of the token endpoint request
func GetTokenBody(data *url.Values, url, clientID, clientSecret string) []byte {

	client := http.Client{}
	request, err := http.NewRequest("POST", url, strings.NewReader(data.Encode()))
	request.SetBasicAuth(clientID, clientSecret)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return body
}

// GetTokenString should unmarshal token Body
func GetTokenString(tok *Token, body []byte) string {
	_ = json.Unmarshal(body, &tok)
	tokenString := tok.Token
	return tokenString
}

// GetIdentityAgentsData should populate Resources struct for agents of a given IDCS instance
func GetIdentityAgentsData(tokenString string, identityAgentsEnpoint string) (*IdentityAgents, error) {
	authToken := "Bearer " + tokenString
	client := http.Client{}
	request, err := http.NewRequest("GET", identityAgentsEnpoint, nil)
	request.Header.Add("Content-Type", "application/scim+json")
	request.Header.Add("Authorization", authToken)

	if err != nil {
		log.Fatalln(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	var agentsData IdentityAgents
	if err := json.NewDecoder(resp.Body).Decode(&agentsData); err != nil {
		resp.Body.Close()
		return nil, err
	}
	// _ = json.Unmarshal(body, &agentsData)
	resp.Body.Close()
	return &agentsData, nil
	//fmt.Println(body)
}
