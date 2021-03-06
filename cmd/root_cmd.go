package cmd

import (
	"errors"
	"fmt"
	"github.com/sahara-go/thinkgo/cmd/migrate_cmd"
	"github.com/sahara-go/thinkgo/cmd/server_cmd"
	"github.com/sahara-go/thinkgo/log"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:               "github.com/sahara-go/thinkgo",
	Short:             "-v",
	SilenceUsage:      true,
	DisableAutoGenTag: true,
	Long:              `thinkgo`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			log.Error("requires at least one arg")
			return errors.New("requires at least one arg")
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		usageStr := `欢迎使用thinkgo，可以是用 -h 查看命令`
		log.Infof("%s\n", usageStr)
	},
}

func init() {
	rootCmd.AddCommand(server_cmd.StartCmd)
	rootCmd.AddCommand(migrate_cmd.StartCmd)
}

//Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(fmt.Errorf("cmd命令执行发生错误:%v", err.Error()))
	}
}
