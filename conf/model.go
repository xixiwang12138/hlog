package conf

type ENV uint8

const (
	Dev ENV = iota
	Prod
)

type MongoDBConfig struct {
	UserName string `yaml:"userName" json:"userName"`
	Password string `yaml:"password" json:"password"`
	Host     string `yaml:"host" json:"host"`
	Port     string `yaml:"port" json:"port"`
	DataBase string `yaml:"dataBase" json:"dataBase"`
	env      ENV
}
