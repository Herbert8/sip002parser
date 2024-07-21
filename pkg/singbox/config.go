package singbox

type Config struct {
	Log       LogConfig `json:"log"`
	DNS       DNSConfig `json:"dns,omitempty"`
	Inbounds  []Bound   `json:"inbounds"`
	Outbounds []Bound   `json:"outbounds"`
	Route     struct {
	} `json:"route,omitempty"`
}

type LogConfig struct {
	Disabled  bool   `json:"disabled"`
	Level     string `json:"level"`
	Output    string `json:"output"`
	Timestamp bool   `json:"timestamp"`
}

type DNSConfig interface{}

//type DNSServer struct {
//	Tag             string `json:"tag"`
//	Address         string `json:"address"`
//	AddressResolver string `json:"address_resolver"`
//	AddressStrategy string `json:"address_strategy"`
//	Strategy        string `json:"strategy"`
//	Detour          string `json:"detour"`
//	ClientSubnet    string `json:"client_subnet"`
//}

type Bound interface {
	GetType() string
	GetTag() string
}

type BaseBound struct {
	Type string `json:"type"`
	Tag  string `json:"tag"`
}

func (bound BaseBound) GetType() string {
	return bound.Type
}

func (bound BaseBound) GetTag() string {
	return bound.Tag
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type MixedInbound struct {
	BaseBound
	Listen         string `json:"listen"`
	ListenPort     uint   `json:"listen_port"`
	Users          []User `json:"users"`
	SetSystemProxy bool   `json:"set_system_proxy"`
}

type ShadowsocksObfsOutbound struct {
	BaseBound
	Server     string `json:"server"`
	ServerPort int    `json:"server_port"`
	Method     string `json:"method"`
	Password   string `json:"password"`
	Plugin     string `json:"plugin"`
	PluginOpts string `json:"plugin_opts"`
}

type SelectorOutbound struct {
	BaseBound
	Outbounds                 []string `json:"outbounds"`
	Default                   string   `json:"default"`
	InterruptExistConnections bool     `json:"interrupt_exist_connections"`
}

type URLTestOutbound struct {
	BaseBound
	Outbounds                 []string `json:"outbounds"`
	Url                       string   `json:"url"`
	Interval                  string   `json:"interval"`
	Tolerance                 int      `json:"tolerance"`
	IdleTimeout               string   `json:"idle_timeout"`
	InterruptExistConnections bool     `json:"interrupt_exist_connections"`
}
