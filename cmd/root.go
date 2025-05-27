package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "loganalyzer",
	Short: "Un outil d'analyse de logs distribuée",
	Long: `LogAnalyzer est un outil en ligne de commande qui permet d'analyser 
des fichiers de logs provenant de diverses sources en parallèle et d'en 
extraire des informations clés.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
