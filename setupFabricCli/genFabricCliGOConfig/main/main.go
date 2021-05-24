package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

const (
	ORG_CRYPTOPATH  = "%s/%s/users/admin.%s/msp"
	ORG_TLS_CERT    = "cryptoconfig/%s/%s/%s/msp/tlscacerts/ca.crt"
	CA_ORG_TLS_CERT = "cryptoconfig/%s/%s/ca/tlsca/cert.pem"
)

func main() {

	//PEER_CONNECTION_PROFILE_PATH
	peerConnectionProfilePath := os.Getenv("PEER_CONNECTION_PROFILE_PATH")
	log.Println("Connection profile path of peer is ", peerConnectionProfilePath)

	//ORDERER_CONNECTION_PROFILE_PATH
	ordererConnectionProfilePath := os.Getenv("ORDERER_CONNECTION_PROFILE_PATH")
	log.Println("Connection profile path of peer is ", ordererConnectionProfilePath)

	peerConnectionProfileBytes, err := ioutil.ReadFile(peerConnectionProfilePath)
	if err != nil {
		log.Fatalln("Error while reading file in peerConnectionProfilePath ", err.Error())
	}

	peerConnectionProfileMap := make(map[string]interface{})

	err = json.Unmarshal(peerConnectionProfileBytes, &peerConnectionProfileMap)
	if err != nil {
		log.Fatalln("Error while converting byte into map ", err.Error())
	}

	ordererConnectionProfileBytes, err := ioutil.ReadFile(ordererConnectionProfilePath)
	if err != nil {
		log.Fatalln("Error while reading file in peerAdminProfilePath ", err.Error())
	}

	ordererConnectionProfileMap := make(map[string]interface{})

	err = json.Unmarshal(ordererConnectionProfileBytes, &ordererConnectionProfileMap)
	if err != nil {
		log.Fatalln("Error while converting byte into map ", err.Error())
	}

	peerOrgName := peerConnectionProfileMap["name"].(string)
	ordererOrgName := ordererConnectionProfileMap["name"].(string)

	config := GoSdkConfig{
		Version: "1.0.0",
		Client: Client{
			Organization: peerOrgName,
			Logging: LoggingC{
				Level: "info",
			},
			Peer: Timeout{
				Timeout: TimeoutDetails{
					Connection:        "3S",
					Response:          "40s",
					QueryResponse:     "45s",
					ExecuteTxResponse: "60s",
					Discovery: Discovery{
						GreyListExpiry: "5s",
					},
				},
			},
			EventService: EventService{
				Type:                             "deliver",
				ResolverStrategy:                 "PreferOrg",
				Balancer:                         "RoundRobin",
				BlockHeightLagThreshold:          "2",
				ReconnectBlockHeightLagThreshold: "2",
				BlockHeightMonitorPeriod:         "3s",
				Timeout: TimeoutDetails{
					Connection: "3s",
					Response:   "5s",
				},
			},
			Orderer: Timeout{
				Timeout: TimeoutDetails{
					Connection: "3s",
					Response:   "5s",
				},
			},
			Global: Global{
				Timeout: TimeoutDetails{
					Query:   "45s",
					Execute: "60s",
					Resmgmt: "60s",
				},
				Cache: Cache{
					ConnectionIdle:    "30s",
					EventServiceIdle:  "2m",
					ChannelConfig:     "60s",
					ChannelMembership: "30s",
				},
			},
			CryptoConfig: Path{
				Path: "cryptoconfig",
			},
			CredentialStore: CredentialStore{
				Path: "temp/msp",
				CryptoStore: Path{
					Path: "temp/store",
				},
			},
			Bccsp: BCCSP{
				Security: Security{
					Enabled: true,
					Default: BCCSPDefaults{
						Provider: "SW",
					},
					HashAlgorithm: "SHA2",
					SoftVerify:    true,
					Ephemeral:     false,
					Level:         256,
				},
			},
			TlsCerts: TlsCert{
				SystemCertPool: false,
				Client: TlsCLient{
					Cert: Path{},
					Key:  Path{},
				},
			},
		},
		Channels: ChannelDefauls{
			Default: Channel{
				Peers: make(map[string]ChannelPeersPermission),
				Policies: ChannelPolicy{
					Discovery: Discovery{
						MaxTargets: 2,
						RetryOpts: RetryOpts{
							Attempts:       4,
							InitialBackoff: "500ms",
							MaxBackoff:     "5s",
							BackoffFactor:  2.0,
						},
					},
					Selection: Selection{
						SortingStrategy:         "BlockHeightPriority",
						Balancer:                "RoundRobin",
						BlockHeightLagThreshold: 5,
					},
					QueryChannelConfig: Discovery{
						MinResponses: 1,
						MaxTargets:   1,
						RetryOpts: RetryOpts{
							Attempts:       5,
							InitialBackoff: "500ms",
							MaxBackoff:     "5s",
							BackoffFactor:  2.0,
						},
					},
				},
			},
		},
		Organizations: make(map[string]Organization),
		Orderers:      make(map[string]OrganizationDetails),
		Peers:         make(map[string]OrganizationDetails),
		CAAuthorities: make(map[string]OrganizationDetails),
	}

	peersInterface := peerConnectionProfileMap["organizations"].(map[string]interface{})[peerOrgName].(map[string]interface{})["peers"].([]interface{})
	var peers []string
	peersMap := make(map[string]string)
	for _, peerInterface := range peersInterface {
		peername := peerInterface.(string)
		peers = append(peers, peername)
		config.Channels.Default.Peers[peername] = ChannelPeersPermission{
			EndorsingPeer:  true,
			ChaincodeQuery: true,
			LedgerQuery:    true,
			EventSource:    true,
		}
		peersMap[peerInterface.(string)] = peerConnectionProfileMap["peers"].(map[string]interface{})[peername].(map[string]interface{})["grpcOptions"].(map[string]interface{})["hostnameOverride"].(string)
	}
	peerCaName := peerConnectionProfileMap["certificateAuthorities"].(map[string]interface{})[peerOrgName+"CA"].(map[string]interface{})["caName"].(string)
	peerOrganization := Organization{
		MspId:      peerOrgName,
		CryptoPath: fmt.Sprintf(ORG_CRYPTOPATH, "peer", peerOrgName, peerOrgName),
		Peers:      peers,
		CA:         []string{peerCaName},
	}

	orderersInterface := ordererConnectionProfileMap["organizations"].(map[string]interface{})[ordererOrgName].(map[string]interface{})["orderers"].([]interface{})
	var orderers []string
	orderersMap := make(map[string]string)
	for _, ordererInterface := range orderersInterface {
		orderername := ordererInterface.(string)
		orderers = append(orderers, orderername)
		orderersMap[ordererInterface.(string)] = ordererConnectionProfileMap["orderers"].(map[string]interface{})[orderername].(map[string]interface{})["grpcOptions"].(map[string]interface{})["hostnameOverride"].(string)
	}
	//ordererCaName := ordererConnectionProfileMap["certificateAuthorities"].(map[string]interface{})[ordererOrgName+"CA"].(map[string]interface{})["caName"].(string)
	ordererOrganization := Organization{
		MspId:      ordererOrgName,
		CryptoPath: fmt.Sprintf(ORG_CRYPTOPATH, "orderer", ordererOrgName, ordererOrgName),
	}

	log.Println("client is ", config)
	config.Organizations[peerOrganization.MspId] = peerOrganization
	config.Organizations[ordererOrganization.MspId] = ordererOrganization

	for _, peer := range peers {
		peerOrgDetails := OrganizationDetails{
			Url: peersMap[peer] + ":443",
			GrpcOpts: GrpcOpts{
				SslTargetName:    peersMap[peer],
				KeepAliveTime:    "0s",
				KeepAliveTimeOut: "20s",
				KeepAlivePermit:  false,
				FailFast:         false,
				AllowInsecure:    false,
			},
			TlsCACert: TLSCACert{
				Path: fmt.Sprintf(ORG_TLS_CERT, "peer", peerOrgName, peer),
			},
		}
		config.Peers[peer] = peerOrgDetails
	}

	for _, orderer := range orderers {
		ordererOrgDetails := OrganizationDetails{
			Url: orderersMap[orderer] + ":443",
			GrpcOpts: GrpcOpts{
				SslTargetName:    orderersMap[orderer],
				KeepAliveTime:    "0s",
				KeepAliveTimeOut: "20s",
				KeepAlivePermit:  false,
				FailFast:         false,
				AllowInsecure:    false,
			},
			TlsCACert: TLSCACert{
				Path: fmt.Sprintf(ORG_TLS_CERT, "orderer", ordererOrgName, orderer),
			},
		}
		config.Orderers[orderer] = ordererOrgDetails
	}

	peerCaDetails := OrganizationDetails{
		Url: peerConnectionProfileMap["certificateAuthorities"].(map[string]interface{})[peerOrgName+"CA"].(map[string]interface{})["url"].(string),
		HttpOpts: HttpOpts{
			Verify: true,
		},
		TlsCACert: TLSCACert{
			Path: fmt.Sprintf(CA_ORG_TLS_CERT, "peer", peerOrgName),
		},
		CAName: peerCaName,
	}
	config.CAAuthorities[peerCaName] = peerCaDetails

	yamlBytes, err := yaml.Marshal(config)
	if err != nil {
		log.Fatalln("Error while marshaling struct ", err.Error())
	}

	f, err := os.Create(peerOrgName + "-config.yaml")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err := f.Write(yamlBytes); err != nil {
		panic(err)
	}
	fmt.Printf("--- m dump:\n%s\n\n", string(yamlBytes))

}
