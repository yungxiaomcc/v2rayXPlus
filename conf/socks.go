package conf

type SocksAccount struct {
	Username string `json:"user"`
	Password string `json:"pass"`
}

type SocksServerConfig struct {
	AuthMethod string          `json:"auth"`
	Accounts   []*SocksAccount `json:"accounts"`
	UDP        bool            `json:"udp"`
	Host       string          `json:"ip"`
	Timeout    uint32          `json:"timeout"`
	UserLevel  uint32          `json:"userLevel"`
}

type SocksRemoteConfig struct {
	Address string         `json:"address"`
	Port    uint16         `json:"port"`
	Users   []SocksAccount `json:"users"`
}
type SocksClientConfig struct {
	Servers []*SocksRemoteConfig `json:"servers"`
}
