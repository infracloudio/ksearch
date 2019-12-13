package cmd

import (

	// Load all known auth plugins
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	"os"

	"github.com/infracloudio/ksearch/pkg/printers"
	"github.com/infracloudio/ksearch/pkg/util"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/client/config"

	"github.com/spf13/cobra"
)

var (
	resName, namespace, kinds string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "ksearch",
	Short:   "run ksearch --help to get the usage",
	Version: "v0.0.1",
	Long:    `ksearch is a command line tool to search for a given pattern in a Kubernetes cluster and will print all of the available resources in a cluster if none is provided`,
	Run: func(cmd *cobra.Command, args []string) {
		getter := make(chan interface{})

		cfg := config.GetConfigOrDie()
		clientset := kubernetes.NewForConfigOrDie(cfg)

		go util.Getter(namespace, clientset, kinds, getter)

		for resource := range getter {
			printers.Printer(resource, resName)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&resName, "pattern", "p", "", "pattern you want to search for")
	rootCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "namespace you want to search in")
	rootCmd.PersistentFlags().StringVarP(&kinds, "kinds", "k", "", "comma separated list of all the kinds that you want to include")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
}
