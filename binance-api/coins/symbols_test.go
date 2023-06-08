package coins

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEthBusd(t *testing.T) {
	c := require.New(t)
	c.Equal("ETHBUSD", EthBusd)
}

func TestBtcBusd(t *testing.T) {
	c := require.New(t)
	c.Equal("BTCBUSD", BtcBusd)
}

func TestMaticBusd(t *testing.T) {
	c := require.New(t)
	c.Equal("MATICBUSD", MaticBusd)
}

func TestBnbBusd(t *testing.T) {
	c := require.New(t)
	c.Equal("BNBBUSD", BnbBusd)
}

func TestAdaBusd(t *testing.T) {
	c := require.New(t)
	c.Equal("ADABUSD", AdaBusd)
}