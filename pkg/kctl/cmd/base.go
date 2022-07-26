package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

func RunCmd() {
	cmd := &cobra.Command{
		Use:          "kctl",
		Short:        "端博客",
		Example:      "kctl",
		SilenceUsage: true,
	}
	cmd.AddCommand(VersionCMD)
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
