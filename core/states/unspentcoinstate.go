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

package states

import (
	"io"
	"github.com/Ontology/common/serialization"
)

type UnspentCoinState struct {
	StateBase
	Item []CoinState
}

func (this *UnspentCoinState) Serialize(w io.Writer) error {
	this.StateBase.Serialize(w)
	serialization.WriteUint32(w, uint32(len(this.Item)))
	for _, v := range this.Item {
		serialization.WriteByte(w, byte(v))
	}
	return nil
}

func (this *UnspentCoinState) Deserialize(r io.Reader) error {
	if this == nil {
		this = new(UnspentCoinState)
	}
	err := this.StateBase.Deserialize(r)
	if err != nil {
		return err
	}
	n, err := serialization.ReadUint32(r)
	if err != nil {
		return err
	}
	for i := uint32(0); i < n; i++ {
		state, err := serialization.ReadByte(r)
		if err != nil {
			return err
		}
		this.Item = append(this.Item, CoinState(state))
	}
	return nil
}
