package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "E-Commerce",
	Short: "Awesome E-Commerce",
	Long: `
Directory Structure
1. API
-- CORE
-- AUTH
-- ORDER
-- PRODUCT
2. ADMIN
3. WEB
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello World")
	},
}

// Execute the Root Command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
