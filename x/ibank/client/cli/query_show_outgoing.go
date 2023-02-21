package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"moon/x/ibank/types"
)

var _ = strconv.Itoa(0)

func CmdShowOutgoing() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-outgoing [sender] [pending]",
		Short: "Query show-outgoing",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqSender := args[0]
			reqPending := args[1]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryShowOutgoingRequest{

				Sender:  reqSender,
				Pending: reqPending,
			}

			res, err := queryClient.ShowOutgoing(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
