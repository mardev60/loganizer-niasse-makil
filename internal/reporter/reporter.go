package reporter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mardev60/loganizer-niasse-makil/internal/analyzer"
)

type Reporter struct {
	results []analyzer.Result
}

func NewReporter(results []analyzer.Result) *Reporter {
	return &Reporter{
		results: results,
	}
}

func (r *Reporter) ExportJSON(path string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("impossible de créer les dossiers parents: %w", err)
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("impossible de créer le fichier de sortie: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Pour un JSON plus lisible
	if err := encoder.Encode(r.results); err != nil {
		return fmt.Errorf("erreur lors de l'encodage JSON: %w", err)
	}

	return nil
}

func (r *Reporter) DisplayResults() {
	fmt.Println("\n=== RÉSULTATS DE L'ANALYSE ===")
	fmt.Printf("Nombre total de logs analysés: %d\n\n", len(r.results))

	successCount := 0
	failureCount := 0

	for _, result := range r.results {
		status := "✓"
		if result.Status == "FAILED" {
			status = "✗"
			failureCount++
		} else {
			successCount++
		}

		fmt.Printf("%s [%s] %s\n", status, result.LogID, result.FilePath)
		fmt.Printf("  Status: %s\n", result.Status)
		fmt.Printf("  Message: %s\n", result.Message)

		if result.ErrorDetails != "" {
			fmt.Printf("  Erreur: %s\n", result.ErrorDetails)
		}
		fmt.Println()
	}

	fmt.Printf("=== RÉSUMÉ ===\n")
	fmt.Printf("Succès: %d\n", successCount)
	fmt.Printf("Échecs: %d\n", failureCount)
}
