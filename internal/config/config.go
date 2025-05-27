package config

type LogConfig struct {
	ID   string `json:"id"`
	Path string `json:"path"`
	Type string `json:"type"`
}

func LoadConfig(path string) ([]LogConfig, error) {
	// TODO: Implémenter le chargement de la configuration
	return nil, nil
}
