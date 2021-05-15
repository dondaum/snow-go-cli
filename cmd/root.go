package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "snowgocli",
	Short: "snowgocli is a leight, fast, concurrent snowflake query monitoring tool.",
	Long:  "A Fast and Flexible concurrent snowflake query monitoring tool",
	Run: func(cmd *cobra.Command, args []string) {
		// Root logic
		currentTime := time.Now()
		fmt.Println("Oh hi there! The current time is: ", currentTime.String())
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
