package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/kwilteam/kuneiform/kfparser"
	"github.com/kwilteam/kuneiform/version"
)

const (
	CodeOk = iota
	CodeFail
	CodeNoInputFile
	CodeErrMakeOutputDir
	CodeErrReadFile
	CodeErrParseFile
	CodeErrSerialize
	CodeErrWriteFile
)

var (
	showVersion bool
	trace       bool
	outputDir   string = "./build"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kuneiform <file>",
	Short: "kuneiform command line parser",
	Long:  `Given a kuneiform file, outputs a JSON file that can be used to deployed to the Kwil.`,
	//Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if showVersion {
			cmd.Println(version.Version)
			return
		}

		if len(args) != 1 {
			cmd.Println("Please provide a file to parse")
			os.Exit(CodeNoInputFile)
		}

		if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
			cmd.Println(err)
			os.Exit(CodeErrMakeOutputDir)
		}

		fileToParse := args[0]
		fileName := filepath.Base(fileToParse)
		fileName = fileName[:len(fileName)-len(filepath.Ext(fileName))]
		data, err := os.ReadFile(fileToParse)
		if err != nil {
			cmd.Println(err)
			os.Exit(CodeErrReadFile)
		}

		parserMode := kfparser.Default
		if trace {
			parserMode = kfparser.Trace
		}
		db, err := kfparser.ParseKF(string(data), nil, parserMode)
		if err != nil {
			cmd.Println(err)
			os.Exit(CodeErrParseFile)
		}

		dbJson, err := db.ToJSON()
		if err != nil {
			cmd.Println(err)
			os.Exit(CodeErrSerialize)
		}

		outputFilePath := filepath.Join(outputDir, fileName+".json")
		err = writeToFile(dbJson, outputFilePath)
		if err != nil {
			cmd.Println(err)
			os.Exit(CodeErrWriteFile)
		}

		cmd.Println("Successfully parsed", fileToParse, "to", outputFilePath)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(CodeFail)
	}
}

func init() {
	rootCmd.Flags().BoolP("trace", "t", false, "toggle trace output")
	rootCmd.Flags().BoolVarP(&showVersion, "version", "v", false, "Show version and exit")
	rootCmd.Flags().StringVarP(&outputDir, "output", "o", outputDir, "Output directory")
}

func writeToFile(data []byte, filePath string) (err error) {
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer func() {
		err = f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	if _, err := f.Write(data); err != nil {
		return err
	}

	return err
}

func main() {
	Execute()
}
