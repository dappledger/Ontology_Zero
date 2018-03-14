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
	. "github.com/Ontology/core/code"
	"io"
	. "github.com/Ontology/errors"
	"github.com/Ontology/smartcontract/types"
)

const DeployCodePayloadVersion byte = 0x00

type DeployCode struct {
	Code        *FunctionCode
	VmType      types.VmType
	NeedStorage bool
	Name        string
	CodeVersion string
	Author      string
	Email       string
	Description string
}

func (dc *DeployCode) Data(version byte) []byte {
	// TODO: Data()

	return []byte{0}
}

func (dc *DeployCode) Serialize(w io.Writer, version byte) error {
	if dc.Code == nil {
		dc.Code = new(FunctionCode)
	}
	err := dc.Code.Serialize(w)
	if err != nil {
		return NewDetailErr(err, ErrNoCode, "Transaction DeployCode Code Serialize failed.")
	}

	err = serialization.WriteByte(w, byte(dc.VmType))
	if err != nil {
		return err
	}

	err = serialization.WriteBool(w, dc.NeedStorage)
	if err != nil {
		return err
	}

	err = serialization.WriteVarString(w, dc.Name)
	if err != nil {
		return err
	}

	err = serialization.WriteVarString(w, dc.CodeVersion)
	if err != nil {
		return err
	}

	err = serialization.WriteVarString(w, dc.Author)
	if err != nil {
		return err
	}

	err = serialization.WriteVarString(w, dc.Email)
	if err != nil {
		return err
	}

	err = serialization.WriteVarString(w, dc.Description)
	if err != nil {
		return err
	}

	return nil
}

func (dc *DeployCode) Deserialize(r io.Reader, version byte) error {
	if dc.Code == nil {
		dc.Code = new(FunctionCode)
	}

	err := dc.Code.Deserialize(r)
	if err != nil {
		return NewDetailErr(err, ErrNoCode, "Transaction DeployCode Code Deserialize failed.")
	}

	vmType, err := serialization.ReadByte(r)
	if err != nil {
		return err
	}
	dc.VmType = types.VmType(vmType)

	dc.NeedStorage, err = serialization.ReadBool(r)
	if err != nil {
		return NewDetailErr(err, ErrNoCode, "Transaction DeployCode NeedStorage Deserialize failed.")
	}

	dc.Name, err = serialization.ReadVarString(r)
	if err != nil {
		return NewDetailErr(err, ErrNoCode, "Transaction DeployCode Name Deserialize failed.")
	}

	dc.CodeVersion, err = serialization.ReadVarString(r)
	if err != nil {
		return NewDetailErr(err, ErrNoCode, "Transaction DeployCode CodeVersion Deserialize failed.")
	}

	dc.Author, err = serialization.ReadVarString(r)
	if err != nil {
		return NewDetailErr(err, ErrNoCode, "Transaction DeployCode Author Deserialize failed.")
	}

	dc.Email, err = serialization.ReadVarString(r)
	if err != nil {
		return NewDetailErr(err, ErrNoCode, "Transaction DeployCode Email Deserialize failed.")
	}

	dc.Description, err = serialization.ReadVarString(r)
	if err != nil {
		return NewDetailErr(err, ErrNoCode, "Transaction DeployCode Description Deserialize failed.")
	}

	return nil
}
