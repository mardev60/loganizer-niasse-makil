package cmd

import (
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
		// TODO: Implémenter la logique d'analyse
		return nil
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)

	analyzeCmd.Flags().StringVarP(&configFile, "config", "c", "", "Chemin vers le fichier de configuration JSON (requis)")
	analyzeCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Chemin vers le fichier de sortie JSON")

	analyzeCmd.MarkFlagRequired("config")
}
