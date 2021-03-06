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

package errors

import (
	"fmt"
)

type ErrCoder interface {
	GetErrCode() ErrCode
}

type ErrCode int32

const (
	ErrNoCode ErrCode = -2
	ErrNoError ErrCode = 0
	ErrUnknown ErrCode = -1
	ErrDuplicatedTx ErrCode = 1
	ErrDuplicateInput ErrCode = 45003
	ErrAssetPrecision ErrCode = 45004
	ErrTransactionBalance ErrCode = 45005
	ErrAttributeProgram ErrCode = 45006
	ErrTransactionContracts ErrCode = 45007
	ErrTransactionPayload ErrCode = 45008
	ErrDoubleSpend ErrCode = 45009
	ErrTxHashDuplicate ErrCode = 45010
	ErrStateUpdaterVaild ErrCode = 45011
	ErrSummaryAsset ErrCode = 45012
	ErrXmitFail ErrCode = 45013
)

func (err ErrCode) Error() string {
	switch err {
	case ErrNoCode:
		return "no error code"
	case ErrNoError:
		return "not an error"
	case ErrUnknown:
		return "unknown error"
	case ErrDuplicatedTx:
		return "duplicated transaction detected"
	case ErrDuplicateInput:
		return "duplicated transaction input detected"
	case ErrAssetPrecision:
		return "invalid asset precision"
	case ErrTransactionBalance:
		return "transaction balance unmatched"
	case ErrAttributeProgram:
		return "attribute program error"
	case ErrTransactionContracts:
		return "invalid transaction contract"
	case ErrTransactionPayload:
		return "invalid transaction payload"
	case ErrDoubleSpend:
		return "double spent transaction detected"
	case ErrTxHashDuplicate:
		return "duplicated transaction hash detected"
	case ErrStateUpdaterVaild:
		return "invalid state updater"
	case ErrSummaryAsset:
		return "invalid summary asset"
	case ErrXmitFail:
		return "transmit error"
	}

	return fmt.Sprintf("Unknown error? Error code = %d", err)
}

func ErrerCode(err error) ErrCode {
	if err, ok := err.(ErrCoder); ok {
		return err.GetErrCode()
	}
	return ErrUnknown
}
