package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/fan-ADN/ready-for-miniDSP-Internship/tools/checker/cmd/dsp"
	"github.com/fan-ADN/ready-for-miniDSP-Internship/tools/checker/cmd/ml"
	"github.com/fan-ADN/ready-for-miniDSP-Internship/tools/checker/cmd/ssp"
	"github.com/fan-ADN/ready-for-miniDSP-Internship/tools/checker/config"
)

func newRootCmd(args []string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "checker",
		Short: "DSP for internship Checker tools",
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("No command")
		},
	}

	rootCmd.SetArgs(args)

	rootCmd.AddCommand(newVersionCmd())
	rootCmd.AddCommand(ssp.NewRootCmd(args))
	rootCmd.AddCommand(ml.NewRootCmd(args))
	rootCmd.AddCommand(dsp.NewRootCmd(args))

	return rootCmd
}

func Execute(args []string) int {
	defer func() {
		r := recover()
		if r != nil {
			return
		}
	}()

	if err := newRootCmd(args).Execute(); err != nil {
		return 255
	}
	return 0
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
}

func newVersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print Version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("intern-dsp version: %s, revision %s\n", config.Version, config.Revision)
		},
	}

	return cmd
}
