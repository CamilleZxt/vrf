package node

import (
	"encoding/json"
	"fmt"
	"github.com/Venachain/Venachain/common"
	"github.com/Venachain/Venachain/common/syscontracts"
	"math/big"
)

type SCNode struct {
	stateDB      StateDB
	contractAddr common.Address
	caller       common.Address
	blockNumber  *big.Int
}

func NewSCNode(db StateDB) *SCNode {
	return &SCNode{stateDB: db, contractAddr: syscontracts.NodeManagementAddress, blockNumber: big.NewInt(0)}
}

func (n *SCNode) importOldNodesData(data string) error {
	str := []byte(data)
	nodes := make([]syscontracts.NodeInfo, 0)
	err := json.Unmarshal(str, &nodes)
	if err != nil {
		fmt.Sprintf("old nodes data unmarshal fail")
		return err
	}
	for index, _ := range nodes {
		//names, err := n.getNames()
		//if err != nil {
		//	if errNodeNotFound != err {
		//		return err
		//	}
		//
		//	names = []string{}
		//}
		//if n.isNameExist(names, nodes[index].Name) {
		//	n.update(nodes[index].Name, &nodes[index])
		//}
		err = n.add(&nodes[index])
		if err != nil {
			return err
		}
	}
	if err != nil {
		return err
	}
	fmt.Sprintf("import old nodes data success")
	return nil
}
