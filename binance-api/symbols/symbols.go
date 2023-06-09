package symbols

type Symbols struct {
	Value string
}

var (
	// EthBusd ETHERIUM Value in BUSD
	EthBusd = &Symbols{Value: "ETHBUSD"}
	// BtcBusd BITCOIN Value in BUSD
	BtcBusd = &Symbols{Value:"BTCBUSD"}
	// MaticBusd POLYHON MATIC Value in BUSD
	MaticBusd = &Symbols{Value:"MATICBUSD"}
	// BnbBusd  POLYHON MATIC Value in BUSD
	BnbBusd = &Symbols{Value:"BNBBUSD"}
	// AdaBusd CARDANO Value in BUSD
	AdaBusd = &Symbols{Value:"ADABUSD"}
)
