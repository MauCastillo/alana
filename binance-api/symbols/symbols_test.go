package symbols

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEthBusd(t *testing.T) {
	c := require.New(t)
	c.Equal("ETHBUSD", EthBusd.Value)
}

func TestBtcBusd(t *testing.T) {
	c := require.New(t)
	c.Equal("BTCBUSD", BtcBusd.Value)
}

func TestMaticBusd(t *testing.T) {
	c := require.New(t)
	c.Equal("MATICBUSD", MaticBusd.Value)
}

func TestBnbBusd(t *testing.T) {
	c := require.New(t)
	c.Equal("BNBBUSD", BnbBusd.Value)
}

func TestAdaBusd(t *testing.T) {
	c := require.New(t)
	c.Equal("ADABUSD", AdaBusd.Value)
}