package keeper

import (
	"errors"

	"github.com/ashraf-mohey/cache/x/ehr/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/v2/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v2/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v2/modules/core/24-host"

	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
)

// TransmitIbcTransferEhrsPacket transmits the packet over IBC with the specified source port and source channel
func (k Keeper) TransmitIbcTransferEhrsPacket(
	ctx sdk.Context,
	packetData types.IbcTransferEhrsPacketData,
	sourcePort,
	sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
) error {

	sourceChannelEnd, found := k.ChannelKeeper.GetChannel(ctx, sourcePort, sourceChannel)
	if !found {
		return sdkerrors.Wrapf(channeltypes.ErrChannelNotFound, "port ID (%s) channel ID (%s)", sourcePort, sourceChannel)
	}

	destinationPort := sourceChannelEnd.GetCounterparty().GetPortID()
	destinationChannel := sourceChannelEnd.GetCounterparty().GetChannelID()

	// get the next sequence
	sequence, found := k.ChannelKeeper.GetNextSequenceSend(ctx, sourcePort, sourceChannel)
	if !found {
		return sdkerrors.Wrapf(
			channeltypes.ErrSequenceSendNotFound,
			"source port: %s, source channel: %s", sourcePort, sourceChannel,
		)
	}

	channelCap, ok := k.ScopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(sourcePort, sourceChannel))
	if !ok {
		return sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	packetBytes, err := packetData.GetBytes()
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, "cannot marshal the packet: "+err.Error())
	}

	packet := channeltypes.NewPacket(
		packetBytes,
		sequence,
		sourcePort,
		sourceChannel,
		destinationPort,
		destinationChannel,
		timeoutHeight,
		timeoutTimestamp,
	)

	if err := k.ChannelKeeper.SendPacket(ctx, channelCap, packet); err != nil {
		return err
	}

	return nil
}

// OnRecvIbcTransferEhrsPacket processes packet reception
func (k Keeper) OnRecvIbcTransferEhrsPacket(ctx sdk.Context, packet channeltypes.Packet, data types.IbcTransferEhrsPacketData) (packetAck types.IbcTransferEhrsPacketAck, err error) {
	// validate packet data upon receiving
	if err := data.ValidateBasic(); err != nil {
		return packetAck, err
	}

	// TODO: packet reception logic

	return packetAck, nil
}

// OnAcknowledgementIbcTransferEhrsPacket responds to the the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementIbcTransferEhrsPacket(ctx sdk.Context, packet channeltypes.Packet, data types.IbcTransferEhrsPacketData, ack channeltypes.Acknowledgement) error {
	switch dispatchedAck := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:

		// TODO: failed acknowledgement logic
		_ = dispatchedAck.Error

		return nil
	case *channeltypes.Acknowledgement_Result:
		// Decode the packet acknowledgment
		var packetAck types.IbcTransferEhrsPacketAck

		if err := types.ModuleCdc.UnmarshalJSON(dispatchedAck.Result, &packetAck); err != nil {
			// The counter-party module doesn't implement the correct acknowledgment format
			return errors.New("cannot unmarshal acknowledgment")
		}

		// TODO: successful acknowledgement logic
		store := ctx.KVStore(k.storeKey)
		ehrsStore := prefix.NewStore(store, []byte(types.EhrKey))

		for _, id := range packetAck.Ids {
			_ = id
			var ehr types.Ehr
			byteKey := make([]byte, 8)
			binary.BigEndian.PutUint64(byteKey, id)
			bz := ehrsStore.Get(byteKey)
			k.cdc.MustUnmarshal(bz, &ehr)
			if ehr.Id < 1 {
				continue
			}

			ehr.Transferred = true
			// Marshal the patient into bytes
			updatedValue := k.cdc.MustMarshal(&ehr)
			// Insert the update ehr bytes using ehr ID as a key
			ehrsStore.Set(byteKey, updatedValue)
		}

		return nil
	default:
		// The counter-party module doesn't implement the correct acknowledgment format
		return errors.New("invalid acknowledgment format")
	}
}

// OnTimeoutIbcTransferEhrsPacket responds to the case where a packet has not been transmitted because of a timeout
func (k Keeper) OnTimeoutIbcTransferEhrsPacket(ctx sdk.Context, packet channeltypes.Packet, data types.IbcTransferEhrsPacketData) error {

	// TODO: packet timeout logic

	return nil
}
