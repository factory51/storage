package config

type AppConfig struct {
	Database     DatabaseServer `json:"database"`
	Redis        RedisServer    `json:"redis"`
	AppIdent     string         `json:"app_ident"`
	IpRestrict   bool           `json:"ip_restricted"`
	IpList       []string       `json:"ip_list"`
	AppSecretKey string         `json:"app_secret_key"`
}

type DatabaseServer struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type RedisServer struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
}
