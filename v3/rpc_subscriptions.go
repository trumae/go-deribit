// Code generated by make generate-methods DO NOT EDIT.

package deribit

import (
	"encoding/json"
	"fmt"

	"github.com/adampointer/go-deribit/v3/client/operations"
	"github.com/adampointer/go-deribit/v3/models"
)

// SubscribeBookInterval subscribes to the book.{instrument_name}.{interval} channel
func (e *Exchange) SubscribeBookInterval(instrument_name, interval string) (chan *models.BookNotificationRaw, error) {
	chans := []string{fmt.Sprintf("book.%s.%s", instrument_name, interval)}

	c := make(chan *RPCNotification)
	out := make(chan *models.BookNotificationRaw)
	sub := &RPCSubscription{Data: c, Channel: chans[0]}
	e.subscriptions[chans[0]] = sub

	client := e.Client()
	if _, err := client.GetPublicSubscribe(&operations.GetPublicSubscribeParams{Channels: chans}); err != nil {
		delete(e.subscriptions, chans[0])
		return nil, fmt.Errorf("error subscribing to channel: %s", err)
	}

	go func() {
	Loop:
		for {
			select {
			case n := <-c:
				if len(n.Params.Data) < 1 {
					e.errors <- fmt.Errorf("invalid json data: %s", string(n.Params.Data))
					continue
				}
				data := preFilter(n.Params.Data)
				switch byte(data[0]) {
				case '{':
					var ret models.BookNotificationRaw
					if err := json.Unmarshal(data, &ret); err != nil {
						e.errors <- fmt.Errorf("error decoding notification: %s", err)
					}
					out <- &ret
				case '[':
					var rets []models.BookNotificationRaw
					if err := json.Unmarshal(data, &rets); err != nil {
						e.errors <- fmt.Errorf("error decoding notification: %s", err)
					}
					for _, ret := range rets {
						out <- &ret
					}
				default:
					e.errors <- fmt.Errorf("invalid json data: %s", string(data))
				}
			case <-e.stop:
				break Loop
			}
		}
	}()
	return out, nil
}

// SubscribeEstimatedExpirationPrice subscribes to the estimated_expiration_price.{index_name} channel
func (e *Exchange) SubscribeEstimatedExpirationPrice(index_name string) (chan *models.EstimatedExpirationPriceNotification, error) {
	chans := []string{fmt.Sprintf("estimated_expiration_price.%s", index_name)}

	c := make(chan *RPCNotification)
	out := make(chan *models.EstimatedExpirationPriceNotification)
	sub := &RPCSubscription{Data: c, Channel: chans[0]}
	e.subscriptions[chans[0]] = sub

	client := e.Client()
	if _, err := client.GetPublicSubscribe(&operations.GetPublicSubscribeParams{Channels: chans}); err != nil {
		delete(e.subscriptions, chans[0])
		return nil, fmt.Errorf("error subscribing to channel: %s", err)
	}

	go func() {
	Loop:
		for {
			select {
			case n := <-c:
				if len(n.Params.Data) < 1 {
					e.errors <- fmt.Errorf("invalid json data: %s", string(n.Params.Data))
					continue
				}
				data := preFilter(n.Params.Data)
				switch byte(data[0]) {
				case '{':
					var ret models.EstimatedExpirationPriceNotification
					if err := json.Unmarshal(data, &ret); err != nil {
						e.errors <- fmt.Errorf("error decoding notification: %s", err)
					}
					out <- &ret
				case '[':
					var rets []models.EstimatedExpirationPriceNotification
					if err := json.Unmarshal(data, &rets); err != nil {
						e.errors <- fmt.Errorf("error decoding notification: %s", err)
					}
					for _, ret := range rets {
						out <- &ret
					}
				default:
					e.errors <- fmt.Errorf("invalid json data: %s", string(data))
				}
			case <-e.stop:
				break Loop
			}
		}
	}()
	return out, nil
}

// SubscribePerpetual subscribes to the perpetual.{instrument_name}.{interval} channel
func (e *Exchange) SubscribePerpetual(instrument_name, interval string) (chan *models.PerpetualNotification, error) {
	chans := []string{fmt.Sprintf("perpetual.%s.%s", instrument_name, interval)}

	c := make(chan *RPCNotification)
	out := make(chan *models.PerpetualNotification)
	sub := &RPCSubscription{Data: c, Channel: chans[0]}
	e.subscriptions[chans[0]] = sub

	client := e.Client()
	if _, err := client.GetPublicSubscribe(&operations.GetPublicSubscribeParams{Channels: chans}); err != nil {
		delete(e.subscriptions, chans[0])
		return nil, fmt.Errorf("error subscribing to channel: %s", err)
	}

	go func() {
	Loop:
		for {
			select {
			case n := <-c:
				if len(n.Params.Data) < 1 {
					e.errors <- fmt.Errorf("invalid json data: %s", string(n.Params.Data))
					continue
				}
				data := preFilter(n.Params.Data)
				switch byte(data[0]) {
				case '{':
					var ret models.PerpetualNotification
					if err := json.Unmarshal(data, &ret); err != nil {
						e.errors <- fmt.Errorf("error decoding notification: %s", err)
					}
					out <- &ret
				case '[':
					var rets []models.PerpetualNotification
					if err := json.Unmarshal(data, &rets); err != nil {
						e.errors <- fmt.Errorf("error decoding notification: %s", err)
					}
					for _, ret := range rets {
						out <- &ret
					}
				default:
					e.errors <- fmt.Errorf("invalid json data: %s", string(data))
				}
			case <-e.stop:
				break Loop
			}
		}
	}()
	return out, nil
}

// SubscribeUserOrdersInstrumentName subscribes to the user.orders.{instrument_name}.{interval} channel
func (e *Exchange) SubscribeUserOrdersInstrumentName(instrument_name, interval string) (chan *models.Order, error) {
	chans := []string{fmt.Sprintf("user.orders.%s.%s", instrument_name, interval)}

	c := make(chan *RPCNotification)
	out := make(chan *models.Order)
	sub := &RPCSubscription{Data: c, Channel: chans[0]}
	e.subscriptions[chans[0]] = sub

	client := e.Client()
	if _, err := client.GetPrivateSubscribe(&operations.GetPrivateSubscribeParams{Channels: chans}); err != nil {
		delete(e.subscriptions, chans[0])
		return nil, fmt.Errorf("error subscribing to channel: %s", err)
	}

	go func() {
	Loop:
		for {
			select {
			case n := <-c:
				if len(n.Params.Data) < 1 {
					e.errors <- fmt.Errorf("invalid json data: %s", string(n.Params.Data))
					continue
				}
				data := preFilter(n.Params.Data)
				switch byte(data[0]) {
				case '{':
					var ret models.Order
					if err := json.Unmarshal(data, &ret); err != nil {
						e.errors <- fmt.Errorf("error decoding notification: %s", err)
					}
					out <- &ret
				case '[':
					var rets []models.Order
					if err := json.Unmarshal(data, &rets); err != nil {
						e.errors <- fmt.Errorf("error decoding notification: %s", err)
					}
					for _, ret := range rets {
						out <- &ret
					}
				default:
					e.errors <- fmt.Errorf("invalid json data: %s", string(data))
				}
			case <-e.stop:
				break Loop
			}
		}
	}()
	return out, nil
}

// SubscribeUserTradesInstrument subscribes to the user.trades.{instrument_name}.{interval} channel
func (e *Exchange) SubscribeUserTradesInstrument(instrument_name, interval string) (chan *models.UserTrade, error) {
	chans := []string{fmt.Sprintf("user.trades.%s.%s", instrument_name, interval)}

	c := make(chan *RPCNotification)
	out := make(chan *models.UserTrade)
	sub := &RPCSubscription{Data: c, Channel: chans[0]}
	e.subscriptions[chans[0]] = sub

	client := e.Client()
	if _, err := client.GetPrivateSubscribe(&operations.GetPrivateSubscribeParams{Channels: chans}); err != nil {
		delete(e.subscriptions, chans[0])
		return nil, fmt.Errorf("error subscribing to channel: %s", err)
	}

	go func() {
	Loop:
		for {
			select {
			case n := <-c:
				if len(n.Params.Data) < 1 {
					e.errors <- fmt.Errorf("invalid json data: %s", string(n.Params.Data))
					continue
				}
				data := preFilter(n.Params.Data)
				switch byte(data[0]) {
				case '{':
					var ret models.UserTrade
					if err := json.Unmarshal(data, &ret); err != nil {
						e.errors <- fmt.Errorf("error decoding notification: %s", err)
					}
					out <- &ret
				case '[':
					var rets []models.UserTrade
					if err := json.Unmarshal(data, &rets); err != nil {
						e.errors <- fmt.Errorf("error decoding notification: %s", err)
					}
					for _, ret := range rets {
						out <- &ret
					}
				default:
					e.errors <- fmt.Errorf("invalid json data: %s", string(data))
				}
			case <-e.stop:
				break Loop
			}
		}
	}()
	return out, nil
}

// SubscribeDeribitPriceIndex subscribes to the deribit_price_index.{index_name} channel
func (e *Exchange) SubscribeDeribitPriceIndex(index_name string) (chan *models.DeribitPriceIndexNotification, error) {
	chans := []string{fmt.Sprintf("deribit_price_index.%s", index_name)}

	c := make(chan *RPCNotification)
	out := make(chan *models.DeribitPriceIndexNotification)
	sub := &RPCSubscription{Data: c, Channel: chans[0]}
	e.subscriptions[chans[0]] = sub

	client := e.Client()
	if _, err := client.GetPublicSubscribe(&operations.GetPublicSubscribeParams{Channels: chans}); err != nil {
		delete(e.subscriptions, chans[0])
		return nil, fmt.Errorf("error subscribing to channel: %s", err)
	}

	go func() {
	Loop:
		for {
			select {
			case n := <-c:
				if len(n.Params.Data) < 1 {
					e.errors <- fmt.Errorf("invalid json data: %s", string(n.Params.Data))
					continue
				}
				data := preFilter(n.Params.Data)
				switch byte(data[0]) {
				case '{':
					var ret models.DeribitPriceIndexNotification
					if err := json.Unmarshal(data, &ret); err != nil {
						e.errors <- fmt.Errorf("error decoding notification: %s", err)
					}
					out <- &ret
				case '[':
					var rets []models.DeribitPriceIndexNotification
					if err := json.Unmarshal(data, &rets); err != nil {
						e.errors <- fmt.Errorf("error decoding notification: %s", err)
					}
					for _, ret := range rets {
						out <- &ret
					}
				default:
					e.errors <- fmt.Errorf("invalid json data: %s", string(data))
				}
			case <-e.stop:
				break Loop
			}
		}
	}()
	return out, nil
}

// SubscribeMarkPriceOptions subscribes to the markprice.options.{index_name} channel
func (e *Exchange) SubscribeMarkPriceOptions(index_name string) (chan *models.MarkpriceOptionsNotification, error) {
	chans := []string{fmt.Sprintf("markprice.options.%s", index_name)}

	c := make(chan *RPCNotification)
	out := make(chan *models.MarkpriceOptionsNotification)
	sub := &RPCSubscription{Data: c, Channel: chans[0]}
	e.subscriptions[chans[0]] = sub

	client := e.Client()
	if _, err := client.GetPublicSubscribe(&operations.GetPublicSubscribeParams{Channels: chans}); err != nil {
		delete(e.subscriptions, chans[0])
		return nil, fmt.Errorf("error subscribing to channel: %s", err)
	}

	go func() {
	Loop:
		for {
			select {
			case n := <-c:
				if len(n.Params.Data) < 1 {
					e.errors <- fmt.Errorf("invalid json data: %s", string(n.Params.Data))
					continue
				}
				data := preFilter(n.Params.Data)
				switch byte(data[0]) {
				case '{':
					var ret models.MarkpriceOptionsNotification
					if err := json.Unmarshal(data, &ret); err != nil {
						e.errors <- fmt.Errorf("error decoding notification: %s", err)
					}
					out <- &ret
				case '[':
					var rets []models.MarkpriceOptionsNotification
					if err := json.Unmarshal(data, &rets); err != nil {
						e.errors <- fmt.Errorf("error decoding notification: %s", err)
					}
					for _, ret := range rets {
						out <- &ret
					}
				default:
					e.errors <- fmt.Errorf("invalid json data: %s", string(data))
				}
			case <-e.stop:
				break Loop
			}
		}
	}()
	return out, nil
}

// SubscribeAnnouncements subscribes to the announcements channel
func (e *Exchange) SubscribeAnnouncements() (chan *models.AnnouncementNotification, error) {
	chans := []string{"announcements"}

	c := make(chan *RPCNotification)
	out := make(chan *models.AnnouncementNotification)
	sub := &RPCSubscription{Data: c, Channel: chans[0]}
	e.subscriptions[chans[0]] = sub

	client := e.Client()
	if _, err := client.GetPrivateSubscribe(&operations.GetPrivateSubscribeParams{Channels: chans}); err != nil {
		delete(e.subscriptions, chans[0])
		return nil, fmt.Errorf("error subscribing to channel: %s", err)
	}

	go func() {
	Loop:
		for {
			select {
			case n := <-c:
				if len(n.Params.Data) < 1 {
					e.errors <- fmt.Errorf("invalid json data: %s", string(n.Params.Data))
					continue
				}
				data := preFilter(n.Params.Data)
				switch byte(data[0]) {
				case '{':
					var ret models.AnnouncementNotification
					if err := json.Unmarshal(data, &ret); err != nil {
						e.errors <- fmt.Errorf("error decoding notification: %s", err)
					}
					out <- &ret
				case '[':
					var rets []models.AnnouncementNotification
					if err := json.Unmarshal(data, &rets); err != nil {
						e.errors <- fmt.Errorf("error decoding notification: %s", err)
					}
					for _, ret := range rets {
						out <- &ret
					}
				default:
					e.errors <- fmt.Errorf("invalid json data: %s", string(data))
				}
			case <-e.stop:
				break Loop
			}
		}
	}()
	return out, nil
}

// SubscribeDeribitPriceRanking subscribes to the deribit_price_ranking.{index_name} channel
func (e *Exchange) SubscribeDeribitPriceRanking(index_name string) (chan *models.DeribitPriceRankingNotification, error) {
	chans := []string{fmt.Sprintf("deribit_price_ranking.%s", index_name)}

	c := make(chan *RPCNotification)
	out := make(chan *models.DeribitPriceRankingNotification)
	sub := &RPCSubscription{Data: c, Channel: chans[0]}
	e.subscriptions[chans[0]] = sub

	client := e.Client()
	if _, err := client.GetPublicSubscribe(&operations.GetPublicSubscribeParams{Channels: chans}); err != nil {
		delete(e.subscriptions, chans[0])
		return nil, fmt.Errorf("error subscribing to channel: %s", err)
	}

	go func() {
	Loop:
		for {
			select {
			case n := <-c:
				if len(n.Params.Data) < 1 {
					e.errors <- fmt.Errorf("invalid json data: %s", string(n.Params.Data))
					continue
				}
				data := preFilter(n.Params.Data)
				switch byte(data[0]) {
				case '{':
					var ret models.DeribitPriceRankingNotification
					if err := json.Unmarshal(data, &ret); err != nil {
						e.errors <- fmt.Errorf("error decoding notification: %s", err)
					}
					out <- &ret
				case '[':
					var rets []models.DeribitPriceRankingNotification
					if err := json.Unmarshal(data, &rets); err != nil {
						e.errors <- fmt.Errorf("error decoding notification: %s", err)
					}
					for _, ret := range rets {
						out <- &ret
					}
				default:
					e.errors <- fmt.Errorf("invalid json data: %s", string(data))
				}
			case <-e.stop:
				break Loop
			}
		}
	}()
	return out, nil
}

// SubscribeTicker subscribes to the ticker.{instrument_name}.{interval} channel
func (e *Exchange) SubscribeTicker(instrument_name, interval string) (chan *models.TickerNotification, error) {
	chans := []string{fmt.Sprintf("ticker.%s.%s", instrument_name, interval)}

	c := make(chan *RPCNotification)
	out := make(chan *models.TickerNotification)
	sub := &RPCSubscription{Data: c, Channel: chans[0]}
	e.subscriptions[chans[0]] = sub

	client := e.Client()
	if _, err := client.GetPublicSubscribe(&operations.GetPublicSubscribeParams{Channels: chans}); err != nil {
		delete(e.subscriptions, chans[0])
		return nil, fmt.Errorf("error subscribing to channel: %s", err)
	}

	go func() {
	Loop:
		for {
			select {
			case n := <-c:
				if len(n.Params.Data) < 1 {
					e.errors <- fmt.Errorf("invalid json data: %s", string(n.Params.Data))
					continue
				}
				data := preFilter(n.Params.Data)
				switch byte(data[0]) {
				case '{':
					var ret models.TickerNotification
					if err := json.Unmarshal(data, &ret); err != nil {
						e.errors <- fmt.Errorf("error decoding notification: %s", err)
					}
					out <- &ret
				case '[':
					var rets []models.TickerNotification
					if err := json.Unmarshal(data, &rets); err != nil {
						e.errors <- fmt.Errorf("error decoding notification: %s", err)
					}
					for _, ret := range rets {
						out <- &ret
					}
				default:
					e.errors <- fmt.Errorf("invalid json data: %s", string(data))
				}
			case <-e.stop:
				break Loop
			}
		}
	}()
	return out, nil
}

// SubscribeTrades subscribes to the trades.{instrument_name}.{interval} channel
func (e *Exchange) SubscribeTrades(instrument_name, interval string) (chan *models.PublicTrade, error) {
	chans := []string{fmt.Sprintf("trades.%s.%s", instrument_name, interval)}

	c := make(chan *RPCNotification)
	out := make(chan *models.PublicTrade)
	sub := &RPCSubscription{Data: c, Channel: chans[0]}
	e.subscriptions[chans[0]] = sub

	client := e.Client()
	if _, err := client.GetPublicSubscribe(&operations.GetPublicSubscribeParams{Channels: chans}); err != nil {
		delete(e.subscriptions, chans[0])
		return nil, fmt.Errorf("error subscribing to channel: %s", err)
	}

	go func() {
	Loop:
		for {
			select {
			case n := <-c:
				if len(n.Params.Data) < 1 {
					e.errors <- fmt.Errorf("invalid json data: %s", string(n.Params.Data))
					continue
				}
				data := preFilter(n.Params.Data)
				switch byte(data[0]) {
				case '{':
					var ret models.PublicTrade
					if err := json.Unmarshal(data, &ret); err != nil {
						e.errors <- fmt.Errorf("error decoding notification: %s", err)
					}
					out <- &ret
				case '[':
					var rets []models.PublicTrade
					if err := json.Unmarshal(data, &rets); err != nil {
						e.errors <- fmt.Errorf("error decoding notification: %s", err)
					}
					for _, ret := range rets {
						out <- &ret
					}
				default:
					e.errors <- fmt.Errorf("invalid json data: %s", string(data))
				}
			case <-e.stop:
				break Loop
			}
		}
	}()
	return out, nil
}

// SubscribeUserTradesKind subscribes to the user.trades.{kind}.{currency}.{interval} channel
func (e *Exchange) SubscribeUserTradesKind(kind, currency, interval string) (chan *models.UserTrade, error) {
	chans := []string{fmt.Sprintf("user.trades.%s.%s.%s", kind, currency, interval)}

	c := make(chan *RPCNotification)
	out := make(chan *models.UserTrade)
	sub := &RPCSubscription{Data: c, Channel: chans[0]}
	e.subscriptions[chans[0]] = sub

	client := e.Client()
	if _, err := client.GetPrivateSubscribe(&operations.GetPrivateSubscribeParams{Channels: chans}); err != nil {
		delete(e.subscriptions, chans[0])
		return nil, fmt.Errorf("error subscribing to channel: %s", err)
	}

	go func() {
	Loop:
		for {
			select {
			case n := <-c:
				if len(n.Params.Data) < 1 {
					e.errors <- fmt.Errorf("invalid json data: %s", string(n.Params.Data))
					continue
				}
				data := preFilter(n.Params.Data)
				switch byte(data[0]) {
				case '{':
					var ret models.UserTrade
					if err := json.Unmarshal(data, &ret); err != nil {
						e.errors <- fmt.Errorf("error decoding notification: %s", err)
					}
					out <- &ret
				case '[':
					var rets []models.UserTrade
					if err := json.Unmarshal(data, &rets); err != nil {
						e.errors <- fmt.Errorf("error decoding notification: %s", err)
					}
					for _, ret := range rets {
						out <- &ret
					}
				default:
					e.errors <- fmt.Errorf("invalid json data: %s", string(data))
				}
			case <-e.stop:
				break Loop
			}
		}
	}()
	return out, nil
}

// SubscribeBookGroup subscribes to the book.{instrument_name}.{group}.{depth}.{interval} channel
func (e *Exchange) SubscribeBookGroup(instrument_name, group, depth, interval string) (chan *models.BookNotification, error) {
	chans := []string{fmt.Sprintf("book.%s.%s.%s.%s", instrument_name, group, depth, interval)}

	c := make(chan *RPCNotification)
	out := make(chan *models.BookNotification)
	sub := &RPCSubscription{Data: c, Channel: chans[0]}
	e.subscriptions[chans[0]] = sub

	client := e.Client()
	if _, err := client.GetPublicSubscribe(&operations.GetPublicSubscribeParams{Channels: chans}); err != nil {
		delete(e.subscriptions, chans[0])
		return nil, fmt.Errorf("error subscribing to channel: %s", err)
	}

	go func() {
	Loop:
		for {
			select {
			case n := <-c:
				if len(n.Params.Data) < 1 {
					e.errors <- fmt.Errorf("invalid json data: %s", string(n.Params.Data))
					continue
				}
				data := preFilter(n.Params.Data)
				switch byte(data[0]) {
				case '{':
					var ret models.BookNotification
					if err := json.Unmarshal(data, &ret); err != nil {
						e.errors <- fmt.Errorf("error decoding notification: %s", err)
					}
					out <- &ret
				case '[':
					var rets []models.BookNotification
					if err := json.Unmarshal(data, &rets); err != nil {
						e.errors <- fmt.Errorf("error decoding notification: %s", err)
					}
					for _, ret := range rets {
						out <- &ret
					}
				default:
					e.errors <- fmt.Errorf("invalid json data: %s", string(data))
				}
			case <-e.stop:
				break Loop
			}
		}
	}()
	return out, nil
}

// SubscribeQuote subscribes to the quote.{instrument_name} channel
func (e *Exchange) SubscribeQuote(instrument_name string) (chan *models.QuoteNotification, error) {
	chans := []string{fmt.Sprintf("quote.%s", instrument_name)}

	c := make(chan *RPCNotification)
	out := make(chan *models.QuoteNotification)
	sub := &RPCSubscription{Data: c, Channel: chans[0]}
	e.subscriptions[chans[0]] = sub

	client := e.Client()
	if _, err := client.GetPublicSubscribe(&operations.GetPublicSubscribeParams{Channels: chans}); err != nil {
		delete(e.subscriptions, chans[0])
		return nil, fmt.Errorf("error subscribing to channel: %s", err)
	}

	go func() {
	Loop:
		for {
			select {
			case n := <-c:
				if len(n.Params.Data) < 1 {
					e.errors <- fmt.Errorf("invalid json data: %s", string(n.Params.Data))
					continue
				}
				data := preFilter(n.Params.Data)
				switch byte(data[0]) {
				case '{':
					var ret models.QuoteNotification
					if err := json.Unmarshal(data, &ret); err != nil {
						e.errors <- fmt.Errorf("error decoding notification: %s", err)
					}
					out <- &ret
				case '[':
					var rets []models.QuoteNotification
					if err := json.Unmarshal(data, &rets); err != nil {
						e.errors <- fmt.Errorf("error decoding notification: %s", err)
					}
					for _, ret := range rets {
						out <- &ret
					}
				default:
					e.errors <- fmt.Errorf("invalid json data: %s", string(data))
				}
			case <-e.stop:
				break Loop
			}
		}
	}()
	return out, nil
}

// SubscribeUserOrdersKind subscribes to the user.orders.{kind}.{currency}.{interval} channel
func (e *Exchange) SubscribeUserOrdersKind(kind, currency, interval string) (chan *models.Order, error) {
	chans := []string{fmt.Sprintf("user.orders.%s.%s.%s", kind, currency, interval)}

	c := make(chan *RPCNotification)
	out := make(chan *models.Order)
	sub := &RPCSubscription{Data: c, Channel: chans[0]}
	e.subscriptions[chans[0]] = sub

	client := e.Client()
	if _, err := client.GetPrivateSubscribe(&operations.GetPrivateSubscribeParams{Channels: chans}); err != nil {
		delete(e.subscriptions, chans[0])
		return nil, fmt.Errorf("error subscribing to channel: %s", err)
	}

	go func() {
	Loop:
		for {
			select {
			case n := <-c:
				if len(n.Params.Data) < 1 {
					e.errors <- fmt.Errorf("invalid json data: %s", string(n.Params.Data))
					continue
				}
				data := preFilter(n.Params.Data)
				switch byte(data[0]) {
				case '{':
					var ret models.Order
					if err := json.Unmarshal(data, &ret); err != nil {
						e.errors <- fmt.Errorf("error decoding notification: %s", err)
					}
					out <- &ret
				case '[':
					var rets []models.Order
					if err := json.Unmarshal(data, &rets); err != nil {
						e.errors <- fmt.Errorf("error decoding notification: %s", err)
					}
					for _, ret := range rets {
						out <- &ret
					}
				default:
					e.errors <- fmt.Errorf("invalid json data: %s", string(data))
				}
			case <-e.stop:
				break Loop
			}
		}
	}()
	return out, nil
}

// SubscribeUserPortfolio subscribes to the user.portfolio.{currency} channel
func (e *Exchange) SubscribeUserPortfolio(currency string) (chan *models.UserPortfolioNotification, error) {
	chans := []string{fmt.Sprintf("user.portfolio.%s", currency)}

	c := make(chan *RPCNotification)
	out := make(chan *models.UserPortfolioNotification)
	sub := &RPCSubscription{Data: c, Channel: chans[0]}
	e.subscriptions[chans[0]] = sub

	client := e.Client()
	if _, err := client.GetPrivateSubscribe(&operations.GetPrivateSubscribeParams{Channels: chans}); err != nil {
		delete(e.subscriptions, chans[0])
		return nil, fmt.Errorf("error subscribing to channel: %s", err)
	}

	go func() {
	Loop:
		for {
			select {
			case n := <-c:
				if len(n.Params.Data) < 1 {
					e.errors <- fmt.Errorf("invalid json data: %s", string(n.Params.Data))
					continue
				}
				data := preFilter(n.Params.Data)
				switch byte(data[0]) {
				case '{':
					var ret models.UserPortfolioNotification
					if err := json.Unmarshal(data, &ret); err != nil {
						e.errors <- fmt.Errorf("error decoding notification: %s", err)
					}
					out <- &ret
				case '[':
					var rets []models.UserPortfolioNotification
					if err := json.Unmarshal(data, &rets); err != nil {
						e.errors <- fmt.Errorf("error decoding notification: %s", err)
					}
					for _, ret := range rets {
						out <- &ret
					}
				default:
					e.errors <- fmt.Errorf("invalid json data: %s", string(data))
				}
			case <-e.stop:
				break Loop
			}
		}
	}()
	return out, nil
}
