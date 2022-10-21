/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// dirCmd represents the dir command
var dirCmd = &cobra.Command{
	Use:   "dir",
	Short: "Crear las carpetas principales para iniciar una ataque de hacking",
	Long: `Creamos las carpetas:
		- Content
		- Exploit
		- nmap`,
	Run: func(cmd *cobra.Command, args []string) {
		dir()
	},
}

func init() {
	rootCmd.AddCommand(dirCmd)

}

func dir() {

	directorys := [3]string{"content", "exploit", "nmap"}
	for i := 0; i < len(directorys); i++ {
		err := os.Mkdir(directorys[i], 0755)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf(directorys[i] + " se ha creado. \n")
	}
}
