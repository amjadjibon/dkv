package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/amjadjibon/dkv/client"
)

var kvGetCmd = &cobra.Command{
	Use:   "get",
	Short: "kv get command",
	Long:  `kv get command`,
	Run: func(cmd *cobra.Command, args []string) {
		var client = client.New(host, port)
		var key = KVKey
		var value, err = client.GetValue(context.Background(), key)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(value))
	},
}

var kvSetCmd = &cobra.Command{
	Use:   "set",
	Short: "kv set command",
	Long:  `kv set command`,
	Run: func(cmd *cobra.Command, args []string) {
		var client = client.New(host, port)
		var key = KVKey
		var value = []byte(KVValue)
		var err = client.SetValue(context.Background(), key, value)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("successful")
	},
}

var kvDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "kv delete command",
	Long:  `kv delete command`,
	Run: func(cmd *cobra.Command, args []string) {
		var client = client.New(host, port)
		var key = KVKey
		var err = client.DeleteValue(context.Background(), key)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("successful")
	},
}

func init() {
	rootCmd.AddCommand(kvDeleteCmd, kvSetCmd, kvGetCmd)
	kvSetCmd.Flags().StringVar(&host, "host", "", "Ex: localhost")
	kvSetCmd.Flags().StringVar(&port, "port", "", "Ex: 6001")
	kvSetCmd.Flags().StringVar(&KVKey, "key", "", "")
	kvSetCmd.Flags().StringVar(&KVValue, "value", "", "")
	_ = kvSetCmd.MarkFlagRequired("key")
	_ = kvSetCmd.MarkFlagRequired("value")
	_ = kvSetCmd.MarkFlagRequired("host")
	_ = kvSetCmd.MarkFlagRequired("port")

	kvGetCmd.Flags().StringVar(&host, "host", "", "Ex: localhost")
	kvGetCmd.Flags().StringVar(&port, "port", "", "Ex: 6001")
	kvGetCmd.Flags().StringVar(&KVKey, "key", "", "")
	_ = kvGetCmd.MarkFlagRequired("key")
	_ = kvGetCmd.MarkFlagRequired("host")
	_ = kvGetCmd.MarkFlagRequired("port")

	kvDeleteCmd.Flags().StringVar(&host, "host", "", "Ex: localhost")
	kvDeleteCmd.Flags().StringVar(&port, "port", "", "Ex: 6001")
	kvDeleteCmd.Flags().StringVar(&KVKey, "key", "", "")
	_ = kvDeleteCmd.MarkFlagRequired("key")
	_ = kvDeleteCmd.MarkFlagRequired("host")
	_ = kvDeleteCmd.MarkFlagRequired("port")
}
