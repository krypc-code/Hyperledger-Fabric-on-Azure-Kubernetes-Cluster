package main

type ClientConfig struct {
	Version       string                         `json:"version" yaml:"version"`
	Client        Client                         `json:"client" yaml:"client"`
	Channels      ChannelDefauls                 `json:"channels" yaml:"channels"`
	Organizations map[string]Organization        `json:"organizations" yaml:"organizations"`
	Orderers      map[string]OrganizationDetails `json:"orderers" yaml:"orderers"`
	Peers         map[string]OrganizationDetails `json:"peers" yaml:"peers"`
	CAAuthorities map[string]OrganizationDetails `json:"ca_authorities" yaml:"certificateAuthorities"`
}

type GoSdkConfig struct {
	Version       string                         `json:"version" yaml:"version"`
	Client        Client                         `json:"client" yaml:"client"`
	Channels      ChannelDefauls                 `json:"channels" yaml:"channels"`
	Organizations map[string]Organization        `json:"organizations" yaml:"organizations"`
	Orderers      map[string]OrganizationDetails `json:"orderers" yaml:"orderers"`
	Peers         map[string]OrganizationDetails `json:"peers" yaml:"peers"`
	CAAuthorities map[string]OrganizationDetails `json:"ca_authorities" yaml:"certificateAuthorities"`
}

type ChannelPeersPermission struct {
	EndorsingPeer  bool `json:"endorsingPeer" yaml:"endorsingPeer"`
	ChaincodeQuery bool `json:"chaincodeQuery" yaml:"chaincodeQuery"`
	LedgerQuery    bool `json:"ledgerQuery" yaml:"ledgerQuery"`
	EventSource    bool `json:"eventSource" yaml:"eventSource"`
}

type ChannelDefauls struct {
	Default Channel `json:"default" yaml:"_default"`
}

type Channel struct {
	//Orderers []string                          `json:"orderers" yaml:"orderers"`
	Peers    map[string]ChannelPeersPermission `json:"peers" yaml:"peers"`
	Policies ChannelPolicy                     `json:"policies" yaml:"policies"`
}

type RetryOpts struct {
	Attempts       int     `json:"attempts" yaml:"attempts"`
	InitialBackoff string  `json:"initialBackoff" yaml:"initialBackoff"`
	MaxBackoff     string  `json:"maxBackoff" yaml:"maxBackoff"`
	BackoffFactor  float64 `json:"backoffFactor" yaml:"backoffFactor"`
}

type ChannelPolicy struct {
	Discovery          Discovery `json:"discovery" yaml:"discovery"`
	Selection          Selection `json:"selection" yaml:"selection"`
	QueryChannelConfig Discovery `json:"queryChannelConfig" yaml:"queryChannelConfig"`
}

type Selection struct {
	SortingStrategy         string `json:"SortingStrategy" yaml:"SortingStrategy,omitempty"`
	Balancer                string `json:"Balancer" yaml:"Balancer,omitempty"`
	BlockHeightLagThreshold int    `json:"BlockHeightLagThreshold" yaml:"BlockHeightLagThreshold,omitempty"`
}

type LoggingC struct {
	Level string `json:"level" yaml:"level"`
}

type Timeout struct {
	Timeout TimeoutDetails `json:"timeout" yaml:"timeout"`
}

type Global struct {
	Timeout TimeoutDetails `json:"timeout" yaml:"timeout"`
	Cache   Cache          `json:"cache" yaml:"cache"`
}

type Cache struct {
	ConnectionIdle    string `json:"connectionIdle" yaml:"connectionIdle,omitempty"`
	EventServiceIdle  string `json:"eventServiceIdle" yaml:"eventServiceIdle,omitempty"`
	ChannelConfig     string `json:"channelConfig" yaml:"channelConfig,omitempty"`
	ChannelMembership string `json:"channelMembership" yaml:"channelMembership,omitempty"`
}

type TimeoutDetails struct {
	Connection           string    `json:"connection" yaml:"connection,omitempty"`
	Response             string    `json:"response" yaml:"response,omitempty"`
	QueryResponse        string    `json:"query_response" yaml:"queryResponse,omitempty"`
	RegistrationResponse string    `json:"registration_response" yaml:"registrationResponse,omitempty"`
	ExecuteTxResponse    string    `json:"execute_tx_response" yaml:"executeTxResponse,omitempty"`
	Discovery            Discovery `json:"discovery" yaml:"discovery,omitempty"`
	Query                string    `json:"query" yaml:"query,omitempty"`
	Execute              string    `json:"execute" yaml:"execute,omitempty"`
	Resmgmt              string    `json:"resmgmt" yaml:"resmgmt,omitempty"`
}

type Discovery struct {
	GreyListExpiry string    `json:"greylistExpiry" yaml:"greylistExpiry,omitempty"`
	MinResponses   int       `json:"minResponses" yaml:"minResponses,omitempty"`
	MaxTargets     int       `json:"maxTargets" yaml:"maxTargets,omitempty"`
	RetryOpts      RetryOpts `json:"retryOpts" yaml:"retryOpts,omitempty"`
}

type Path struct {
	Path string `json:"path" yaml:"path"`
}

type CredentialStore struct {
	Path        string `json:"path" yaml:"path,omitempty"`
	CryptoStore Path   `json:"crypto_store" yaml:"cryptoStore,omitempty"`
	Wallet      string `json:"wallet" yaml:"wallet,omitempty"`
}

type Security struct {
	Enabled       bool          `json:"enabled" yaml:"enabled"`
	Default       BCCSPDefaults `json:"default" yaml:"default"`
	HashAlgorithm string        `json:"hash_algorithm" yaml:"hashAlgorithm"`
	SoftVerify    bool          `json:"soft_verify" yaml:"softVerify"`
	Ephemeral     bool          `json:"ephemeral" yaml:"ephemeral"`
	Level         int           `json:"level" yaml:"level"`
}

type BCCSPDefaults struct {
	Provider string `json:"provider" yaml:"provider"`
}

type BCCSP struct {
	Security Security `json:"security" yaml:"security"`
}

type Organization struct {
	MspId      string   `json:"msp_id" yaml:"mspid"`
	CryptoPath string   `json:"crypto_path" yaml:"cryptoPath"`
	Peers      []string `json:"peers" yaml:"peers,omitempty"`
	CA         []string `json:"ca" yaml:"certificateAuthorities,omitempty"`
}

type GrpcOpts struct {
	SslTargetName    string `json:"ssl_target_name" yaml:"ssl-target-name-override"`
	MaxMessage       int    `json:"max_message" yaml:"grpc-max-send-message-length,omitempty"`
	KeepAliveTime    string `json:"keep-alive-time" yaml:"keep-alive-time,omitempty"`
	KeepAliveTimeOut string `json:"keep-alive-timeout" yaml:"keep-alive-timeout,omitempty"`
	KeepAlivePermit  bool   `json:"keep-alive-permit" yaml:"keep-alive-permit"`
	FailFast         bool   `json:"fail-fast" yaml:"fail-fast"`
	AllowInsecure    bool   `json:"allow-insecure" yaml:"allow-insecure"`
}

type HttpOpts struct {
	Verify bool `json:"verify" yaml:"verify"`
}

type CAClient struct {
	Keyfile  string `json:"keyfile" yaml:"keyfile"`
	CertFile string `json:"cert_file" yaml:"certfile"`
}

type TLSCACert struct {
	Path   string   `json:"path" yaml:"path"`
	Client CAClient `json:"client" yaml:"client,omitempty"`
}

type Registrar struct {
	EnrollId     string `json:"enroll_id" yaml:"enrollId"`
	EnrollSecret string `json:"enroll_secret" yaml:"enrollSecret"`
}

type OrganizationDetails struct {
	Url       string    `json:"url" yaml:"url"`
	EventUrl  string    `json:"event_url" yaml:"eventUrl,omitempty"`
	GrpcOpts  GrpcOpts  `json:"grpc_opts" yaml:"grpcOptions,omitempty"`
	HttpOpts  HttpOpts  `json:"http_opts" yaml:"httpOptions,omitempty"`
	TlsCACert TLSCACert `json:"tls_ca_cert" yaml:"tlsCACerts,omitempty"`
	Registrar Registrar `json:"registrar" yaml:"registrar,omitempty"`
	CAName    string    `json:"ca_name" yaml:"caName,omitempty"`
}

type Client struct {
	Organization    string          `json:"organization" yaml:"organization"`
	Logging         LoggingC        `json:"logging" yaml:"logging"`
	Peer            Timeout         `json:"peer" yaml:"peer"`
	EventService    EventService    `json:"event_service" yaml:"eventService"`
	Orderer         Timeout         `json:"orderer" yaml:"orderer"`
	Global          Global          `json:"global" yaml:"global"`
	CryptoConfig    Path            `json:"crypto_config" yaml:"cryptoconfig"`
	CredentialStore CredentialStore `json:"credential_store" yaml:"credentialStore"`
	Bccsp           BCCSP           `json:"bccsp" yaml:"BCCSP"`
	TlsCerts        TlsCert         `json:"tlsCerts" yaml:"tlsCerts"`
}

type TlsCert struct {
	SystemCertPool bool      `json:"systemCertPool" yaml:"systemCertPool"`
	Client         TlsCLient `json:"client" yaml:"client"`
}

type TlsCLient struct {
	Key  Path `json:"key" yaml:"key"`
	Cert Path `json:"cert" yaml:"cert"`
}

type EventService struct {
	Type                             string         `json:"type" yaml:"type,omitempty"`
	ResolverStrategy                 string         `json:"resolverStrategy" yaml:"resolverStrategy,omitempty"`
	Balancer                         string         `json:"balancer" yaml:"balancer,omitempty"`
	BlockHeightLagThreshold          string         `json:"blockHeightLagThreshold" yaml:"blockHeightLagThreshold,omitempty"`
	ReconnectBlockHeightLagThreshold string         `json:"reconnectBlockHeightLagThreshold" yaml:"reconnectBlockHeightLagThreshold,omitempty"`
	BlockHeightMonitorPeriod         string         `json:"blockHeightMonitorPeriod" yaml:"blockHeightMonitorPeriod,omitempty"`
	Timeout                          TimeoutDetails `json:"timeout" yaml:"timeout,omitempty"`
}
