package reporter

import (
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
	// TODO: Implémenter l'export des résultats en JSON
	return nil
}

func (r *Reporter) DisplayResults() {
	// TODO: Implémenter l'affichage des résultats
}
