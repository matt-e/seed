/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/fx"

	"github.com/matt-e/otelfx"

	"github.com/matt-e/seed/pkg/log"
	"github.com/matt-e/seed/pkg/stage"
)

const (
	OtelDevEndpoint   = otelfx.CollectorEndpoint("127.0.0.1:4317")
	OtelDevIsInsecure = otelfx.CollectorIsInsecure(true)
)

func cobraFxAdapter(opts ...fx.Option) error {
	ctx := context.Background()

	args := []fx.Option{
		log.Module,
		stage.Module,
		otelfx.Module,
		fx.Provide(func() context.Context { return ctx }),
	}

	switch stage.MustGet() {
	case stage.Test:
		args = append(args, otelfx.StdoutModule)
	case stage.Dev:
		args = append(args,
			fx.Supply(OtelDevEndpoint, OtelDevIsInsecure),
			otelfx.CollectorModule,
		)
	default:
		args = append(args,
			otelfx.CollectorModule,
		)
	}

	args = append(args, opts...)
	app := fx.New(args...)
	app.Run()
	// TODO: handle errors and exit codes
	return nil
}

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "seed",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	stage.BindStageFlag(rootCmd)
}
