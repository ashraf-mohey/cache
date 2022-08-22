package cli

import (
	"strconv"

	"github.com/ashraf-mohey/cache/x/ehr/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdAddEhr() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-ehr [patient-id] [patient-private-key] [data-hash]",
		Short: "Broadcast message addEhr",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPatientId, err := strconv.ParseUint(string(args[0]), 10, 64)
			argPatientPrivateKey := args[1]
			argDataHash := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAddEhr(
				clientCtx.GetFromAddress().String(),
				argPatientId,
				argPatientPrivateKey,
				argDataHash,
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
