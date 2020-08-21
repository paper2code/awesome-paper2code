package cmd

import (
	"os"
	"fmt"

	"github.com/sirupsen/logrus"	
)

var (
	log       = logrus.New()
)

var options struct {
	verbose     bool
	debug       bool
	generateDoc bool
}

var RootCmd = &cobra.Command{
	Use:   "awesome-paper2code",
	Short: "Awesome paper2code generates an awesome list with all trending papers into categories.",
	Long:  `Awesome paper2code generates an awesome list with all trending papers into categories.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	if options.generateDoc {
		err := doc.GenMarkdownTree(RootCmd, "../docs")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func init() {
	flags := RootCmd.PersistentFlags()
	flags.BoolVarP(&options.generateDoc, "generate-doc", "g", false, "generate sub-commands documentation.")
	flags.BoolVarP(&options.debug, "debug", "d", false, "debug mode.")
	flags.BoolVarP(&options.verbose, "verbose", "v", false, "verbose output.")
}
