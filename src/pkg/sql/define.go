package sql

type Config struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
	DBPath      string
	Port        int
	Charset     string
}
