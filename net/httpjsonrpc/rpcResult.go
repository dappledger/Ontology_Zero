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

package httpjsonrpc

var (
	DnaRpcInvalidHash = responsePacking("invalid hash")
	DnaRpcInvalidBlock = responsePacking("invalid block")
	DnaRpcInvalidTransaction = responsePacking("invalid transaction")
	DnaRpcInvalidParameter = responsePacking("invalid parameter")

	DnaRpcUnknownBlock = responsePacking("unknown block")
	DnaRpcUnknownTransaction = responsePacking("unknown transaction")

	DnaRpcNil = responsePacking(nil)
	DnaRpcUnsupported = responsePacking("Unsupported")
	DnaRpcInternalError = responsePacking("internal error")
	DnaRpcIOError = responsePacking("internal IO error")
	DnaRpcAPIError = responsePacking("internal API error")
	DnaRpcSuccess = responsePacking(true)
	DnaRpcFailed = responsePacking(false)
	DnaRpcAccountNotFound = responsePacking(("Account not found"))

	DnaRpc = responsePacking
)
