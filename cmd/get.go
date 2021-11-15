package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"nacos-cli/config"
	"nacos-cli/config/constant"
	"strconv"
)

var getConfigCmd = &cobra.Command{
	Use:   "get [namespaceId] [dataId] [group]",
	Short: "Get configuration information",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 3 {
			return errors.New("Three parameters are required,[namespaceId] and [dataId] and [group]")
		}
		if args[2] == "" {
			args[2] = constant.DefaultGroup
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		viper := config.GetViper()
		address := viper.Get("server.address").(string)
		port := strconv.Itoa(viper.Get("server.port").(int))

		// 如果命令指定了flag  就替换配置文件
		if cmd.Flag("address").Value.String() != constant.DefaultAddress {
			address = cmd.Flag("address").Value.String()
		}
		if cmd.Flag("port").Value.String() != constant.DefaultPort {
			port = cmd.Flag("port").Value.String()
		}
		config.GetCommand(args[0], args[1], args[2], address, port)
	},
}

func init() {
	getConfigCmd.Flags().StringP("address", "b", constant.DefaultAddress, "nacos server ip address")
	getConfigCmd.Flags().StringP("port", "p", constant.DefaultPort, "nacos server port")
	configCmd.AddCommand(getConfigCmd)
}
