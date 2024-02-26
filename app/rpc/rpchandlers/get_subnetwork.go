package rpchandlers

import (
	"github.com/seracoin/serad/app/appmessage"
	"github.com/seracoin/serad/app/rpc/rpccontext"
	"github.com/seracoin/serad/infrastructure/network/netadapter/router"
)

// HandleGetSubnetwork handles the respectively named RPC command
func HandleGetSubnetwork(context *rpccontext.Context, _ *router.Router, request appmessage.Message) (appmessage.Message, error) {
	response := &appmessage.GetSubnetworkResponseMessage{}
	response.Error = appmessage.RPCErrorf("not implemented")
	return response, nil
}
