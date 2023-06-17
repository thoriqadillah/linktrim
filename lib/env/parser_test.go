package env

import (
	"bytes"
	"testing"
	"time"
)

func TestParseInt(t *testing.T) {
	key := "11"
	parsed := parse(key).ToInt()
	expected := 11

	if expected != parsed {
		t.Errorf("error parsing ToInt() function. expected %d, got %d", expected, parsed)
	}

}

func TestParseString(t *testing.T) {
	key := "singo"
	parsed := parse(key).ToString()
	expected := "singo"

	if expected != parsed {
		t.Errorf("error parsing ToInt() function. expected %v, got %v", expected, parsed)
	}
}

func TestParseBool(t *testing.T) {
	key := "false"
	parsed := parse(key).ToBool()
	expected := false

	if expected != parsed {
		t.Errorf("error parsing ToInt() function. expected %v, got %v", expected, parsed)
	}
}

func TestParseFLoat(t *testing.T) {
	key := "0.05"
	parsed := parse(key).ToFloat()
	expected := 0.05

	if expected != parsed {
		t.Errorf("error parsing ToInt() function. expected %v, got %v", expected, parsed)
	}
}

func TestParseDuration(t *testing.T) {
	key := "24h"
	parsed := parse(key).ToDuration()
	expected := time.Duration(24 * time.Hour)

	if expected.Hours() != parsed.Hours() {
		t.Errorf("error parsing ToInt() function. expected %v, got %v", expected, parsed)
	}
}

func TestParseBytes(t *testing.T) {
	key := "singo"
	parsed := parse(key).ToBytes()
	expected := []byte("singo")

	if !bytes.Equal(parsed, expected) {
		t.Errorf("error parsing ToInt() function. expected %v, got %v", expected, parsed)
	}
}
