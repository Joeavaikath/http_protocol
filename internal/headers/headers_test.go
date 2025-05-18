package headers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHeadersParse(t *testing.T) {

	headers := NewHeaders()
	data := []byte("Host: localhost:42069\r\n\r\n")
	n, done, err := headers.Parse(data)
	require.NoError(t, err)
	require.NotNil(t, headers)
	assert.Equal(t, "localhost:42069", headers["host"])
	assert.Equal(t, 23, n)
	assert.False(t, done)

	// Test: Invalid spacing header
	headers = NewHeaders()
	data = []byte("       Host : localhost:42069       \r\n\r\n")
	n, done, err = headers.Parse(data)
	require.Error(t, err)
	assert.Equal(t, 0, n)
	assert.False(t, done)

	// Test: Valid single header
	headers = NewHeaders()
	data = []byte("host: localhost123 \r\n\r\n")
	n, done, err = headers.Parse(data)
	require.NoError(t, err)
	require.NotNil(t, headers)
	require.Equal(t, "localhost123", headers["host"])
	assert.False(t, done)

	// Test: Valid single header w/ whitespace
	headers = NewHeaders()
	data = []byte(" host: localhost123 \r\n\r\n")
	n, done, err = headers.Parse(data)
	require.NoError(t, err)
	require.NotNil(t, headers)
	require.Equal(t, "localhost123", headers["host"])
	assert.False(t, done)

	// Test: Valid two headers
	headers = NewHeaders()

	data = []byte(" host: localhost123 \r\n hosty: localhost1234\r\n\r\n")

	n, done, err = headers.Parse(data)
	require.NoError(t, err)
	require.NotNil(t, headers)
	require.Equal(t, "localhost123", headers["host"])

	n, done, err = headers.Parse(data[n:])
	require.NoError(t, err)
	require.NotNil(t, headers)
	require.Equal(t, "localhost1234", headers["hosty"])
	assert.False(t, done)

}
