// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import (
	"github.com/ethereum/go-ethereum/common"
)

// GENESYSRefundContract is the address of the refund contract to send genesys balances to.
var GENESYSRefundContract = common.HexToAddress("0x84103fcaeb28be34f8b4da8017c126a260385a1e")

// GENESYSDrainList is the list of accounts whose full balances will be moved into a
// refund contract at the beginning of the genesys-fork block.
func GENESYSDrainList() []common.Address {
	return []common.Address{
		common.HexToAddress("0x00000000219ab540356cBB839Cbe05303d7705Fa"),
	}
}
