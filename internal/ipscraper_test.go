package internal

import (
	"os"
	"testing"
)

func TestScrapeIpV6(t *testing.T) {
	raw := "::1 55234 22"
	os.Setenv("SSH_CLIENT", raw)
}

func TestScrapeIpV4(t *testing.T) {
	raw := "123.123.123.123 55234 22"
	os.Setenv("SSH_CLIENT", raw)
}