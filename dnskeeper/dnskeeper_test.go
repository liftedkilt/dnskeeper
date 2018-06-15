package dnskeeper_test

import (
	"testing"

	"github.com/liftedkilt/dnskeeper/dnskeeper"
	"github.com/stretchr/testify/assert"
)

var (
	b bool
	i int
	s string
)

func TestUnMarshal(t *testing.T) {
	assert := assert.New(t)
	raw := `{
		"40t132836m0k6q35lmn3ds82d38uv00d": {
			"content": "127.0.0.1",
			"name": "bar.example.org",
			"proxied": false,
			"ttl": 1,
			"type": "CNAME"
		},
		"bc0179rei87981c6e36y4931b23iwu13": {
			"content": "bar.example.org",
			"name": "foo.example.org",
			"proxied": true,
			"ttl": 1,
			"type": "CNAME"
		}}`
	zone, err := dnskeeper.Unmarshal([]byte(raw))
	assert.Nil(err)

	if assert.NotNil(zone) {
		// Assert that Records are parsed into struct properly
		assert.Equal(2, len(zone), "There should be 4 records")

		// Assert that fields are parsed properly
		assert.Equal("127.0.0.1", zone["40t132836m0k6q35lmn3ds82d38uv00d"].Content, "Content field should match")
		assert.Equal("bar.example.org", zone["40t132836m0k6q35lmn3ds82d38uv00d"].Name, "Name field should match")
		assert.Equal(1, zone["40t132836m0k6q35lmn3ds82d38uv00d"].TTL, "TTL field should match")
		assert.Equal("CNAME", zone["40t132836m0k6q35lmn3ds82d38uv00d"].Type, "Type field should match")
		assert.Equal(false, zone["40t132836m0k6q35lmn3ds82d38uv00d"].Proxied, "Proxied field should match")
		assert.NotEqual("127.0.1.1", zone["40t132836m0k6q35lmn3ds82d38uv00d"].Content, "Content field should not match")
		assert.NotEqual(" bar.example.org", zone["40t132836m0k6q35lmn3ds82d38uv00d"].Name, "Name field should not match")
		assert.NotEqual(0, zone["40t132836m0k6q35lmn3ds82d38uv00d"].TTL, "TTL field should not match")
		assert.NotEqual("A", zone["40t132836m0k6q35lmn3ds82d38uv00d"].Type, "Type field should not match")
		assert.NotEqual(true, zone["40t132836m0k6q35lmn3ds82d38uv00d"].Proxied, "Proxied field should match")
	}

	// Test field validation

	// Content must be a string
	raw = `{
		"40t132836m0k6q35lmn3ds82d38uv00d": {
			"content": 12345
		}}`

	_, err = dnskeeper.Unmarshal([]byte(raw))
	assert.NotEqual(nil, err)

	// Name must be a string
	raw = `{
		"40t132836m0k6q35lmn3ds82d38uv00d": {
			"name": true
		}}`

	_, err = dnskeeper.Unmarshal([]byte(raw))
	assert.NotEqual(nil, err)

	// Proxied must be a bool
	raw = `{
		"40t132836m0k6q35lmn3ds82d38uv00d": {
			"proxied": "false",
		}}`

	_, err = dnskeeper.Unmarshal([]byte(raw))
	assert.NotEqual(nil, err)

	// TTL must be an int
	raw = `{
		"40t132836m0k6q35lmn3ds82d38uv00d": {
			"ttl": "0",
		}}`

	_, err = dnskeeper.Unmarshal([]byte(raw))
	assert.NotEqual(nil, err)

	// Type must be a string
	raw = `{
		"40t132836m0k6q35lmn3ds82d38uv00d": {
			"type": false
		}}`

	_, err = dnskeeper.Unmarshal([]byte(raw))
	assert.NotEqual(nil, err)
}

func TestZoneName(t *testing.T) {
	assert := assert.New(t)

	zoneFile := "/some/path/to/a/file/example.org.json"

	zoneName := dnskeeper.GetZoneName(zoneFile)

	assert.Equal("example.org", zoneName)
}
