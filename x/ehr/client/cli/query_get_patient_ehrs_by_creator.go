package cli

import (
	"strconv"

	"github.com/ashraf-mohey/cache/x/ehr/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdGetPatientEhrsByCreator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get-patient-ehrs-by-creator [patient-id] [creator]",
		Short: "Query getPatientEhrsByCreator",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqPatientId, err := strconv.ParseUint(string(args[0]), 10, 64)
			reqCreator := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetPatientEhrsByCreatorRequest{

				PatientId: reqPatientId,
				Creator:   reqCreator,
			}

			res, err := queryClient.GetPatientEhrsByCreator(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
