package coins

type Coin struct {
	Value string
}

var (
	// EthBusd ETHERIUM Value in BUSD
	EthBusd = &Coin{Value: "ETHBUSD"}
	// BtcBusd BITCOIN Value in BUSD
	BtcBusd = &Coin{Value:"BTCBUSD"}
	// MaticBusd POLYHON MATIC Value in BUSD
	MaticBusd = &Coin{Value:"MATICBUSD"}
	// BnbBusd  POLYHON MATIC Value in BUSD
	BnbBusd = &Coin{Value:"BNBBUSD"}
	// AdaBusd CARDANO Value in BUSD
	AdaBusd = &Coin{Value:"ADABUSD"}
)
