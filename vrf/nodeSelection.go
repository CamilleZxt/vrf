package vrf

import (
	"github.com/Venachain/Venachain/common"
	"github.com/Venachain/Venachain/common/syscontracts"
	"sort"
)

const ValidatorCount = 1

type NodeForElection struct {
	*syscontracts.NodeInfo
	rank common.Hash
}

type NodesForElection []NodeForElection

func vrfElection(nonce []byte) (int32, error) {
	var nodes NodesForElection

	h1 := common.RlpHash(nonce)

	consensusNodes := NodesForElection{}
	//	// todo 获取所有节点
	for _, node := range nodes {
		if node.Status == 1 && node.Typ == 1 {
			h2 := common.RlpHash(node.PublicKey)
			h := common.Hash{}
			for i, _ := range h {
				h[i] = h1[i] ^ h2[i]
			}
			consensusNodes = append(consensusNodes, NodeForElection{node.NodeInfo, h})
		}
	}
	sort.Sort(consensusNodes)

	if len(consensusNodes) > int(ValidatorCount) {
		consensusNodes = consensusNodes[:ValidatorCount]
	}

	names := make([]string, 0)
	for _, v := range consensusNodes {
		names = append(names, v.Name)
	}
	return 0, nil
}
