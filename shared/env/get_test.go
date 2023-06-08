package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetBool(t *testing.T) {
	c := require.New(t)

	envName := "LOCAL_BOOL"

	value := GetBool(envName, true)
	c.True(value)

	os.Setenv(envName, "false")

	value = GetBool(envName, true)
	c.False(value)
}

func TestGetInt64 (t *testing.T) {
	c := require.New(t)
	envName := "LOCAL_INT"

	value := GetInt64(envName, 64)
	c.Equal( int64(64), value)

	os.Setenv(envName, "675")

	value = GetInt64(envName, 8989)
	c.Equal( int64(675), value)

	os.Setenv(envName, "puss") 

	value = GetInt64(envName, 685)
	c.Equal( int64(685), value)

}

func TestGetFloat64 (t *testing.T) {
	c := require.New(t)
	envName := "LOCAL_FLOAT"

	value := GetFloat64(envName, 64)
	c.Equal( float64(64), value)

	os.Setenv(envName, "675")

	value = GetFloat64(envName, 8989)
	c.Equal( float64(675), value)

	os.Setenv(envName, "675n")

	value = GetFloat64(envName, 8989)
	c.Equal( float64(8989), value)

}

func TestGetString (t *testing.T) {
	c := require.New(t)
	envName := "LOCAL_STRING"

	value := GetString(envName, "home")
	c.Equal( "home", value)

	os.Setenv(envName, "cat")

	value = GetString(envName, "dog")
	c.Equal( "cat", value)

}

