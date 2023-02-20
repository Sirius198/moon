package cli

import (
	"strconv"

	"moon/x/ibank/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdReceive() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "receive [transaction-id]",
		Short: "Broadcast message receive",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTransactionId := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			txnId, err := strconv.ParseInt(argTransactionId, 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgReceive(
				clientCtx.GetFromAddress().String(),
				txnId,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
