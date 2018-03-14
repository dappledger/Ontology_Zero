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
	"github.com/Ontology/common"
	"io"
	. "github.com/Ontology/common/serialization"
)

type SpentCoinState struct {
	StateBase
	TransactionHash   common.Uint256
	TransactionHeight uint32
	Items             []*Item
}

type Item struct {
	StateBase
	PrevIndex uint16
	EndHeight uint32
}

func (this *Item) Serialize(w io.Writer) error {
	this.StateBase.Serialize(w)
	err := WriteUint16(w, this.PrevIndex)
	if err != nil {
		return err
	}
	err = WriteUint32(w, this.EndHeight)
	if err != nil {
		return err
	}
	return nil
}

func (this *Item) Deserialize(r io.Reader) error {
	var err error
	err = this.StateBase.Deserialize(r)
	if err != nil {
		return err
	}
	this.PrevIndex, err = ReadUint16(r)
	if err != nil {
		return err
	}
	this.EndHeight, err = ReadUint32(r)
	if err != nil {
		return err
	}
	return nil
}

func (this *SpentCoinState) Serialize(w io.Writer) error {
	this.StateBase.Serialize(w)
	_, err := this.TransactionHash.Serialize(w)
	if err != nil {
		return err
	}
	err = WriteUint32(w, this.TransactionHeight)
	if err != nil {
		return err
	}
	err = WriteUint32(w, uint32(len(this.Items)))
	if err != nil {
		return err
	}
	for _, v := range this.Items {
		err = v.Serialize(w)
		if err != nil {
			return err
		}
	}

	return nil
}

// Deserialize is the implement of SignableData interface.
func (this *SpentCoinState) Deserialize(r io.Reader) error {
	if this == nil {
		this = new(SpentCoinState)
	}
	var err error
	err = this.StateBase.Deserialize(r)
	if err != nil {
		return err
	}
	this.TransactionHash.Deserialize(r)
	if err != nil {
		return err
	}
	this.TransactionHeight, err = ReadUint32(r)
	if err != nil {
		return err
	}
	count, err := ReadUint32(r)
	if err != nil {
		return err
	}
	for i := 0; i < int(count); i++ {
		item_ := new(Item)
		err := item_.Deserialize(r)
		if err != nil {
			return err
		}
		this.Items = append(this.Items, item_)
	}
	return nil
}

func (this *SpentCoinState) RemoveItem(i int) {
	this.Items[i] = this.Items[len(this.Items) - 1]
	this.Items = this.Items[:len(this.Items) - 1]
}

