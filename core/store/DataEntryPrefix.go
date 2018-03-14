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

package store

// DataEntryPrefix
type DataEntryPrefix byte

const (
	// DATA
	DATA_Block DataEntryPrefix = iota
	DATA_Header
	DATA_Transaction

	// ASSET
	ST_Account
	ST_Coin
	ST_SpentCoin
	ST_BookKeeper
	ST_Asset
	ST_Contract
	ST_Storage
	ST_Identity
	ST_Program_Coin
	ST_Validator
	ST_Vote

	IX_HeaderHashList

	//SYSTEM
	SYS_CurrentBlock
	SYS_Version
	Sys_CurrentStateRoot
	SYS_BlockMerkleTree
)
