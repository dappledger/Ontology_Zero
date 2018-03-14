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

package utxo

import (
	"github.com/Ontology/common"
	"github.com/Ontology/common/serialization"
	"io"
	"bytes"
)

type UTXOUnspent struct {
	Txid  common.Uint256
	Index uint32
	Value common.Fixed64
}

func (uu *UTXOUnspent) Serialize(w io.Writer) {
	uu.Txid.Serialize(w)
	serialization.WriteUint32(w, uu.Index)
	uu.Value.Serialize(w)
}

func (uu *UTXOUnspent) Deserialize(r io.Reader) error {
	uu.Txid.Deserialize(r)

	index, err := serialization.ReadUint32(r)
	uu.Index = uint32(index)
	if err != nil {
		return err
	}

	uu.Value.Deserialize(r)

	return nil
}

func (uu *UTXOUnspent) ToArray() []byte {
	bf := new(bytes.Buffer)
	uu.Serialize(bf)
	return bf.Bytes()
}