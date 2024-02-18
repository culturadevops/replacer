/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "replacer",
	Short: "reemplazar palabras clave en un archivo de destino",
	Long: `toma un archivo de leyenda y remplaza en otro archivo, necesitas la ruta del archivo base con las definiciones o leyendas y la ruta final del archivo destino y por ultimo el separador del archivo base es decir un = o un : 
ejemplo del archivo base
	 nombre=juan
	 direccion= casa roja luego de otra
	 
el proceso tomara estos datos y los remplazara en el archivo destino donde consiga la palabra nombre y la palabra direccion
	 `,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.RemplazadorDeDatos.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
