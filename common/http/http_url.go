package http

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"nacosctl/common"
	"nacosctl/common/constant"
)

func GetNamespaceUrl(cmd *cobra.Command) string {
	return getServerAddressByViper(cmd) + constant.NamespaceUrl + "?accessToken=" + viper.GetString("nacosctl.accessToken")
}

func GetConfigUrl(cmd *cobra.Command) string {
	return getServerAddressByViper(cmd) + constant.ConfigUrl + "?accessToken=" + viper.GetString("nacosctl.accessToken")
}

func GetConfigVersionUrl(cmd *cobra.Command) string {
	return getServerAddressByViper(cmd) + constant.VersionUrl + "?accessToken=" + viper.GetString("nacosctl.accessToken")
}

func GetUserUrl(cmd *cobra.Command) string {
	return getServerAddressByViper(cmd) + constant.UserUrl + "?accessToken=" + viper.GetString("nacosctl.accessToken")
}

func GetLoginUrl(cmd *cobra.Command) string {
	return getServerAddressByViper(cmd) + constant.LoginUrl
}

func getServerAddressByViper(cmd *cobra.Command) string {
	address := common.GetServerAddress(cmd)
	prefix := fmt.Sprintf(constant.Prefix, address)
	return prefix
}
