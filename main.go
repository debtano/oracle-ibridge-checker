package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/debtano/oci/idcs/ibridge/check"
	"github.com/tkanos/gonfig"
)

const (
	tokenEndpoint          = ".identity.oraclecloud.com/oauth2/v1/token"
	identityAgentsEndpoint = ".identity.oraclecloud.com/admin/v1/IdentityAgents"
)

// Configuration struct is used to obtain config file data
type Configuration struct {
	Tenant       string
	IDCS         string
	ClientID     string
	ClientSecret string
}

func main() {

	configuration := Configuration{}
	err := gonfig.GetConf("config.json", &configuration)
	tenant, idcs, clientID, clientSecret := configuration.Tenant, configuration.IDCS, configuration.ClientID, configuration.ClientSecret
	// fmt.Println(configuration.Tenant)

	var tok check.Token
	// var resource []struct{ check.Resources }

	val := url.Values{}
	val.Add("grant_type", "client_credentials")
	val.Add("scope", "urn:opc:idm:__myscopes__")

	tokenURL := idcs + tokenEndpoint
	identityAgentsEndpointURL := idcs + identityAgentsEndpoint

	body := check.GetTokenBody(&val, tokenURL, clientID, clientSecret)
	tokenString := check.GetTokenString(&tok, body)
	agentsData, err := check.GetIdentityAgentsData(tokenString, identityAgentsEndpointURL)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("%d", agentsData.TotalResources)
	fmt.Printf("Tenant name : %s\n", tenant)
	for _, agent := range agentsData.Resources {
		fmt.Printf("Hostname => %s\tStatus => %s\n", agent.Hostname, agent.Status)
		for _, instance := range agent.Sources {
			fmt.Printf("SyncState => %s\tDomain => %s\n", instance.SyncState, instance.Domain)
		}
	}
}
