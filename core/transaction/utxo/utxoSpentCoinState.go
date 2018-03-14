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
	. "github.com/Ontology/errors"
	"errors"
	"io"
)

//define the gas stucture in onchain DNA
type SpentCoinState struct {
	TransactionHash   common.Uint256
	TransactionHeight uint32
	Items             []*Item
}

// Serialize is the implement of SignableData interface.
func (this *SpentCoinState) Serialize(w io.Writer) error {
	_, err := this.TransactionHash.Serialize(w)
	if err != nil {
		return NewDetailErr(err, ErrNoCode, "[SpentCoinState], TransactionHash serialize failed.")
	}
	err = serialization.WriteUint32(w, this.TransactionHeight)
	if err != nil {
		return NewDetailErr(err, ErrNoCode, "[SpentCoinState], StartHeight serialize failed.")
	}
	err = serialization.WriteUint32(w, uint32(len(this.Items)))
	if err != nil {
		return NewDetailErr(err, ErrNoCode, "[SpentCoinState], count serialize failed.")
	}
	for _, v := range this.Items {
		err = v.Serialize(w)
		if err != nil {
			return NewDetailErr(err, ErrNoCode, "[SpentCoinState], Item serialize failed.")
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
	this.TransactionHash.Deserialize(r)
	if err != nil {
		return NewDetailErr(errors.New("[SpentCoinState], TransactionHash deserialize failed."), ErrNoCode, "")
	}
	this.TransactionHeight, err = serialization.ReadUint32(r)
	if err != nil {
		return NewDetailErr(errors.New("[SpentCoinState], TransactionHeight deserialize failed."), ErrNoCode, "")
	}
	count, err := serialization.ReadUint32(r)
	if err != nil {
		return NewDetailErr(errors.New("[SpentCoinState], count deserialize failed."), ErrNoCode, "")
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

type Item struct {
	PrevIndex uint32
	EndHeight uint32
}

// Serialize is the implement of SignableData interface.
func (this *Item) Serialize(w io.Writer) error {
	err := serialization.WriteUint32(w, this.PrevIndex)
	if err != nil {
		return NewDetailErr(err, ErrNoCode, "[Items], PrevIndex serialize failed.")
	}
	err = serialization.WriteUint32(w, this.EndHeight)
	if err != nil {
		return NewDetailErr(err, ErrNoCode, "[Items], EndHeight serialize failed.")
	}
	return nil
}

// Deserialize is the implement of SignableData interface.
func (this *Item) Deserialize(r io.Reader) error {
	if this == nil {
		this = new(Item)
	}
	var err error
	this.PrevIndex, err = serialization.ReadUint32(r)
	if err != nil {
		return NewDetailErr(errors.New("[Items], PrevIndex deserialize failed."), ErrNoCode, "")
	}
	this.EndHeight, err = serialization.ReadUint32(r)
	if err != nil {
		return NewDetailErr(errors.New("[Items], EndHeight deserialize failed."), ErrNoCode, "")
	}
	return nil
}
