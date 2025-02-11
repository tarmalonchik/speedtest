package svc

type Config struct {
	BankHost    string `envconfig:"BANK_HOST" default:"127.0.0.1"`
	BankPort    string `envconfig:"BANK_PORT" default:"8081"`
	MyIpAddress string `envconfig:"MY_IP_ADDRESS" required:"true"`
}
