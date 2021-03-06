package main

import (
	. "github.com/coinrust/crex"
	"github.com/coinrust/crex/brokers"
	"log"
)

func main() {
	ws := brokers.NewWS(brokers.HBDM,
		"[accessKey]", "[secretKey]", false, nil)

	// 订单薄事件方法
	ws.On(WSEventL2Snapshot, func(ob *OrderBook) {
		log.Printf("ob: %#v", ob)
	})
	// 成交记录事件方法
	ws.On(WSEventTrade, func(trades []Trade) {
		log.Printf("trades: %#v", trades)
	})

	// 订单事件方法
	ws.On(WSEventOrder, func(orders []Order) {
		log.Printf("orders: %#v", orders)
	})
	// 持仓事件方法
	ws.On(WSEventPosition, func(positions []Position) {
		log.Printf("positions: %#v", positions)
	})

	// 订阅订单薄
	ws.SubscribeLevel2Snapshots(Market{
		ID:     "BTC",
		Params: ContractTypeW1,
	})
	// 订阅成交记录
	ws.SubscribeTrades(Market{
		ID:     "BTC",
		Params: ContractTypeW1,
	})
	// 订阅订单成交信息
	ws.SubscribeOrders(Market{
		ID:     "BTC",
		Params: ContractTypeW1,
	})
	// 订阅持仓信息
	ws.SubscribePositions(Market{
		ID:     "BTC",
		Params: ContractTypeW1,
	})

	select {}
}
