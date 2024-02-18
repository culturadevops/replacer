/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/spf13/cobra"
)

func BuscarCoincidencias(pat, archivoEntrada string) (map[string]bool, error) {
	// Abrir el archivo de entrada
	fileIn, err := os.Open(archivoEntrada)
	if err != nil {
		return nil, err
	}
	defer fileIn.Close()

	// Crear un mapa para almacenar las palabras encontradas
	palabras := make(map[string]bool)

	// Crear un compilador de expresión regular para el patrón
	regex := regexp.MustCompile(pat)

	// Crear un scanner para leer el archivo de entrada línea por línea
	scanner := bufio.NewScanner(fileIn)

	// Escanear el archivo de entrada y buscar coincidencias
	for scanner.Scan() {
		linea := scanner.Text()
		// Encontrar todas las coincidencias del patrón en la línea
		matches := regex.FindAllString(linea, -1)
		// Iterar sobre las coincidencias y agregarlas al mapa
		for _, match := range matches {
			// Eliminar los corchetes dobles del inicio y del final de la coincidencia
			// Agregar la palabra al mapa
			palabras[match] = true
		}
	}

	// Verificar errores durante el escaneo del archivo
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return palabras, nil
}
func EscribirMapaEnArchivo(m map[string]bool, archivoSalida string, separador string) error {
	// Abrir el archivo de salida
	fileOut, err := os.Create(archivoSalida)
	if err != nil {
		return err
	}
	defer fileOut.Close()

	// Escribir el contenido del mapa en el archivo
	for palabra := range m {
		_, err := fmt.Fprintf(fileOut, "%s%s\n", palabra, separador)
		if err != nil {
			return err
		}
	}

	return nil
}

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "busca en un archivo todas las palabras que coinciden con un patrón específico y las almacena en un documento",
	Long: `toma dos parámetros: pat, que es una expresión regular que define el patrón a buscar, y archivoEntrada, que es el nombre del archivo donde se buscarán las coincidencias. La función escanea el archivo línea por línea, busca todas las coincidencias del patrón especificado en cada línea utilizando expresiones regulares y luego almacena todas las palabras encontradas en un archivo.
ejemplo:
	remplace generate -i archivoaRemplazar.json -o salida.vars -s : -p "\[\[.*?\]\]"

	`,
	Run: func(cmd *cobra.Command, args []string) {

		inputFile, _ := cmd.Flags().GetString("input")
		output, _ := cmd.Flags().GetString("output")
		separated, _ := cmd.Flags().GetString("separated")
		pat, _ := cmd.Flags().GetString("pat")

		if inputFile == "" {
			fmt.Println("Falta el archivo inputFile agregue -i")
			return
		}
		if output == "" {
			fmt.Println("Falta el archivo output agregue -o")
			return
		}
		if separated == "" {
			fmt.Println("Falta el archivo separacion agregue -s")
			return
		}
		if pat == "" {
			fmt.Println("Falta el archivo patron agregue -p")
			return
		}

		//palabrasEncontradas, err := BuscarCoincidencias(`\[\[.*?\]\]`, "ejemplos/recursos/archivoaRemplazar.json")
		fmt.Println("Patron a buscar ", pat)
		palabrasEncontradas, err := BuscarCoincidencias(pat, inputFile)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Palabras encontradas:")
		for palabra := range palabrasEncontradas {
			fmt.Println(palabra)
		}
		EscribirMapaEnArchivo(palabrasEncontradas, output, separated)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.PersistentFlags().StringP("input", "i", "", "Archivo plantilla a analizar normalmente es el archivo final donde estan los parametros a remplazar")
	generateCmd.PersistentFlags().StringP("pat", "p", "", `"patron a usar ejemplo \[\[.*?\]\]"`)
	generateCmd.PersistentFlags().StringP("separated", "s", "", "en el archivo leyenda puede usar separadores para datos y valor normalmente es =")
	generateCmd.PersistentFlags().StringP("output", "o", "", "nombre del archivo final que tendra los valores encontrado en la plantilla")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
