//go:generate go-enum --marshal --noprefix --nocase --names
package stage

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

var Module = fx.Module("stage",
	fx.Provide(
		Get,
	),
)

// ENUM(test, dev, staging, prod).
type Stage int

const Default = Dev

var stageStr string

func BindStageFlag(rootCmd *cobra.Command) {
	rootCmd.PersistentFlags().StringVar(&stageStr, "stage", Dev.String(), "Application runtime stages. Valid values are: "+strings.Join(StageNames(), ". "))
	if err := viper.BindPFlag("stage", rootCmd.PersistentFlags().Lookup("stage")); err != nil {
		panic(err)
	}
}

func SetTestStage() {
	stageStr = Test.String()
}

func Get() (Stage, error) {
	if stageStr == "" {
		return Default, nil
	}
	return ParseStage(stageStr)
}

func MustGet() Stage {
	s, err := Get()
	if err != nil {
		panic(err)
	}
	return s
}
