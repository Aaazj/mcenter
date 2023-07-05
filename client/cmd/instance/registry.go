package instance

import (
	"fmt"

	"github.com/Aaazj/mcenter/apps/instance"
	"github.com/Aaazj/mcenter/client/rest"
	"github.com/spf13/cobra"
)

// Cmd represents the start command
var registry = &cobra.Command{
	Use:   "registry",
	Short: "服务实例注册",
	Long:  "服务实例注册",
	RunE: func(cmd *cobra.Command, args []string) error {
		req := instance.NewRegistryRequest()
		ins, err := rest.C().Instance().RegistryInstance(cmd.Context(), req)
		if err != nil {
			return err
		}
		fmt.Println(ins)
		return nil
	},
}

func init() {
	Cmd.AddCommand(registry)
}
