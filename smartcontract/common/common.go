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
	"github.com/Ontology/vm/neovm/types"
	"github.com/Ontology/common"
	"fmt"
	"reflect"
)

type States struct {
	Type string
	Value interface{}
}

func ConvertTypes(item types.StackItemInterface) (results []States) {
	if item == nil {
		return
	}
	switch v := item.(type) {
	case *types.ByteArray:
		results = append(results, States{"ByteArray", common.ToHexString(v.GetByteArray())})
	case *types.Integer:
		if v.GetBigInteger().Sign() == 0 {
			results = append(results, States{"Integer", common.ToHexString([]byte{0})})
		} else {
			results = append(results, States{"Integer", common.ToHexString(types.ConvertBigIntegerToBytes(v.GetBigInteger()))})
		}
	case *types.Boolean:
		if v.GetBoolean() {
			results = append(results, States{"Boolean", common.ToHexString([]byte{1})})
		} else {
			results = append(results, States{"Boolean", common.ToHexString([]byte{0})})
		}
	case *types.Array:
		var arr []States
		for _, val := range v.GetArray() {
			arr = append(arr, ConvertTypes(val)...)
		}
		results = append(results, States{"Array", arr})
	case *types.InteropInterface:
		results = append(results, States{"InteropInterface", common.ToHexString(v.GetInterface().ToArray())})
	case types.StackItemInterface:
		ConvertTypes(v)
	default:
		panic(fmt.Sprintf("[ConvertTypes] Invalid Types: %v", reflect.TypeOf(v)))
	}
	return
}

func ConvertReturnTypes(item types.StackItemInterface) (results []interface{}) {
	if item == nil {
		return
	}
	switch v := item.(type) {
	case *types.ByteArray:
		results = append(results, common.ToHexString(v.GetByteArray()))
	case *types.Integer:
		results = append(results, v.GetBigInteger())
	case *types.Boolean:
		results = append(results, v.GetBoolean())
	case *types.Array:
		var arr []interface{}
		for _, val := range v.GetArray() {
			arr = append(arr, ConvertReturnTypes(val)...)
		}
		results = append(results, arr)
	case *types.InteropInterface:
		results = append(results, common.ToHexString(v.GetInterface().ToArray()))
	case types.StackItemInterface:
		ConvertTypes(v)
	default:
		panic("[ConvertTypes] Invalid Types!")
	}
	return
}