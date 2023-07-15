package symbols

type Symbols struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

var (
	// EthBusd ETHERIUM Value in BUSD
	EthBusd = &Symbols{Name: "eth_busd", Value: "ETHBUSD"}
	// BtcBusd BITCOIN Value in BUSD
	BtcBusd = &Symbols{Name: "btc_busd", Value: "BTCBUSD"}
	// EthUsdt ETHERIUM Value in USDT
	EthUsdt = &Symbols{Name: "eth_usdt", Value: "ETHUSDT"}
	// BtcUsdt BITCOIN Value in USDT
	BtcUsdt = &Symbols{Name: "btc_usdt", Value: "BTCUSDT"}
	// MaticBusd POLYHON MATIC Value in BUSD
	MaticBusd = &Symbols{Name: "matic_busd", Value: "MATICBUSD"}
	// SolUsdt POLYHON MATIC Value in BUSD
	SolUsdt = &Symbols{Name: "sol_usdt", Value: "SOLUSDT"}
	// MaticUsdt POLYHON MATIC Value in USDT
	MaticUsdt = &Symbols{Name: "matic_usdt", Value: "MATICUSDT"}
	// BnbBusd  POLYHON MATIC Value in BUSD
	BnbBusd = &Symbols{Name: "bnb_busd", Value: "BNBBUSD"}
	// AdaUsdt CARDANO Value in USDT
	AdaUsdt = &Symbols{Name: "ada_usdt", Value: "ADAUSDT"}
	// BnbUsdt BNB Value in USDT
	BnbUsdt = &Symbols{Name: "bnb_usdt", Value: "BNBUSDT"}
)
