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

package payload

import (
	"github.com/Ontology/common/serialization"
	"github.com/Ontology/crypto"
	. "github.com/Ontology/errors"
	"io"
)

const DataFilePayloadVersion byte = 0x00

type DataFile struct {
	IPFSPath string
	Filename string
	Note     string
	Issuer   *crypto.PubKey
	//TODO: add hash or key to verify data
}

func (a *DataFile) Data(version byte) []byte {
	//TODO: implement RegisterRecord.Data()
	return []byte{0}
}

// Serialize is the implement of SignableData interface.
func (a *DataFile) Serialize(w io.Writer, version byte) error {
	err := serialization.WriteVarString(w, a.IPFSPath)
	if err != nil {
		return NewDetailErr(err, ErrNoCode, "[DataFileDetail], IPFSPath serialize failed.")
	}
	err = serialization.WriteVarString(w, a.Filename)
	if err != nil {
		return NewDetailErr(err, ErrNoCode, "[DataFileDetail], Filename serialize failed.")
	}
	err = serialization.WriteVarString(w, a.Note)
	if err != nil {
		return NewDetailErr(err, ErrNoCode, "[DataFileDetail], Note serialize failed.")
	}
	a.Issuer.Serialize(w)

	return nil
}

// Deserialize is the implement of SignableData interface.
func (a *DataFile) Deserialize(r io.Reader, version byte) error {
	var err error
	a.IPFSPath, err = serialization.ReadVarString(r)
	if err != nil {
		return NewDetailErr(err, ErrNoCode, "[DataFileDetail], IPFSPath deserialize failed.")
	}
	a.Filename, err = serialization.ReadVarString(r)
	if err != nil {
		return NewDetailErr(err, ErrNoCode, "[DataFileDetail], Filename deserialize failed.")
	}
	a.Note, err = serialization.ReadVarString(r)
	if err != nil {
		return NewDetailErr(err, ErrNoCode, "[DataFileDetail], Note deserialize failed.")
	}
	//Issuer     *crypto.PubKey
	a.Issuer = new(crypto.PubKey)
	err = a.Issuer.DeSerialize(r)
	if err != nil {
		return NewDetailErr(err, ErrNoCode, "[DataFileDetail], Issuer deserialize failed.")
	}

	return nil
}
