/*
 * Copyright (C) 2018 Onchain <onchain@onchain.com>
 *
 * This file is part of The ontology_Zero.
 *
 * The ontology_Zero is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The ontology_Zero is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The ontology_Zero.  If not, see <http://www.gnu.org/licenses/>.
 */

package consensus

import (
	"fmt"
	cl "github.com/Ontology/account"
	"github.com/Ontology/common/config"
	"github.com/Ontology/common/log"
	"github.com/Ontology/consensus/dbft"
	"github.com/Ontology/consensus/solo"
	"github.com/Ontology/net"
	"strings"
	"time"
)

type ConsensusService interface {
	Start() error
	Halt() error
}

func Log(message string) {
	logMsg := fmt.Sprintf("[%s] %s", time.Now().Format("02/01/2006 15:04:05"), message)
	fmt.Println(logMsg)
	log.Info(logMsg)
}

const (
	CONSENSUS_DBFT = "dbft"
	CONSENSUS_SOLO  = "solo"
)

var ConsensusMgr = NewConsensuManager()

type ConsensusManager struct {}

func NewConsensuManager() *ConsensusManager {
	return &ConsensusManager{}
}

func (this *ConsensusManager)NewConsensusService(client cl.Client , localNet net.Neter)ConsensusService{
	consensusType := strings.ToLower(config.Parameters.ConsensusType)
	if consensusType == "" {
		consensusType = CONSENSUS_DBFT
	}
	var consensus ConsensusService
	switch consensusType {
	case CONSENSUS_DBFT:
		consensus = dbft.NewDbftService(client, "dbft", localNet)
	case CONSENSUS_SOLO:
		consensus = solo.NewSoloService(client, localNet)
	}
	log.Infof("ConsensusType:%s", consensusType)
	return consensus
}
