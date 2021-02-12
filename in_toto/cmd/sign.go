package cmd

import (
	"fmt"
	"os"
	intoto "github.com/boxboat/in-toto-golang/in_toto"
	"github.com/spf13/cobra"
)

var prompt bool
var append bool
var output string
var gpgID string

var signCmd = &cobra.Command{
	Use:   "sign",
	Short: "Sign will sign a link or a layout",
	Run: signLayout,
}

func init() {
	rootCmd.AddCommand(signCmd)
	signCmd.PersistentFlags().StringVarP(&keyPath,
		"key", "k", "",
		`Path to a PEM formatted private key file used to sign
the resulting layout metadata. (passing one of '--key'
or '--gpg' is required) `)
}

func signLayout(cmd *cobra.Command, args []string) {
	var key intoto.Key

	if err := key.LoadKey(keyPath, "rsassa-pss-sha256", []string{"sha256", "sha512"}); err != nil {
		fmt.Println("Invalid Key Error:", err.Error())
		os.Exit(1)
	}
}
