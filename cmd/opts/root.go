package opts

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "connect-idp",
	Short:   "Connect IdP",
	Long:    "Connect IdP - An identity provider for Connect",
	Version: "0.1.0",
}
