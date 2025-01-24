package http

type Config struct {
	Port              uint16 `yaml:"port"`
	ReadHeaderTimeout string `yaml:"readHeaderTimeout"`
	IdleTimeout       string `yaml:"idleTimeout"`
}
