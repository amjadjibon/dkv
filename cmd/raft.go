package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/amjadjibon/dkv/client"
)

var raftJoinCmd = &cobra.Command{
	Use:   "join",
	Short: "join raft server",
	Long:  `join raft server`,
	Run: func(cmd *cobra.Command, args []string) {
		var client = client.New(host, port)
		var err = client.RaftJoin(context.Background(), nodeID, nodeAddress)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("raft join successful")
	},
}

var raftLeaveCmd = &cobra.Command{
	Use:   "leave",
	Short: "leave raft server",
	Long:  `leave raft server`,
	Run: func(cmd *cobra.Command, args []string) {
		var client = client.New(host, port)
		var err = client.RaftLeave(context.Background(), nodeID)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("raft join successful")
	},
}

var raftStateCmd = &cobra.Command{
	Use:   "state",
	Short: "join raft server",
	Long:  `join raft server`,
	Run: func(cmd *cobra.Command, args []string) {
		var client = client.New(host, port)
		var res, err = client.RaftState(context.Background())
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("raft state: ", res)
	},
}

func init() {
	rootCmd.AddCommand(raftStateCmd, raftLeaveCmd, raftJoinCmd)

	raftJoinCmd.Flags().StringVar(&host, "host", "", "Ex: localhost")
	raftJoinCmd.Flags().StringVar(&port, "port", "", "Ex: 6001")
	raftJoinCmd.Flags().StringVar(&nodeID, "node-id", "", "Ex: node_id_1")
	raftJoinCmd.Flags().StringVar(&nodeAddress, "node-address", "", "Ex: localhost:6001")
	_ = raftJoinCmd.MarkFlagRequired("host")
	_ = raftJoinCmd.MarkFlagRequired("port")
	_ = raftJoinCmd.MarkFlagRequired("node-id")
	_ = raftJoinCmd.MarkFlagRequired("node-address")

	raftLeaveCmd.Flags().StringVar(&host, "host", "", "Ex: localhost")
	raftLeaveCmd.Flags().StringVar(&port, "port", "", "Ex: 6001")
	raftLeaveCmd.Flags().StringVar(&nodeID, "node-id", "", "Ex: node_id_1")
	_ = raftLeaveCmd.MarkFlagRequired("host")
	_ = raftLeaveCmd.MarkFlagRequired("port")
	_ = raftLeaveCmd.MarkFlagRequired("node-id")

	raftStateCmd.Flags().StringVar(&host, "host", "", "Ex: localhost")
	raftStateCmd.Flags().StringVar(&port, "port", "", "Ex: 6001")
	_ = raftStateCmd.MarkFlagRequired("host")
	_ = raftStateCmd.MarkFlagRequired("port")
}
