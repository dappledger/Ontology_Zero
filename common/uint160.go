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

package common

import (
	"github.com/Ontology/common/log"
	. "github.com/Ontology/errors"
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"io"
	"math/big"

	"github.com/itchyny/base58-go"
)

const UINT160SIZE int = 20

type Uint160 [UINT160SIZE]uint8

func (u *Uint160) CompareTo(o Uint160) int {
	x := u.ToArray()
	y := o.ToArray()

	for i := len(x) - 1; i >= 0; i-- {
		if x[i] > y[i] {
			return 1
		}
		if x[i] < y[i] {
			return -1
		}
	}

	return 0
}

func (u *Uint160) ToArray() []byte {
	var x []byte = make([]byte, UINT160SIZE)
	for i := 0; i < 20; i++ {
		x[i] = byte(u[i])
	}

	return x
}
func (u *Uint160) ToArrayReverse() []byte {
	var x []byte = make([]byte, UINT160SIZE)
	for i, j := 0, UINT160SIZE - 1; i < j; i, j = i + 1, j - 1 {
		x[i], x[j] = byte(u[j]), byte(u[i])
	}
	return x
}
func (u *Uint160) Serialize(w io.Writer) (int, error) {
	b_buf := bytes.NewBuffer([]byte{})
	binary.Write(b_buf, binary.LittleEndian, u)

	len, err := w.Write(b_buf.Bytes())

	if err != nil {
		return 0, err
	}

	return len, nil
}

func (f *Uint160) Deserialize(r io.Reader) error {
	p := make([]byte, UINT160SIZE)
	n, err := r.Read(p)

	if n <= 0 || err != nil {
		return err
	}

	b_buf := bytes.NewBuffer(p)
	binary.Read(b_buf, binary.LittleEndian, f)

	return nil
}

func (f *Uint160) ToAddress() (string, error) {
	data := append([]byte{23}, f.ToArray()...)
	temp := sha256.Sum256(data)
	temps := sha256.Sum256(temp[:])
	data = append(data, temps[0:4]...)

	bi := new(big.Int).SetBytes(data).String()
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode([]byte(bi))
	if err != nil {
		return "", err
	}
	return string(encoded), nil
}

func Uint160ParseFromBytes(f []byte) (Uint160, error) {
	if len(f) != UINT160SIZE {
		return Uint160{}, NewDetailErr(errors.New("[Common]: Uint160ParseFromBytes err, len != 20"), ErrNoCode, "")
	}

	var hash [20]uint8
	for i := 0; i < 20; i++ {
		hash[i] = f[i]
	}
	return Uint160(hash), nil
}
func ToScriptHash(address string) (Uint160, error) {
	encoding := base58.BitcoinEncoding

	decoded, err := encoding.Decode([]byte(address))
	if err != nil {
		return Uint160{}, err
	}

	x, _ := new(big.Int).SetString(string(decoded), 10)
	log.Tracef("[ToAddress] x: ", x.Bytes())

	ph, err := Uint160ParseFromBytes(x.Bytes()[1:21])
	if err != nil {
		return Uint160{}, err
	}

	log.Tracef("[AddressToProgramHash] programhash: %x", ph.ToArray())

	addr, err := ph.ToAddress()
	if err != nil {
		return Uint160{}, err
	}

	log.Tracef("[AddressToProgramHash] address: %s", addr)

	if addr != address {
		return Uint160{}, errors.New("[AddressToProgramHash]: decode address verify failed.")
	}

	return ph, nil
}
