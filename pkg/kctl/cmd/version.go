package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

type VersionInfo struct {
	Version   string
	GoVersion string
}

func NewVersionInfo() *VersionInfo {
	return &VersionInfo{
		Version:   "v0.1",
		GoVersion: "go1.17.3",
	}
}

func (this *VersionInfo) String() string {
	return fmt.Sprintf("version info: {version: %s, goversion: %s}", this.Version, this.GoVersion)
}

var VersionCMD = &cobra.Command{
	Use:          "version",
	Short:        "v",
	Example:      "kctl version",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("%s \n", NewVersionInfo())
		return nil
	},
}
