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
	"io"
	"github.com/Ontology/common"
	"github.com/Ontology/common/serialization"
)

type InvokeCode struct {
	CodeHash common.Uint160
	Code     []byte
}

func (ic *InvokeCode) Data(version byte) []byte {
	return []byte{0}
}

func (ic *InvokeCode) Serialize(w io.Writer, version byte) error {
	ic.CodeHash.Serialize(w)
	err := serialization.WriteVarBytes(w, ic.Code)
	if err != nil {
		return err
	}
	return nil
}

func (ic *InvokeCode) Deserialize(r io.Reader, version byte) error {
	u := new(common.Uint160)
	if err := u.Deserialize(r); err != nil {
		return err
	}
	ic.CodeHash = *u
	code, err := serialization.ReadVarBytes(r)
	if err != nil {
		return err
	}
	ic.Code = code
	return nil
}