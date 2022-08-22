package keeper

import (
	"context"

	"github.com/ashraf-mohey/cache/x/ehr/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	
	"github.com/cosmos/cosmos-sdk/store/prefix"
)

func (k Keeper) GetPatientPendingTransferEhrs(goCtx context.Context, req *types.QueryGetPatientPendingTransferEhrsRequest) (*types.QueryGetPatientPendingTransferEhrsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx
	
	var ehrs []*types.Ehr
	store := ctx.KVStore(k.storeKey)
	// Get the part of the store that keeps ehrs
	ehrsStore := prefix.NewStore(store, []byte(types.EhrKey))

	iter := sdk.KVStorePrefixIterator(ehrsStore, []byte{}) //sdk.KVStorePrefixIterator(store, []byte(types.EhrKey))
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var ehr types.Ehr
		if err := k.cdc.Unmarshal(iter.Value(), &ehr); err != nil {
			continue
		}
		if !ehr.Transferred && ehr.PatientId == req.PatientId {
			ehrs = append(ehrs, &ehr)
		}
	}

	return &types.QueryGetPatientPendingTransferEhrsResponse{Ehrs: ehrs}, nil
}
