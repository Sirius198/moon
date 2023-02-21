package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"moon/x/ibank/types"
)

var _ = strconv.Itoa(0)

func CmdShowIncoming() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-incoming [receiver] [pending]",
		Short: "Query show-incoming",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqReceiver := args[0]
			reqPending := args[1]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryShowIncomingRequest{

				Receiver: reqReceiver,
				Pending:  reqPending,
			}

			res, err := queryClient.ShowIncoming(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
