package postgresql

type Config struct {
	Host                  string
	Port                  string
	Username              string
	Password              string
	DBname                string
	MaxConnection         string // connection pool'da max kac connection olabilir
	MaxConnectionIdleTime string
}
