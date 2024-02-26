package main

import (
	"context"
	"fmt"

	"github.com/seracoin/serad/cmd/serawallet/daemon/client"
	"github.com/seracoin/serad/cmd/serawallet/daemon/pb"
	"github.com/seracoin/serad/cmd/serawallet/utils"
)

func balance(conf *balanceConfig) error {
	daemonClient, tearDown, err := client.Connect(conf.DaemonAddress)
	if err != nil {
		return err
	}
	defer tearDown()

	ctx, cancel := context.WithTimeout(context.Background(), daemonTimeout)
	defer cancel()
	response, err := daemonClient.GetBalance(ctx, &pb.GetBalanceRequest{})
	if err != nil {
		return err
	}

	pendingSuffix := ""
	if response.Pending > 0 {
		pendingSuffix = " (pending)"
	}
	if conf.Verbose {
		pendingSuffix = ""
		println("Address                                                                       Available             Pending")
		println("-----------------------------------------------------------------------------------------------------------")
		for _, addressBalance := range response.AddressBalances {
			fmt.Printf("%s %s %s\n", addressBalance.Address, utils.FormatSdr(addressBalance.Available), utils.FormatSdr(addressBalance.Pending))
		}
		println("-----------------------------------------------------------------------------------------------------------")
		print("                                                 ")
	}
	fmt.Printf("Total balance, SERA %s %s%s\n", utils.FormatSdr(response.Available), utils.FormatSdr(response.Pending), pendingSuffix)

	return nil
}
