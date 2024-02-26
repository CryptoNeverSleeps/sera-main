package rpchandlers

import (
	"github.com/seracoin/serad/app/appmessage"
	"github.com/seracoin/serad/app/rpc/rpccontext"
	"github.com/seracoin/serad/domain/consensus/utils/constants"
	"github.com/seracoin/serad/infrastructure/network/netadapter/router"
)

// HandleGetCoinSupply handles the respectively named RPC command
func HandleGetCoinSupply(context *rpccontext.Context, _ *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	if !context.Config.UTXOIndex {
		errorMessage := &appmessage.GetCoinSupplyResponseMessage{}
		errorMessage.Error = appmessage.RPCErrorf("Method unavailable when serad is run without --utxoindex")
		return errorMessage, nil
	}

	circulatingSeepSupply, err := context.UTXOIndex.GetCirculatingSeepSupply()
	if err != nil {
		return nil, err
	}

	response := appmessage.NewGetCoinSupplyResponseMessage(
		constants.MaxSeep,
		circulatingSeepSupply,
	)

	return response, nil
}
