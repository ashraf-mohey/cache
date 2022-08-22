package keeper

import (
	"context"

	"github.com/ashraf-mohey/cache/x/ehr/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v2/modules/core/02-client/types"

	//"strconv"
)



func (k msgServer) SendIbcTransferEhrs(goCtx context.Context, msg *types.MsgSendIbcTransferEhrs) (*types.MsgSendIbcTransferEhrsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: logic before transmitting the packet

	// Construct the packet
	var packet types.IbcTransferEhrsPacketData
	packet.Creator = msg.Creator
	packet.OrganizationAddress = types.OrganizationAddress
	packet.PatientId = msg.PatientId
	packet.PendingTransferUrl = types.PendingTransferUrl //types.PendingTransferUrl + strconv.FormatUint(msg.PatientId, 10) + "/" + "accessToken"
	

	// Transmit the packet
	err := k.TransmitIbcTransferEhrsPacket(
		ctx,
		packet,
		msg.Port,
		msg.ChannelID,
		clienttypes.ZeroHeight(),
		msg.TimeoutTimestamp,
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgSendIbcTransferEhrsResponse{}, nil
}


