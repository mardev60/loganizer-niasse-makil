package analyzer

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

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
	var wg sync.WaitGroup
	resultsChan := make(chan Result, len(a.configs))

	for _, cfg := range a.configs {
		wg.Add(1)
		go a.analyzeLog(cfg, resultsChan, &wg)
	}

	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	var results []Result
	for result := range resultsChan {
		results = append(results, result)
	}

	return results, nil
}

func (a *Analyzer) analyzeLog(cfg config.LogConfig, resultsChan chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	result := Result{
		LogID:    cfg.ID,
		FilePath: cfg.Path,
	}

	if _, err := os.Stat(cfg.Path); err != nil {
		if os.IsNotExist(err) {
			fileErr := &FileNotFoundError{Path: cfg.Path}
			result.Status = "FAILED"
			result.Message = "Fichier introuvable."
			result.ErrorDetails = fileErr.Error()
		} else {
			result.Status = "FAILED"
			result.Message = "Erreur d'accès au fichier."
			result.ErrorDetails = err.Error()
		}
		resultsChan <- result
		return
	}

	analysisTime := 50 + rand.Intn(151)
	time.Sleep(time.Duration(analysisTime) * time.Millisecond)

	if rand.Float32() < 0.1 {
		parseErr := &ParseError{
			Path:    cfg.Path,
			Message: "format de log invalide détecté",
		}
		result.Status = "FAILED"
		result.Message = "Erreur de parsing."
		result.ErrorDetails = parseErr.Error()
	} else {
		result.Status = "OK"
		result.Message = "Analyse terminée avec succès."
		result.ErrorDetails = ""
	}

	resultsChan <- result
}

func IsFileNotFoundError(err error) bool {
	var fileErr *FileNotFoundError
	return errors.As(err, &fileErr)
}

func IsParseError(err error) bool {
	var parseErr *ParseError
	return errors.As(err, &parseErr)
}
