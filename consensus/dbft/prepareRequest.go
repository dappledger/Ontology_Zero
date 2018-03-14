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

package dbft

import (
	. "github.com/Ontology/common"
	"github.com/Ontology/common/log"
	ser "github.com/Ontology/common/serialization"
	tx "github.com/Ontology/core/transaction"
	. "github.com/Ontology/errors"
	"io"
)

type PrepareRequest struct {
	msgData        ConsensusMessageData
	Nonce          uint64
	NextBookKeeper Uint160
	Transactions   []*tx.Transaction
	Signature      []byte
}

func (pr *PrepareRequest) Serialize(w io.Writer) error {
	log.Debug()

	pr.msgData.Serialize(w)
	if err := ser.WriteVarUint(w, pr.Nonce); err != nil {
		return NewDetailErr(err, ErrNoCode, "[PrepareRequest] nonce serialization failed")
	}
	if _, err := pr.NextBookKeeper.Serialize(w); err != nil {
		return NewDetailErr(err, ErrNoCode, "[PrepareRequest] nextbookKeeper serialization failed")
	}
	if err := ser.WriteVarUint(w, uint64(len(pr.Transactions))); err != nil {
		return NewDetailErr(err, ErrNoCode, "[PrepareRequest] length serialization failed")
	}
	for _, t := range pr.Transactions {
		if err := t.Serialize(w); err != nil {
			return NewDetailErr(err, ErrNoCode, "[PrepareRequest] transactions serialization failed")
		}
	}
	if err := ser.WriteVarBytes(w, pr.Signature); err != nil {
		return NewDetailErr(err, ErrNoCode, "[PrepareRequest] signature serialization failed")
	}
	return nil
}

func (pr *PrepareRequest) Deserialize(r io.Reader) error {
	log.Debug()
	pr.msgData = ConsensusMessageData{}
	pr.msgData.Deserialize(r)
	pr.Nonce, _ = ser.ReadVarUint(r, 0)

	pr.NextBookKeeper = Uint160{}
	if err := pr.NextBookKeeper.Deserialize(r); err != nil {
		return NewDetailErr(err, ErrNoCode, "[PrepareRequest] nextbookKeeper deserialization failed")
	}

	length, err := ser.ReadVarUint(r, 0)
	if err != nil {
		return NewDetailErr(err, ErrNoCode, "[PrepareRequest] length deserialization failed")
	}

	pr.Transactions = make([]*tx.Transaction, length)
	for i := 0; i < len(pr.Transactions); i++ {
		var t tx.Transaction
		if err := t.Deserialize(r); err != nil {
			return NewDetailErr(err, ErrNoCode, "[PrepareRequest] transactions deserialization failed")
		}
		pr.Transactions[i] = &t
	}

	pr.Signature, err = ser.ReadVarBytes(r)
	if err != nil {
		return NewDetailErr(err, ErrNoCode, "[PrepareRequest] signature deserialization failed")
	}

	return nil
}

func (pr *PrepareRequest) Type() ConsensusMessageType {
	log.Debug()
	return pr.ConsensusMessageData().Type
}

func (pr *PrepareRequest) ViewNumber() byte {
	log.Debug()
	return pr.msgData.ViewNumber
}

func (pr *PrepareRequest) ConsensusMessageData() *ConsensusMessageData {
	log.Debug()
	return &(pr.msgData)
}
