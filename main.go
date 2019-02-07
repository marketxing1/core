package main

import (
	"fmt"
	"time"

	"github.com/DOSNetwork/core/configuration"
	"github.com/DOSNetwork/core/dosnode"
	"github.com/DOSNetwork/core/log"
	"github.com/DOSNetwork/core/onchain"
	"github.com/DOSNetwork/core/p2p"
	"github.com/DOSNetwork/core/share/dkg/pedersen"
	"github.com/DOSNetwork/core/suites"
)

// main
func main() {
	//Read Configuration
	config := configuration.Config{}
	if err := config.LoadConfig(); err != nil {
		log.Fatal(err)
	}

	role := config.NodeRole
	port := config.Port
	bootstrapIp := config.BootStrapIp
	chainConfig := config.GetChainConfig()

	//Set up an onchain adapter
	chainConn, err := onchain.AdaptTo(config.GetCurrentType())
	if err != nil {
		log.Fatal(err)
	}

	if err = chainConn.SetAccount(config.GetCredentialPath()); err != nil {
		log.Fatal(err)
	}
	//Init log module with nodeID that is an onchain account address
	log.Init(chainConn.GetId()[:])
	if err = chainConn.Init(chainConfig); err != nil {
		log.Fatal(err)
	}

	rootCredentialPath := "testAccounts/bootCredential/fundKey"
	if err := chainConn.BalanceMaintain(rootCredentialPath); err != nil {
		log.Fatal(err)
	}

	go func() {
		fmt.Println("regular balanceMaintain started")
		ticker := time.NewTicker(time.Hour * 8)
		for range ticker.C {
			if err := chainConn.BalanceMaintain(rootCredentialPath); err != nil {
				log.Fatal(err)
			}
		}
	}()

	//Build a p2p network
	p, err := p2p.CreateP2PNetwork(chainConn.GetId(), port)
	if err != nil {
		log.Fatal(err)
	}

	if err := p.Listen(); err != nil {
		log.Fatal(err)
	}

	//Bootstrapping p2p network
	if role != "BootstrapNode" {
		if err = p.Join(bootstrapIp); err != nil {
			log.Fatal(err)
		}
	}

	//Build a p2pDKG
	suite := suites.MustFind("bn256")
	p2pDkg, err := dkg.CreateP2PDkg(p, suite)
	if err != nil {
		log.Fatal(err)
	}

	dosclient := dosnode.NewDosNode(suite, p, chainConn, p2pDkg)
	if err = dosclient.Start(); err != nil {
		log.Fatal(err)
	}

	done := make(chan interface{})
	<-done
}
