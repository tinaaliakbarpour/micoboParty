package config

type Service struct {
	Name    string `yaml:"service.name"`
	ID      uint32 `yaml:"service.id"`
	BaseURL string `yaml:"service.baseURL"`
	HTTP    struct {
		Host           string `yaml:"http.host"`
		Port           string `yaml:"http.port"`
		RequestTimeout string `yaml:"http.requestTimeout"`
	}
}
