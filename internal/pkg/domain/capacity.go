package domain

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type CapacityUnit rune

const (
	Byte     CapacityUnit = 'B'
	Kilobyte CapacityUnit = 'K'
	Megabyte CapacityUnit = 'M'
	Gigabyte CapacityUnit = 'G'
	Terabyte CapacityUnit = 'T'
)

var AllCapacityUnits = []CapacityUnit{Byte, Kilobyte, Megabyte, Gigabyte, Terabyte}

var CapacityUnitsSize = map[CapacityUnit]uint64{
	Byte:     1,
	Kilobyte: 1024,
	Megabyte: 1024 * 1024,
	Gigabyte: 1024 * 1024 * 1024,
	Terabyte: 1024 * 1024 * 1024 * 1024,
}

type Capacity uint64

func ParseCapacity(s string) (Capacity, error) {
	var amountBuilder strings.Builder
	var unit CapacityUnit

	for _, c := range s {
		if unicode.IsDigit(c) {
			amountBuilder.WriteRune(c)
		} else {
			unit = CapacityUnit(c)
			break
		}
	}

	amount, err := strconv.ParseUint(amountBuilder.String(), 10, 64)
	if err != nil {
		return 0, err
	}

	unitSize, valid := CapacityUnitsSize[unit]
	if !valid {
		return 0, fmt.Errorf("invalid capacity unit: %s, the valid ones are: %v", string(unit), AllCapacityUnits)
	}

	return Capacity(amount * unitSize), nil
}
