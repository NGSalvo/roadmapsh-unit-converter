package tests

import (
	"fmt"
	"testing"
	"unit-converter/services"

	"github.com/stretchr/testify/assert"
)

func TestConverter(t *testing.T) {
	asserts := assert.New(t)

	tests := []struct {
		name      string
		unitType  services.UnitType
		fromUnit  services.Unit
		toUnit    services.Unit
		value     float64
		expected  float64
		expectErr bool
	}{
		{
			name:      "✅ celsius to fahrenheit",
			unitType:  services.Temperature,
			fromUnit:  services.Celsius,
			toUnit:    services.Fahrenheit,
			value:     0,
			expected:  32,
			expectErr: false,
		},
		{
			name:      "✅ fahrenheit to celsius",
			unitType:  services.Temperature,
			fromUnit:  services.Fahrenheit,
			toUnit:    services.Celsius,
			value:     32,
			expected:  0,
			expectErr: false,
		},
		{
			name:      "✅ kelvin to celsius",
			unitType:  services.Temperature,
			fromUnit:  services.Celsius,
			toUnit:    services.Kelvin,
			value:     0,
			expected:  273.15,
			expectErr: false,
		},
		{
			name:      "✅ kelvin to fahrenheit",
			unitType:  services.Temperature,
			fromUnit:  services.Kelvin,
			toUnit:    services.Celsius,
			value:     273.15,
			expected:  0,
			expectErr: false,
		},
		{
			name:      "❌ invalid unit type",
			unitType:  "invalid",
			fromUnit:  services.Celsius,
			toUnit:    services.Fahrenheit,
			value:     0,
			expected:  0,
			expectErr: true,
		},
		{
			name:      "❌ invalid from unit",
			unitType:  services.Temperature,
			fromUnit:  "invalid",
			toUnit:    services.Fahrenheit,
			value:     0,
			expected:  0,
			expectErr: true,
		},
		{
			name:      "❌ invalid to unit",
			unitType:  services.Temperature,
			fromUnit:  services.Celsius,
			toUnit:    "invalid",
			value:     0,
			expected:  0,
			expectErr: true,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf(test.name, test.unitType, test.fromUnit, test.toUnit, test.value), func(t *testing.T) {
			actual, err := services.Convert(test.unitType, test.fromUnit, test.toUnit, test.value)
			asserts.Equal(test.expected, actual)
			asserts.Equal(test.expectErr, err != nil)
		})
	}
}

func TestLengthConverter(t *testing.T) {
	asserts := assert.New(t)

	tests := []struct {
		name      string
		unitType  services.UnitType
		fromUnit  services.Unit
		toUnit    services.Unit
		value     float64
		expected  float64
		expectErr bool
	}{
		{
			name:      "✅ meters to feet",
			unitType:  services.Length,
			fromUnit:  services.Meters,
			toUnit:    services.Feet,
			value:     100,
			expected:  328.08,
			expectErr: false,
		},
		{
			name:      "✅ feet to meters",
			unitType:  services.Length,
			fromUnit:  services.Feet,
			toUnit:    services.Meters,
			value:     328.08,
			expected:  100,
			expectErr: false,
		},
		{
			name:      "✅ meters to kilometers",
			unitType:  services.Length,
			fromUnit:  services.Meters,
			toUnit:    services.Kilometers,
			value:     100,
			expected:  0.1,
			expectErr: false,
		},
		{
			name:      "✅ kilometers to meters",
			unitType:  services.Length,
			fromUnit:  services.Kilometers,
			toUnit:    services.Meters,
			value:     0.1,
			expected:  100,
			expectErr: false,
		},
		{
			name:      "✅ kilometers to feet",
			unitType:  services.Length,
			fromUnit:  services.Kilometers,
			toUnit:    services.Feet,
			value:     0.1,
			expected:  328.08,
			expectErr: false,
		},
		{
			name:      "✅ feet to kilometers",
			unitType:  services.Length,
			fromUnit:  services.Feet,
			toUnit:    services.Kilometers,
			value:     328.08,
			expected:  0.1,
			expectErr: false,
		},
		{
			name:      "✅ kilometers to miles",
			unitType:  services.Length,
			fromUnit:  services.Kilometers,
			toUnit:    services.Miles,
			value:     100,
			expected:  62.14,
			expectErr: false,
		},
		{
			name:      "✅ miles to kilometers",
			unitType:  services.Length,
			fromUnit:  services.Miles,
			toUnit:    services.Kilometers,
			value:     62.14,
			expected:  100,
			expectErr: false,
		},
		{
			name:      "✅ miles to feet",
			unitType:  services.Length,
			fromUnit:  services.Miles,
			toUnit:    services.Feet,
			value:     1,
			expected:  5280,
			expectErr: false,
		},
		{
			name:      "✅ feet to miles",
			unitType:  services.Length,
			fromUnit:  services.Feet,
			toUnit:    services.Miles,
			value:     5280,
			expected:  1,
			expectErr: false,
		},

		{
			name:      "❌ invalid unit type",
			unitType:  "invalid",
			fromUnit:  services.Meters,
			toUnit:    services.Feet,
			value:     100,
			expected:  0,
			expectErr: true,
		},
		{
			name:      "❌ invalid from unit",
			unitType:  services.Length,
			fromUnit:  "invalid",
			toUnit:    services.Feet,
			value:     100,
			expected:  0,
			expectErr: true,
		},
		{
			name:      "❌ invalid to unit",
			unitType:  services.Length,
			fromUnit:  services.Meters,
			toUnit:    "invalid",
			value:     100,
			expected:  0,
			expectErr: true,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf(test.name, test.unitType, test.fromUnit, test.toUnit, test.value), func(t *testing.T) {
			actual, err := services.Convert(test.unitType, test.fromUnit, test.toUnit, test.value)
			asserts.Equal(test.expected, actual, test.name)
			asserts.Equal(test.expectErr, err != nil)
		})
	}
}

func TestWeightConverter(t *testing.T) {
	asserts := assert.New(t)

	tests := []struct {
		name      string
		unitType  services.UnitType
		fromUnit  services.Unit
		toUnit    services.Unit
		value     float64
		expected  float64
		expectErr bool
	}{
		{
			name:      "✅ grams to ounces",
			unitType:  services.Weight,
			fromUnit:  services.Grams,
			toUnit:    services.Ounces,
			value:     100,
			expected:  3.53,
			expectErr: false,
		},
		{
			name:      "✅ ounces to grams",
			unitType:  services.Weight,
			fromUnit:  services.Ounces,
			toUnit:    services.Grams,
			value:     100,
			expected:  2834.95,
			expectErr: false,
		},
		{
			name:      "✅ grams to pounds",
			unitType:  services.Weight,
			fromUnit:  services.Grams,
			toUnit:    services.Pounds,
			value:     1000,
			expected:  2.2,
			expectErr: false,
		},
		{
			name:      "✅ pounds to grams",
			unitType:  services.Weight,
			fromUnit:  services.Pounds,
			toUnit:    services.Grams,
			value:     10,
			expected:  4535.92,
			expectErr: false,
		},
		{
			name:      "✅ pounds to ounces",
			unitType:  services.Weight,
			fromUnit:  services.Pounds,
			toUnit:    services.Ounces,
			value:     1,
			expected:  16,
			expectErr: false,
		},
		{
			name:      "✅ ounces to pounds",
			unitType:  services.Weight,
			fromUnit:  services.Ounces,
			toUnit:    services.Pounds,
			value:     16,
			expected:  1,
			expectErr: false,
		},
		{
			name:      "✅ grams to kilograms",
			unitType:  services.Weight,
			fromUnit:  services.Grams,
			toUnit:    services.Kilograms,
			value:     10000,
			expected:  10,
			expectErr: false,
		},
		{
			name:      "✅ kilograms to grams",
			unitType:  services.Weight,
			fromUnit:  services.Kilograms,
			toUnit:    services.Grams,
			value:     1,
			expected:  1000,
			expectErr: false,
		},
		{
			name:      "✅ grams to milligrams",
			unitType:  services.Weight,
			fromUnit:  services.Grams,
			toUnit:    services.Milligrams,
			value:     1,
			expected:  1000,
			expectErr: false,
		},
		{
			name:      "✅ milligrams to grams",
			unitType:  services.Weight,
			fromUnit:  services.Milligrams,
			toUnit:    services.Grams,
			value:     1,
			expected:  0,
			expectErr: false,
		},
		{
			name:      "✅ milligrams to kilograms",
			unitType:  services.Weight,
			fromUnit:  services.Milligrams,
			toUnit:    services.Kilograms,
			value:     10000,
			expected:  0.01,
			expectErr: false,
		},
		{
			name:      "✅ kilograms to milligrams",
			unitType:  services.Weight,
			fromUnit:  services.Kilograms,
			toUnit:    services.Milligrams,
			value:     0.1,
			expected:  100000,
			expectErr: false,
		},
		{
			name:      "✅ milligrams to ounces",
			unitType:  services.Weight,
			fromUnit:  services.Milligrams,
			toUnit:    services.Ounces,
			value:     100000,
			expected:  3.53,
			expectErr: false,
		},
		{
			name:      "✅ ounces to milligrams",
			unitType:  services.Weight,
			fromUnit:  services.Ounces,
			toUnit:    services.Milligrams,
			value:     10,
			expected:  283495.2,
			expectErr: false,
		},
		{
			name:      "✅ milligrams to pounds",
			unitType:  services.Weight,
			fromUnit:  services.Milligrams,
			toUnit:    services.Pounds,
			value:     100000,
			expected:  0.22,
			expectErr: false,
		},
		{
			name:      "✅ pounds to milligrams",
			unitType:  services.Weight,
			fromUnit:  services.Pounds,
			toUnit:    services.Milligrams,
			value:     0.6,
			expected:  272155.44,
			expectErr: false,
		},
		{
			name:      "❌ invalid unit type",
			unitType:  "invalid",
			fromUnit:  services.Pounds,
			toUnit:    services.Grams,
			value:     1,
			expected:  0,
			expectErr: true,
		},
		{
			name:      "❌ invalid from unit",
			unitType:  services.Weight,
			fromUnit:  "invalid",
			toUnit:    services.Grams,
			value:     1,
			expected:  0,
			expectErr: true,
		},
		{
			name:      "❌ invalid to unit",
			unitType:  services.Weight,
			fromUnit:  services.Pounds,
			toUnit:    "invalid",
			value:     1,
			expected:  0,
			expectErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := services.Convert(test.unitType, test.fromUnit, test.toUnit, test.value)
			asserts.Equal(test.expected, actual, test.name)
			asserts.Equal(test.expectErr, err != nil)
		})
	}
}
