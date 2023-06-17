package env

import (
	"strconv"
	"time"
)

type parser struct {
	val string
}

func parse(val string) *parser {
	return &parser{
		val: val,
	}
}

func (p *parser) ToInt(def ...int) int {
	v := 0
	if len(def) > 0 {
		v = def[0]
	}

	val, err := strconv.Atoi(p.val)
	if err != nil {
		return v
	}

	return val
}

func (p *parser) ToString(def ...string) string {
	v := ""
	if len(def) > 0 {
		v = def[0]
	}

	if p.val == "" {
		return v
	}

	return p.val
}

func (p *parser) ToBool(def ...bool) bool {
	v := false
	if len(def) > 0 {
		v = def[0]
	}

	val, err := strconv.ParseBool(p.val)
	if err != nil {
		return v
	}

	return val
}

func (p *parser) ToFloat(def ...float64) float64 {
	v := 0.0
	if len(def) > 0 {
		v = def[0]
	}

	val, err := strconv.ParseFloat(p.val, 64)
	if err != nil {
		return v
	}

	return val
}

func (p *parser) ToBytes(def ...[]byte) []byte {
	v := []byte{}
	if len(def) > 0 {
		v = def[0]
	}

	if p.val == "" {
		return v
	}

	return []byte(p.val)
}

func (p *parser) ToDuration(def ...time.Duration) time.Duration {
	v := time.Duration(0)
	if len(def) > 0 {
		v = def[0]
	}

	val, err := time.ParseDuration(p.val)
	if err != nil {
		return v
	}

	return val
}
