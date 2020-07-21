package cui

import (
	"github.com/spf13/cobra"
)

var appVersion = "unknown"

// VersionCmd represents the version command.
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the app's version",
	Long: `Prints the app's version. This app uses semver and the version value
is automatically generated while building.`,
	Run: func(cmd *cobra.Command, args []string) {
		Info(appVersion)
	},
}
