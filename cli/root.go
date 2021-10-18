package cli

import (
	"github.com/spf13/cobra"
)

// RootCmd main cobra command
var RootCmd = &cobra.Command{
	Use:   "anime-cli",
	Short: "A cli to browse and watch anime.",
}
