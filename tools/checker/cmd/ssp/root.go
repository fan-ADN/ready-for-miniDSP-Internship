package ssp

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	host    = ""
	port    = 0
	floor   = 30
	verbose = false
)

func NewRootCmd(args []string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "ssp",
		Short: "SSP Tools",
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("No Command")
		},
	}

	rootCmd.SetArgs(args)
	rootCmd.AddCommand(newCheckCmd())
	rootCmd.AddCommand(newFinalCheckCmd())

	rootCmd.PersistentFlags().StringVar(&host, "host", "localhost", "Host Name")
	rootCmd.PersistentFlags().IntVar(&port, "port", 8080, "Port Number")
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", true, "Verbose mode")

	return rootCmd
}

func newCheckCmd() *cobra.Command {
	checkCmd := &cobra.Command{
		Use:   "check",
		Short: "1回AdRequestを指定したURLに投げる。",
		Run: func(cmd *cobra.Command, args []string) {
			runBidRequestOnce(host, port, floor)
			fmt.Printf("Target URL: %s:%d\n", host, port)
		},
	}
	checkCmd.PersistentFlags().IntVar(&floor, "floor", 30, "Floor Price")

	return checkCmd
}

func newFinalCheckCmd() *cobra.Command {
	finalCheckCmd := &cobra.Command{
		Use:   "final",
		Short: "AdRequestを投げる",
		Run: func(cmd *cobra.Command, args []string) {
			runBidRequestFinal(host, port, floor)
			fmt.Printf("Target URL: %s:%d\n", host, port)
		},
	}
	finalCheckCmd.PersistentFlags().IntVar(&floor, "floor", 30, "Floor Price")

	return finalCheckCmd
}
