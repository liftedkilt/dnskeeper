// Package dnskeeper is a set of utilities for storing and handling Cloudflare DNS records locally
package dnskeeper

import (
	"encoding/json"
	"path/filepath"
	"strings"
)

// DNSRecord is a struct that holds the required fields to construct a Cloudflare record
type DNSRecord struct {
	Content string `json:"content"`
	Name    string `json:"name"`
	Proxied bool   `json:"proxied"`
	TTL     int    `json:"ttl"`
	Type    string `json:"type"`
}

// Zone is a map of Record IDs to a pointer to a DNSRecord struct
type Zone map[string]*DNSRecord

// Unmarshal takes a json array as a slice of bytes as an argument, and returns a Zone struct
func Unmarshal(b []byte) (z Zone, err error) {
	err = json.Unmarshal(b, &z)
	if err != nil {
		return nil, err
	}

	return z, nil
}

// GetZoneName returns the name of the zone that is being operated on by examining the filepath
func GetZoneName(s string) string {
	return strings.TrimSuffix(filepath.Base(s), ".json")
}
