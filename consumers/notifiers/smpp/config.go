package smpp

import (
	"crypto/tls"
)

// Config represents SMPP transmitter configuration.
type Config struct {
	Address       string `env:"SINA_SMPP_ADDRESS"       envDefault:""`
	Username      string `env:"SINA_SMPP_USERNAME"      envDefault:""`
	Password      string `env:"SINA_SMPP_PASSWORD"      envDefault:""`
	SystemType    string `env:"SINA_SMPP_SYSTEM_TYPE"   envDefault:""`
	SourceAddrTON uint8  `env:"SINA_SMPP_SRC_ADDR_TON"  envDefault:"0"`
	SourceAddrNPI uint8  `env:"SINA_SMPP_DST_ADDR_TON"  envDefault:"0"`
	DestAddrTON   uint8  `env:"SINA_SMPP_SRC_ADDR_NPI"  envDefault:"0"`
	DestAddrNPI   uint8  `env:"SINA_SMPP_DST_ADDR_NPI"  envDefault:"0"`
	TLS           *tls.Config
}