package cmd

import (
	"fmt"

	"github.com/mardev60/loganizer-niasse-makil/internal/analyzer"
	"github.com/mardev60/loganizer-niasse-makil/internal/config"
	"github.com/mardev60/loganizer-niasse-makil/internal/reporter"
	"github.com/spf13/cobra"
)

var (
	configFile string
	outputFile string
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyse les fichiers de logs spécifiés",
	Long: `Analyse les fichiers de logs spécifiés dans le fichier de configuration JSON.
Les résultats peuvent être exportés dans un fichier JSON de sortie.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Chargement de la configuration depuis: %s\n", configFile)
		configs, err := config.LoadConfig(configFile)
		if err != nil {
			return fmt.Errorf("erreur lors du chargement de la configuration: %w", err)
		}

		if len(configs) == 0 {
			return fmt.Errorf("aucune configuration de log trouvée")
		}

		fmt.Printf("Configuration chargée: %d logs à analyser\n", len(configs))

		fmt.Println("Démarrage de l'analyse en parallèle...")
		analyzer := analyzer.NewAnalyzer(configs)
		results, err := analyzer.AnalyzeLogs()
		if err != nil {
			return fmt.Errorf("erreur lors de l'analyse: %w", err)
		}

		rep := reporter.NewReporter(results)
		rep.DisplayResults()

		if outputFile != "" {
			fmt.Printf("\nExport des résultats vers: %s\n", outputFile)
			if err := rep.ExportJSON(outputFile); err != nil {
				return fmt.Errorf("erreur lors de l'export JSON: %w", err)
			}
			fmt.Println("Export terminé avec succès!")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)

	analyzeCmd.Flags().StringVarP(&configFile, "config", "c", "", "Chemin vers le fichier de configuration JSON (requis)")
	analyzeCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Chemin vers le fichier de sortie JSON")

	analyzeCmd.MarkFlagRequired("config")
}
