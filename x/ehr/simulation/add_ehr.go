package simulation

import (
	"math/rand"

	"github.com/ashraf-mohey/cache/x/ehr/keeper"
	"github.com/ashraf-mohey/cache/x/ehr/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgAddEhr(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgAddEhr{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the AddEhr simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "AddEhr simulation not implemented"), nil, nil
	}
}
