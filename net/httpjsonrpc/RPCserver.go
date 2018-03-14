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

import (
	. "github.com/Ontology/common/config"
	"github.com/Ontology/common/log"
	"net/http"
	"strconv"
)

func StartRPCServer() {
	log.Debug()
	http.HandleFunc("/", Handle)

	HandleFunc("getbestblockhash", getBestBlockHash)
	HandleFunc("getblock", getBlock)
	HandleFunc("getblockcount", getBlockCount)
	HandleFunc("getblockhash", getBlockHash)
	HandleFunc("getunspendoutput", getUnspendOutput)
	HandleFunc("getconnectioncount", getConnectionCount)
	HandleFunc("getrawmempool", getRawMemPool)
	HandleFunc("getrawtransaction", getRawTransaction)
	HandleFunc("sendrawtransaction", sendRawTransaction)
	HandleFunc("getstorage", getStorage)
	HandleFunc("getbalance", getBalance)
	HandleFunc("submitblock", submitBlock)
	HandleFunc("getversion", getVersion)
	HandleFunc("getdataile", getDataFile)
	HandleFunc("catdatarecord", catDataRecord)
	HandleFunc("regdatafile", regDataFile)
	HandleFunc("uploadDataFile", uploadDataFile)
	HandleFunc("getsmartcodeevent", getSmartCodeEvent)

	err := http.ListenAndServe(":" + strconv.Itoa(Parameters.HttpJsonPort), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
