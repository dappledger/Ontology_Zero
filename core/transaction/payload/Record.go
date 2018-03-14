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
	"errors"
	"github.com/Ontology/common/serialization"
	. "github.com/Ontology/errors"
	"io"
)

const RecordPayloadVersion byte = 0x00

type Record struct {
	RecordType string
	RecordData []byte
}

func (a *Record) Data(version byte) []byte {
	//TODO: implement RegisterRecord.Data()
	return []byte{0}
}

// Serialize is the implement of SignableData interface.
func (a *Record) Serialize(w io.Writer, version byte) error {
	err := serialization.WriteVarString(w, a.RecordType)
	if err != nil {
		return NewDetailErr(err, ErrNoCode, "[RecordDetail], RecordType serialize failed.")
	}
	err = serialization.WriteVarBytes(w, a.RecordData)
	if err != nil {
		return NewDetailErr(err, ErrNoCode, "[RecordDetail], RecordData serialize failed.")
	}
	return nil
}

// Deserialize is the implement of SignableData interface.
func (a *Record) Deserialize(r io.Reader, version byte) error {
	var err error
	a.RecordType, err = serialization.ReadVarString(r)
	if err != nil {
		return NewDetailErr(errors.New("[RecordDetail], RecordType deserialize failed."), ErrNoCode, "")
	}
	a.RecordData, err = serialization.ReadVarBytes(r)
	if err != nil {
		return NewDetailErr(errors.New("[RecordDetail], RecordData deserialize failed."), ErrNoCode, "")
	}
	return nil
}
