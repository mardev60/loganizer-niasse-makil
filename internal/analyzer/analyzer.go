package analyzer

import (
	"fmt"

	"github.com/mardev60/loganizer-niasse-makil/internal/config"
)

type Result struct {
	LogID        string `json:"log_id"`
	FilePath     string `json:"file_path"`
	Status       string `json:"status"`
	Message      string `json:"message"`
	ErrorDetails string `json:"error_details"`
}

type FileNotFoundError struct {
	Path string
}

func (e *FileNotFoundError) Error() string {
	return fmt.Sprintf("fichier introuvable: %s", e.Path)
}

type ParseError struct {
	Path    string
	Message string
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("erreur de parsing pour %s: %s", e.Path, e.Message)
}

type Analyzer struct {
	configs []config.LogConfig
}

func NewAnalyzer(configs []config.LogConfig) *Analyzer {
	return &Analyzer{
		configs: configs,
	}
}

func (a *Analyzer) AnalyzeLogs() ([]Result, error) {
	// TODO: Implémenter l'analys  des logs
	return nil, nil
}
