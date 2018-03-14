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

package transaction

import (
	"github.com/Ontology/common"
	"io"
)

type BalanceTxInput struct {
	AssetID     common.Uint256
	Value       common.Fixed64
	ProgramHash common.Uint160
}

func (bi *BalanceTxInput) Serialize(w io.Writer) {
	bi.AssetID.Serialize(w)
	bi.Value.Serialize(w)
	bi.ProgramHash.Serialize(w)
}

func (bi *BalanceTxInput) Deserialize(r io.Reader) error {
	err := bi.AssetID.Deserialize(r)
	if err != nil {
		return err
	}

	err = bi.Value.Deserialize(r)
	if err != nil {
		return err
	}

	err = bi.ProgramHash.Deserialize(r)
	if err != nil {
		return err
	}

	return nil
}
