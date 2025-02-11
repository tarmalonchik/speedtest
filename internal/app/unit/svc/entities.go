package svc

type Config struct {
	BankHost    string `envconfig:"BANK_HOST" required:"true"`
	BankPort    string `envconfig:"BANK_PORT" required:"true"`
	MyIpAddress string `envconfig:"MY_IP_ADDRESS" required:"true"`
}
