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
	"github.com/Ontology/common"
	"github.com/Ontology/core/asset"
	"github.com/Ontology/crypto"
	. "github.com/Ontology/errors"
	"io"
)

const RegisterPayloadVersion byte = 0x00

type RegisterAsset struct {
	Asset      *asset.Asset
	Amount     common.Fixed64
	//Precision  byte
	Issuer     *crypto.PubKey
	Controller common.Uint160
}

func (a *RegisterAsset) Data(version byte) []byte {
	//TODO: implement RegisterAsset.Data()
	return []byte{0}

}

func (a *RegisterAsset) Serialize(w io.Writer, version byte) error {
	a.Asset.Serialize(w)
	a.Amount.Serialize(w)
	//w.Write([]byte{a.Precision})
	a.Issuer.Serialize(w)
	a.Controller.Serialize(w)
	return nil
}

func (a *RegisterAsset) Deserialize(r io.Reader, version byte) error {

	//asset
	a.Asset = new(asset.Asset)
	err := a.Asset.Deserialize(r)
	if err != nil {
		return NewDetailErr(err, ErrNoCode, "[RegisterAsset], Asset Deserialize failed.")
	}

	//Amount
	a.Amount = *new(common.Fixed64)
	err = a.Amount.Deserialize(r)
	if err != nil {
		return NewDetailErr(err, ErrNoCode, "[RegisterAsset], Ammount Deserialize failed.")
	}

	//Precision  byte 02/10 comment out by wjj
	//p := make([]byte, 1)
	//n, err := r.Read(p)
	//if n > 0 {
	//	a.Precision = p[0]
	//} else {
	//	return NewDetailErr(err, ErrNoCode, "[RegisterAsset], Precision Deserialize failed.")
	//}

	//Issuer     *crypto.PubKey
	a.Issuer = new(crypto.PubKey)
	err = a.Issuer.DeSerialize(r)
	if err != nil {
		return NewDetailErr(err, ErrNoCode, "[RegisterAsset], Ammount Deserialize failed.")
	}

	//Controller *common.Uint160
	a.Controller = *new(common.Uint160)
	err = a.Controller.Deserialize(r)
	if err != nil {
		return NewDetailErr(err, ErrNoCode, "[RegisterAsset], Ammount Deserialize failed.")
	}
	return nil
}
