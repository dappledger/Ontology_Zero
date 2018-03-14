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

package code

import (
	"fmt"
	. "github.com/Ontology/common"
	"github.com/Ontology/common/log"
	"github.com/Ontology/common/serialization"
	. "github.com/Ontology/core/contract"
	. "github.com/Ontology/errors"
	"io"
)

type FunctionCode struct {
	// Contract Code
	Code           []byte

	// Contract parameter type list
	ParameterTypes []ContractParameterType

	// Contract return type
	ReturnType     ContractParameterType

	codeHash       Uint160
}

// method of SerializableData
func (fc *FunctionCode) Serialize(w io.Writer) error {
	var err error
	err = serialization.WriteVarBytes(w, fc.Code)
	if err != nil {
		return err
	}

	err = serialization.WriteVarBytes(w, ContractParameterTypeToByte(fc.ParameterTypes))
	if err != nil {
		return err
	}

	err = serialization.WriteByte(w, byte(fc.ReturnType))
	if err != nil {
		return err
	}

	return nil
}

// method of SerializableData
func (fc *FunctionCode) Deserialize(r io.Reader) error {
	var err error

	fc.Code, err = serialization.ReadVarBytes(r)
	if err != nil {
		return NewDetailErr(err, ErrNoCode, "Transaction FunctionCode Code Deserialize failed.")
	}

	p, err := serialization.ReadVarBytes(r)
	if err != nil {
		return NewDetailErr(err, ErrNoCode, "Transaction FunctionCode ParameterTypes Deserialize failed.")
	}
	fc.ParameterTypes = ByteToContractParameterType(p)

	returnType, err := serialization.ReadByte(r)
	if err != nil {
		return NewDetailErr(err, ErrNoCode, "Transaction FunctionCode returnType Deserialize failed.")
	}
	fc.ReturnType = ContractParameterType(returnType)
	return nil
}

// method of ICode
// Get the hash of the smart contract
func (fc *FunctionCode) CodeHash() Uint160 {
	u160 := Uint160{}
	if fc.codeHash == u160 {
		u160, err := ToCodeHash(fc.Code)
		if err != nil {
			log.Debug(fmt.Sprintf("[FunctionCode] ToCodeHash err=%s", err))
			return u160
		}
		fc.codeHash = u160
	}
	return fc.codeHash
}
