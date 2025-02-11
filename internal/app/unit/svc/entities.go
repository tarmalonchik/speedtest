package svc

type Config struct {
	BankServerAddress string `envconfig:"BANK_SERVER_ADDRESS" default:"127.0.0.1:8081"`
	MyIpAddress       string `envconfig:"MY_IP_ADDRESS" required:"true"`
}
