package keeper

import (
	"context"

	"github.com/ashraf-mohey/cache/x/ehr/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"encoding/base64"
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
)

func (k msgServer) AddEhr(goCtx context.Context, msg *types.MsgAddEhr) (*types.MsgAddEhrResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	// Sign data using patient private key
	privateKeyPem, err := loadKeyFromFile(msg.PatientPrivateKey)
	if err != nil || privateKeyPem == "" {
		return nil, err
	}
	privateKey, err := privateKeyPemStringToRsaKey(privateKeyPem)
	if err != nil {
		return nil, err
	}
	signatureBytes, err := signData(privateKey, []byte(msg.DataHash))
	if err != nil {
		return nil, err
	}
	patientSignature := base64.StdEncoding.EncodeToString(signatureBytes)

	// Sign data using organization private key
	privateKeyPem, err = loadKeyFromFile(types.OrganizationPrivateKeyPath)
	if err != nil || privateKeyPem == "" {
		return nil, err
	}
	privateKey, err = privateKeyPemStringToRsaKey(privateKeyPem)
	if err != nil {
		return nil, err
	}
	signatureBytes, err = signData(privateKey, []byte(msg.DataHash))
	if err != nil {
		return nil, err
	}
	organizationSignature := base64.StdEncoding.EncodeToString(signatureBytes)

	// Encapsulate ehr fields as an object Ehr
	var ehr = types.Ehr{
		Creator:               types.OrganizationAddress,
		PatientId:             msg.PatientId,
		DataHash:              msg.DataHash,
		PatientSignature:      patientSignature,
		OrganizationSignature: organizationSignature,
		Transferred:           false,
	}
	// Add an ehr to the blockchain and get back the ID
	id := k.AppendEhr(ctx, ehr)
	ehr.Id = id
	// Return the ID of the patient's metadata

	return &types.MsgAddEhrResponse{Ehr: &ehr}, nil
}

func (k Keeper) GetNextEhrId(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.EhrIdKey))
	// Convert the EhrIdKey to bytes
	byteKey := []byte(types.EhrIdKey)
	// Get the value of the id
	bz := store.Get(byteKey)
	// Return one if the id value is not found for first record
	if bz == nil {
		return 1
	}
	// Convert the id into a uint64
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetNextEhrId(ctx sdk.Context, nextId uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.EhrIdKey))
	// Convert the EhrIdKey to bytes
	byteKey := []byte(types.EhrIdKey)
	// Convert id from uint64 to string and get bytes
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, nextId)
	// Set the value of EhrIdKey to nextId
	store.Set(byteKey, bz)
}

func (k Keeper) AppendEhr(ctx sdk.Context, ehr types.Ehr) uint64 {
	// Get the next ehr id in the store
	id := k.GetNextEhrId(ctx)
	ehr.Id = id
	// Get the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.EhrKey))
	// Convert the ehr ID into bytes
	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, ehr.Id)
	// Marshal the ehr into bytes
	appendedValue := k.cdc.MustMarshal(&ehr)
	// Insert the ehr bytes using ehr ID as a key
	store.Set(byteKey, appendedValue)
	// Update the next ehr id
	k.SetNextEhrId(ctx, id+1)
	return id
}
