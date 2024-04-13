package main

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/koustubh25/go-coinbase/pkg/advancedtradeapi"
	"github.com/koustubh25/go-coinbase/pkg/advancedtradeapi/websocket"
	"github.com/koustubh25/trader/pkg/algorithm/candle"
)

func main() {
	// client, err := advancedtradeapi.NewRESTClient(advancedtradeapi.WithProduction())
	// if err != nil {
	// 	panic(err)
	// }
	// ctx := context.Background()
	// a, err := rest.ListAccounts[rest.AccountsEnvelope](ctx, client)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(a)

	slog.SetLogLoggerLevel(slog.LevelDebug)
	websocketClient, err := advancedtradeapi.NewWebSocketClient(advancedtradeapi.WithProduction())
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	msgs, err := websocket.Subscribe[candle.CandlesMessage](ctx, "candles", []string{"BIGTIME-AUD"}, websocketClient)
	if err != nil {
		panic(err)
	}
	hammerCandle := candle.HammerCandleAgorithm{}
	for msg := range msgs {
		if len(msg.Events) > 0 {
			latestEvent := msg.Events[len(msg.Events)-1]
			if len(latestEvent.Candles) > 0 {
				latestCandle := latestEvent.Candles[len(latestEvent.Candles)-1]
				hammerCandle.Candle = latestCandle
				isBuy, _ := hammerCandle.IsBuy()
				isSell, _ := hammerCandle.IsSell()
				if isBuy {
					slog.Info(fmt.Sprintf("Hammer detected, buying: %+v", msg))
				}
				if isSell {
					slog.Info(fmt.Sprintf("Inverted hammer detected, selling: %+v", msg))
				}
			}
		}
	}
}
