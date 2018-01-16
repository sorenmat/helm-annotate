package main

import (
	"os"

	"github.com/spf13/cobra"
	"k8s.io/helm/pkg/helm"
	helm_env "k8s.io/helm/pkg/helm/environment"
)

var (
	settings helm_env.EnvSettings
)

func setupConnection(c *cobra.Command, args []string) error {
	settings.TillerHost = os.Getenv("TILLER_HOST")

	return nil
}

func ensureHelmClient(h helm.Interface) helm.Interface {
	if h != nil {
		return h
	}
	return helm.NewClient(helm.Host(settings.TillerHost))
}

func main() {
	var rootCmd = &cobra.Command{Use: "annotate"}
	rootCmd.AddCommand(newSetCmd())
	rootCmd.AddCommand(newGetCmd())
	rootCmd.Execute()
}
