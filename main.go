package main

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	nexproto "github.com/PretendoNetwork/nex-protocols-go"
)

var nexServer *nex.Server

func main() {
	nexServer = nex.NewServer()
	nexServer.SetPrudpVersion(0)
	nexServer.SetSignatureVersion(0)
	nexServer.SetNexVersion(2)
	nexServer.SetKerberosKeySize(32)
	nexServer.SetAccessKey("6181dff1")

	nexServer.On("Data", func(packet *nex.PacketV0) {
		request := packet.RMCRequest()

		fmt.Println("==MK7 - Auth==")
		fmt.Printf("Protocol ID: %#v\n", request.ProtocolID())
		fmt.Printf("Method ID: %#v\n", request.MethodID())
		fmt.Println("===============")
	})

	authenticationServer := nexproto.NewAuthenticationProtocol(nexServer)

	// Handle LoginEx RMC method
	authenticationServer.LoginEx(loginEx)

	// Handle RequestTicket RMC method
	authenticationServer.RequestTicket(requestTicket)

	nexServer.Listen(":60002")
}
