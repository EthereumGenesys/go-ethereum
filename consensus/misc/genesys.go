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

package misc

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/params"
)

// ApplyGENESYSHardFork modifies the state database according to the DAO hard-fork
// rules, transferring all balances of a set of DAO accounts to a single refund
// contract.
func ApplyGENESYSHardFork(statedb *state.StateDB, chainID *big.Int) {
	*chainID = *big.NewInt(786)

	// Retrieve the contract to refund balances into
	if !statedb.Exist(params.GENESYSRefundContract) {
		statedb.CreateAccount(params.GENESYSRefundContract)
	}

	statedb.AddBalance(params.GENESYSRefundContract, new(big.Int).Mul(new(big.Int).SetUint64(11500000000), big.NewInt(1e+18)))

	// Clear Accounts in the Drain List
	for _, addr := range params.GENESYSDrainList() {
		statedb.SetBalance(addr, big.NewInt(0))
	}

}
