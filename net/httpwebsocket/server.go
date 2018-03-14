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

package httpwebsocket

import (
	"bytes"
	. "github.com/Ontology/common"
	. "github.com/Ontology/common/config"
	"github.com/Ontology/core/ledger"
	"github.com/Ontology/events"
	"github.com/Ontology/net/httprestful/common"
	Err "github.com/Ontology/net/httprestful/error"
	"github.com/Ontology/net/httpwebsocket/websocket"
	. "github.com/Ontology/net/protocol"
	"github.com/Ontology/smartcontract/event"
	sc "github.com/Ontology/smartcontract/common"
)

var ws *websocket.WsServer
var (
	pushBlockFlag bool = false
	pushRawBlockFlag bool = false
	pushBlockTxsFlag bool = false
)

func StartServer(n Noder) {
	common.SetNode(n)
	ledger.DefaultLedger.Blockchain.BCEvents.Subscribe(events.EventBlockPersistCompleted, SendBlock2WSclient)
	ledger.DefaultLedger.Blockchain.BCEvents.Subscribe(events.EventSmartCode, PushSmartCodeEvent)
	go func() {
		ws = websocket.InitWsServer(common.CheckAccessToken)
		ws.Start()
	}()
}
func SendBlock2WSclient(v interface{}) {
	if Parameters.HttpWsPort != 0 && pushBlockFlag {
		go func() {
			PushBlock(v)
		}()
	}
	if Parameters.HttpWsPort != 0 && pushBlockTxsFlag {
		go func() {
			PushBlockTransactions(v)
		}()
	}
}
func Stop() {
	if ws == nil {
		return
	}
	ws.Stop()
}
func ReStartServer() {
	if ws == nil {
		ws = websocket.InitWsServer(common.CheckAccessToken)
		ws.Start()
		return
	}
	ws.Restart()
}
func GetWsPushBlockFlag() bool {
	return pushBlockFlag
}
func SetWsPushBlockFlag(b bool) {
	pushBlockFlag = b
}
func GetPushRawBlockFlag() bool {
	return pushRawBlockFlag
}
func SetPushRawBlockFlag(b bool) {
	pushRawBlockFlag = b
}
func GetPushBlockTxsFlag() bool {
	return pushBlockTxsFlag
}
func SetPushBlockTxsFlag(b bool) {
	pushBlockTxsFlag = b
}
func SetTxHashMap(txhash string, sessionid string) {
	if ws == nil {
		return
	}
	ws.SetTxHashMap(txhash, sessionid)
}

func PushSmartCodeEvent(v interface{}) {
	if ws != nil {
		rs,ok := v.(map[string]interface{})
		if !ok {
			return
		}
		go func() {
			switch object := rs["Result"].(type) {
			case event.LogEventArgs:
				type LogEventArgsInfo struct {
					Container string
					CodeHash  string
					Message   string
					BlockHeight uint32
				}
				msg :=LogEventArgsInfo{
					Container: ToHexString(object.Container.ToArray()),
					CodeHash:  ToHexString(object.CodeHash.ToArray()),
					Message:   object.Message,
					BlockHeight: ledger.DefaultLedger.Store.GetHeight(),
				}
				PushEvent(rs["TxHash"].(string),rs["Error"].(int64),rs["Action"].(string),msg)
				return
			case event.NotifyEventArgs:
				type NotifyEventArgsInfo struct {
					Container string
					CodeHash  string
					State     []sc.States
					BlockHeight uint32
				}
				msg := NotifyEventArgsInfo{
					Container: ToHexString(object.Container.ToArray()),
					CodeHash:  ToHexString(object.CodeHash.ToArray()),
					State:   sc.ConvertTypes(object.State),
					BlockHeight: ledger.DefaultLedger.Store.GetHeight(),
				}
				PushEvent(rs["TxHash"].(string),rs["Error"].(int64),rs["Action"].(string),msg)
				return
			default:
				PushEvent(rs["TxHash"].(string),rs["Error"].(int64),rs["Action"].(string),rs["Result"])
				return
			}
		}()
	}
}

func PushEvent(txHash string, errcode int64, action string, result interface{}) {
	if ws != nil {
		resp := common.ResponsePack(Err.SUCCESS)
		resp["Result"] = result
		resp["Error"] = errcode
		resp["Action"] = action
		resp["Desc"] = Err.ErrMap[resp["Error"].(int64)]
		ws.PushTxResult(txHash, resp)
		//ws.BroadcastResult(resp)
	}
}

func PushBlock(v interface{}) {
	if ws == nil {
		return
	}
	resp := common.ResponsePack(Err.SUCCESS)
	if block, ok := v.(*ledger.Block); ok {
		if pushRawBlockFlag {
			w := bytes.NewBuffer(nil)
			block.Serialize(w)
			resp["Result"] = ToHexString(w.Bytes())
		} else {
			resp["Result"] = common.GetBlockInfo(block)
		}
		resp["Action"] = "sendrawblock"
		ws.BroadcastResult(resp)
	}
}
func PushBlockTransactions(v interface{}) {
	if ws == nil {
		return
	}
	resp := common.ResponsePack(Err.SUCCESS)
	if block, ok := v.(*ledger.Block); ok {
		if pushBlockTxsFlag {
			resp["Result"] = common.GetBlockTransactions(block)
		}
		resp["Action"] = "sendblocktransactions"
		ws.BroadcastResult(resp)
	}
}
