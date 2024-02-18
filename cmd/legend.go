/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/culturadevops/gu/file"
	"github.com/spf13/cobra"
)

var boolFlagValue bool

func validarArchivo(rutaArchivo string, separador string) (bool, error) {
	// Abrir el archivo
	archivo, err := os.Open(rutaArchivo)
	if err != nil {
		return false, err
	}
	defer archivo.Close()

	// Crear un scanner para leer el archivo línea por línea
	scanner := bufio.NewScanner(archivo)
	lineaanterior := ""
	// Iterar sobre cada línea del archivo
	for scanner.Scan() {
		linea := strings.TrimSpace(scanner.Text())

		// Verificar si la línea contiene al menos un signo "="
		if !strings.Contains(linea, separador) {
			fmt.Println("linea anterior al error", lineaanterior)
			return false, nil
		}
		lineaanterior = linea
	}

	// Verificar errores durante el escaneo del archivo
	if err := scanner.Err(); err != nil {
		return false, err
	}

	// Si todas las líneas tienen al menos un signo "=", retornar verdadero
	return true, nil
}
func ReadLegendFile(archivo string, separador string) (map[string]string, error) {
	// Abre el archivo

	file, err := os.Open(archivo)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Crea un mapa para almacenar los datos
	datos := make(map[string]string)

	// Crea un scanner para leer el archivo línea por línea
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// Divide cada línea del archivo en nombre y valor usando "=" como delimitador
		partes := strings.Split(scanner.Text(), separador)
		if len(partes) == 2 {
			// Agrega la entrada al mapa
			datos[strings.TrimSpace(partes[0])] = strings.TrimSpace(partes[1])
		}
	}

	// Verifica errores durante el escaneo del archivo
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return datos, nil
}

// legendCmd represents the legend command
var legendCmd = &cobra.Command{
	Use:   "legend",
	Short: "para formato leyenda",
	Long: `toma la direccion del archivo leyenda y lo remplaza en el target tambien usa el separador un comando aceptado podria ser
	go run main.go legend -l recursos/base.vars -t recursos/archivoaRemplazar.json -s : -o archivonuevo.json
	go run main.go legend -l recursos/base.vars -t recursos/archivoaRemplazar.json -s : -r
	go run main.go legend -l recursos/base.vars -t recursos/archivoaRemplazar.json -s : 
	
	`,
	Run: func(cmd *cobra.Command, args []string) {
		legend, _ := cmd.Flags().GetString("legend")
		target, _ := cmd.Flags().GetString("target")
		separated, _ := cmd.Flags().GetString("separated")
		output, _ := cmd.Flags().GetString("output")

		if legend == "" {
			fmt.Println("Falta el archivo legend")
			return
		}
		if target == "" {
			fmt.Println("Falta el archivo target")
			return
		}

		if boolFlagValue && output != "" {
			fmt.Println("no puede agregar un output y remplazar el archivo debe escoger uno solo")
			return
		}
		j := file.New()
		// Imprime el mapa+
		if separated == "" {
			fmt.Println("Falta el separador agrega -s = o -s :")
			return
		}

		fileok, err := validarArchivo(legend, separated)
		if err != nil {
			fmt.Println("Error en el proceso de validacion del archivo legend:", err)
			return
		}
		if !fileok {
			fmt.Println("Error en el archivo legend Falta el separador en el documento")
			return
		}
		mapaDatos, err := ReadLegendFile(legend, separated)
		if err != nil {
			fmt.Println("Error al leer el archivo:", err)
			return
		}
		/*fmt.Println("Mapa de datos:")
		for clave, valor := range mapaDatos {
			fmt.Printf("%s: %s\n", clave, valor)
		}
		*/

		if boolFlagValue {
			fmt.Println("se va a remplazar el archivo:", target)
			fmt.Println(j.ReplaceContentAndCreateNewFile(target, target, mapaDatos))
		} else {
			if output != "" {
				fmt.Println("se va a crear el archivo:", output)
				fmt.Println(j.ReplaceContentAndCreateNewFile(target, output, mapaDatos))
			} else {
				x, err := j.ReplaceTextInFile(target, mapaDatos)
				if err != nil {
					print(err)
					return
				}

				fmt.Println(x)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(legendCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	legendCmd.PersistentFlags().StringP("legend", "l", "", "Archivo leyenda donde estan todos los datos que seran remplazado")
	legendCmd.PersistentFlags().StringP("target", "t", "", "archivo plantilla donde se van a remplazar todo o se usara para crear un nuevo archivo con los datos remplazados")
	legendCmd.PersistentFlags().StringP("separated", "s", "", "en el archivo leyenda puede usar separadores para datos y valor normalmente es =")
	legendCmd.PersistentFlags().StringP("output", "o", "", "nombre del archivo final que tendra tanto los valores de leyenda y las plantillas")
	legendCmd.Flags().BoolVarP(&boolFlagValue, "remplace", "r", false, "valor bool que define si el target sera remplazado o el resultado se mostraran en la pantalla ")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// legendCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
