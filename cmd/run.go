/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"log/slog"

	"github.com/spf13/cobra"
	"go.uber.org/fx"

	"github.com/matt-e/seed/pkg/stage"
)

// runCmd represents the run command.
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(_ *cobra.Command, args []string) error {
		return cobraFxAdapter(fx.Invoke(func(ctx context.Context, s stage.Stage, logger *slog.Logger) {
			logger.InfoContext(ctx, "Hello, world", "stage", s.String(), "args", args)
		}))
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
