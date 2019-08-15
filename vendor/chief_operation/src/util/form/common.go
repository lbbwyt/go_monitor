package form

//此文件保存公共的结构

type Mysql struct {
	Host   string `json:"host"`
	Db     string `json:"db"`
	MaxCon int    `json:"max_con"`
}

type Redis struct {
	Host     string `json:"host"`
	Password string `json:"password"`
	Db       int    `json:"db"`
}
type Dingding struct {
	Path     string `json:"path"`
	Send     int `json:"send"`
	People        []string `json:"people"`
	Limit int `json:"limit_send"`
}



type CommonConf struct {
	Mysql         Mysql `json:"mysql"`
	MysqlLog      Mysql `json:"mysql_log"`
	MysqlActivity Mysql `json:"mysql_activity"`
	Redis         Redis `json:"redis"`
	Dingding         Dingding `json:"dingding"`
}


type PortConf struct {
	Org string `json:"org"`
	Url   string `json:"url"`
	Ports         []string `json:"ports"`
}
