package main

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConverter(t *testing.T) {
	asserts := assert.New(t)

	tests := []struct {
		name      string
		unitType  UnitType
		fromUnit  Unit
		toUnit    Unit
		value     float64
		expected  float64
		expectErr bool
	}{
		{
			name:      "✅ celsius to fahrenheit",
			unitType:  Temperature,
			fromUnit:  Celsius,
			toUnit:    Fahrenheit,
			value:     0,
			expected:  32,
			expectErr: false,
		},
		{
			name:      "✅ fahrenheit to celsius",
			unitType:  Temperature,
			fromUnit:  Fahrenheit,
			toUnit:    Celsius,
			value:     32,
			expected:  0,
			expectErr: false,
		},
		{
			name:      "✅ kelvin to celsius",
			unitType:  Temperature,
			fromUnit:  Celsius,
			toUnit:    Kelvin,
			value:     0,
			expected:  273.15,
			expectErr: false,
		},
		{
			name:      "✅ kelvin to fahrenheit",
			unitType:  Temperature,
			fromUnit:  Kelvin,
			toUnit:    Celsius,
			value:     273.15,
			expected:  0,
			expectErr: false,
		},
		{
			name:      "❌ invalid unit type",
			unitType:  "invalid",
			fromUnit:  Celsius,
			toUnit:    Fahrenheit,
			value:     0,
			expected:  0,
			expectErr: true,
		},
		{
			name:      "❌ invalid from unit",
			unitType:  Temperature,
			fromUnit:  "invalid",
			toUnit:    Fahrenheit,
			value:     0,
			expected:  0,
			expectErr: true,
		},
		{
			name:      "❌ invalid to unit",
			unitType:  Temperature,
			fromUnit:  Celsius,
			toUnit:    "invalid",
			value:     0,
			expected:  0,
			expectErr: true,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf(test.name, test.unitType, test.fromUnit, test.toUnit, test.value), func(t *testing.T) {
			actual, err := Convert(test.unitType, test.fromUnit, test.toUnit, test.value)
			asserts.Equal(test.expected, actual)
			asserts.Equal(test.expectErr, err != nil)
		})
	}
}

func TestLengthConverter(t *testing.T) {
	asserts := assert.New(t)

	tests := []struct {
		name      string
		unitType  UnitType
		fromUnit  Unit
		toUnit    Unit
		value     float64
		expected  float64
		expectErr bool
	}{
		{
			name:      "✅ meters to feet",
			unitType:  Length,
			fromUnit:  Meters,
			toUnit:    Feet,
			value:     100,
			expected:  328.08,
			expectErr: false,
		},
		{
			name:      "✅ feet to meters",
			unitType:  Length,
			fromUnit:  Feet,
			toUnit:    Meters,
			value:     328.08,
			expected:  100,
			expectErr: false,
		},
		{
			name:      "✅ meters to kilometers",
			unitType:  Length,
			fromUnit:  Meters,
			toUnit:    Kilometers,
			value:     100,
			expected:  0.1,
			expectErr: false,
		},
		{
			name:      "✅ kilometers to meters",
			unitType:  Length,
			fromUnit:  Kilometers,
			toUnit:    Meters,
			value:     0.1,
			expected:  100,
			expectErr: false,
		},
		{
			name:      "✅ kilometers to feet",
			unitType:  Length,
			fromUnit:  Kilometers,
			toUnit:    Feet,
			value:     0.1,
			expected:  328.08,
			expectErr: false,
		},
		{
			name:      "✅ feet to kilometers",
			unitType:  Length,
			fromUnit:  Feet,
			toUnit:    Kilometers,
			value:     328.08,
			expected:  0.1,
			expectErr: false,
		},
		{
			name:      "✅ kilometers to miles",
			unitType:  Length,
			fromUnit:  Kilometers,
			toUnit:    Miles,
			value:     100,
			expected:  62.14,
			expectErr: false,
		},
		{
			name:      "✅ miles to kilometers",
			unitType:  Length,
			fromUnit:  Miles,
			toUnit:    Kilometers,
			value:     62.14,
			expected:  100,
			expectErr: false,
		},
		{
			name:      "✅ miles to feet",
			unitType:  Length,
			fromUnit:  Miles,
			toUnit:    Feet,
			value:     1,
			expected:  5280,
			expectErr: false,
		},
		{
			name:      "✅ feet to miles",
			unitType:  Length,
			fromUnit:  Feet,
			toUnit:    Miles,
			value:     5280,
			expected:  1,
			expectErr: false,
		},

		{
			name:      "❌ invalid unit type",
			unitType:  "invalid",
			fromUnit:  Meters,
			toUnit:    Feet,
			value:     100,
			expected:  0,
			expectErr: true,
		},
		{
			name:      "❌ invalid from unit",
			unitType:  Length,
			fromUnit:  "invalid",
			toUnit:    Feet,
			value:     100,
			expected:  0,
			expectErr: true,
		},
		{
			name:      "❌ invalid to unit",
			unitType:  Length,
			fromUnit:  Meters,
			toUnit:    "invalid",
			value:     100,
			expected:  0,
			expectErr: true,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf(test.name, test.unitType, test.fromUnit, test.toUnit, test.value), func(t *testing.T) {
			actual, err := Convert(test.unitType, test.fromUnit, test.toUnit, test.value)
			asserts.Equal(test.expected, actual, test.name)
			asserts.Equal(test.expectErr, err != nil)
		})
	}
}

func TestWeightConverter(t *testing.T) {
	asserts := assert.New(t)

	tests := []struct {
		name      string
		unitType  UnitType
		fromUnit  Unit
		toUnit    Unit
		value     float64
		expected  float64
		expectErr bool
	}{
		{
			name:      "✅ grams to ounces",
			unitType:  Weight,
			fromUnit:  Grams,
			toUnit:    Ounces,
			value:     100,
			expected:  3.53,
			expectErr: false,
		},
		{
			name:      "✅ ounces to grams",
			unitType:  Weight,
			fromUnit:  Ounces,
			toUnit:    Grams,
			value:     100,
			expected:  2834.95,
			expectErr: false,
		},
		{
			name:      "✅ grams to pounds",
			unitType:  Weight,
			fromUnit:  Grams,
			toUnit:    Pounds,
			value:     1000,
			expected:  2.2,
			expectErr: false,
		},
		{
			name:      "✅ pounds to grams",
			unitType:  Weight,
			fromUnit:  Pounds,
			toUnit:    Grams,
			value:     10,
			expected:  4535.92,
			expectErr: false,
		},
		{
			name:      "✅ pounds to ounces",
			unitType:  Weight,
			fromUnit:  Pounds,
			toUnit:    Ounces,
			value:     1,
			expected:  16,
			expectErr: false,
		},
		{
			name:      "✅ ounces to pounds",
			unitType:  Weight,
			fromUnit:  Ounces,
			toUnit:    Pounds,
			value:     16,
			expected:  1,
			expectErr: false,
		},
		{
			name:      "✅ grams to kilograms",
			unitType:  Weight,
			fromUnit:  Grams,
			toUnit:    Kilograms,
			value:     10000,
			expected:  10,
			expectErr: false,
		},
		{
			name:      "✅ kilograms to grams",
			unitType:  Weight,
			fromUnit:  Kilograms,
			toUnit:    Grams,
			value:     1,
			expected:  1000,
			expectErr: false,
		},
		{
			name:      "✅ grams to milligrams",
			unitType:  Weight,
			fromUnit:  Grams,
			toUnit:    Milligrams,
			value:     1,
			expected:  1000,
			expectErr: false,
		},
		{
			name:      "✅ milligrams to grams",
			unitType:  Weight,
			fromUnit:  Milligrams,
			toUnit:    Grams,
			value:     1,
			expected:  0,
			expectErr: false,
		},
		{
			name:      "✅ milligrams to kilograms",
			unitType:  Weight,
			fromUnit:  Milligrams,
			toUnit:    Kilograms,
			value:     10000,
			expected:  0.01,
			expectErr: false,
		},
		{
			name:      "✅ kilograms to milligrams",
			unitType:  Weight,
			fromUnit:  Kilograms,
			toUnit:    Milligrams,
			value:     0.1,
			expected:  100000,
			expectErr: false,
		},
		{
			name:      "✅ milligrams to ounces",
			unitType:  Weight,
			fromUnit:  Milligrams,
			toUnit:    Ounces,
			value:     100000,
			expected:  3.53,
			expectErr: false,
		},
		{
			name:      "✅ ounces to milligrams",
			unitType:  Weight,
			fromUnit:  Ounces,
			toUnit:    Milligrams,
			value:     10,
			expected:  283495.2,
			expectErr: false,
		},
		{
			name:      "✅ milligrams to pounds",
			unitType:  Weight,
			fromUnit:  Milligrams,
			toUnit:    Pounds,
			value:     100000,
			expected:  0.22,
			expectErr: false,
		},
		{
			name:      "✅ pounds to milligrams",
			unitType:  Weight,
			fromUnit:  Pounds,
			toUnit:    Milligrams,
			value:     0.6,
			expected:  272155.44,
			expectErr: false,
		},
		{
			name:      "❌ invalid unit type",
			unitType:  "invalid",
			fromUnit:  Pounds,
			toUnit:    Grams,
			value:     1,
			expected:  0,
			expectErr: true,
		},
		{
			name:      "❌ invalid from unit",
			unitType:  Weight,
			fromUnit:  "invalid",
			toUnit:    Grams,
			value:     1,
			expected:  0,
			expectErr: true,
		},
		{
			name:      "❌ invalid to unit",
			unitType:  Weight,
			fromUnit:  Pounds,
			toUnit:    "invalid",
			value:     1,
			expected:  0,
			expectErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := Convert(test.unitType, test.fromUnit, test.toUnit, test.value)
			asserts.Equal(test.expected, actual, test.name)
			asserts.Equal(test.expectErr, err != nil)
		})
	}
}

// ConverterFunc is a function that converts one value to another
type ConverterFunc func(float64) float64

// UnitType defines the category of units (e.g., Temperature, Length, Weight)
type UnitType string

// Unit represents a generic unit for conversion
type Unit string

// Supported unit types
const (
	Temperature UnitType = "temperature"
	Length      UnitType = "length"
	Weight      UnitType = "weight"
)

// Supported units for Temperature
const (
	Celsius    Unit = "celsius"
	Fahrenheit Unit = "fahrenheit"
	Kelvin     Unit = "kelvin"
)

// Supported units for Length
const (
	Meters     Unit = "meters"
	Kilometers Unit = "kilometers"
	Feet       Unit = "feet"
	Yards      Unit = "yards"
	Miles      Unit = "miles"
)

// Supported units for Weight
const (
	Milligrams Unit = "milligrams"
	Grams      Unit = "grams"
	Kilograms  Unit = "kilograms"
	Ounces     Unit = "ounces"
	Pounds     Unit = "pounds"
)

// ConversionTable holds the conversion functions for different unit types
var ConversionTable = map[UnitType]map[Unit]map[Unit]ConverterFunc{
	Temperature: {
		Celsius: {
			Fahrenheit: celsiusToFahrenheit,
			Kelvin:     celsiusToKelvin,
		},
		Fahrenheit: {
			Celsius: fahrenheitToCelsius,
			Kelvin:  fahrenheitToKelvin,
		},
		Kelvin: {
			Celsius:    kelvinToCelsius,
			Fahrenheit: kelvinToFahrenheit,
		},
	},

	Length: {
		Meters: {
			Kilometers: metersToKilometers,
			Feet:       metersToFeet,
			Yards:      metersToYards,
			Miles:      metersToMiles,
		},
		Kilometers: {
			Meters: kilometersToMeters,
			Feet:   kilometersToFeet,
			Yards:  kilometersToYards,
			Miles:  kilometersToMiles,
		},
		Feet: {
			Meters:     feetToMeters,
			Kilometers: feetToKilometers,
			Yards:      feetToYards,
			Miles:      feetToMiles,
		},
		Yards: {
			Meters:     yardsToMeters,
			Kilometers: yardsToKilometers,
			Feet:       yardsToFeet,
			Miles:      yardsToMiles,
		},
		Miles: {
			Meters:     milesToMeters,
			Kilometers: milesToKilometers,
			Feet:       milesToFeet,
			Yards:      milesToYards,
		},
	},

	Weight: {
		Milligrams: {
			Grams:     milligramsToGrams,
			Kilograms: milligramsToKilograms,
			Ounces:    milligramsToOunces,
			Pounds:    milligramsToPounds,
		},
		Grams: {
			Milligrams: gramsToMilligrams,
			Kilograms:  gramsToKilograms,
			Ounces:     gramsToOunces,
			Pounds:     gramsToPounds,
		},
		Kilograms: {
			Milligrams: kilogramsToMilligrams,
			Grams:      kilogramsToGrams,
			Ounces:     kilogramsToOunces,
			Pounds:     kilogramsToPounds,
		},
		Ounces: {
			Milligrams: ouncesToMilligrams,
			Grams:      ouncesToGrams,
			Kilograms:  ouncesToKilograms,
			Pounds:     ouncesToPounds,
		},
		Pounds: {
			Milligrams: poundsToMilligrams,
			Grams:      poundsToGrams,
			Kilograms:  poundsToKilograms,
			Ounces:     poundsToOunces,
		},
	},
}

// Convert performs a conversion between two units of the same type
func Convert(unitType UnitType, fromUnit, toUnit Unit, value float64) (float64, error) {
	if fromUnit == toUnit {
		return value, nil
	}

	conversion, ok := ConversionTable[unitType][fromUnit][toUnit]
	if !ok {
		return 0, fmt.Errorf("conversion from %q to %q not supported", fromUnit, toUnit)
	}

	return math.Round(conversion(value)*100) / 100, nil
}

// Temperature conversion functions
func celsiusToFahrenheit(c float64) float64 { return c*9/5 + 32 }
func celsiusToKelvin(c float64) float64     { return c + 273.15 }
func fahrenheitToCelsius(f float64) float64 { return (f - 32) * 5 / 9 }
func fahrenheitToKelvin(f float64) float64  { return (f + 459.67) * 5 / 9 }
func kelvinToCelsius(k float64) float64     { return k - 273.15 }
func kelvinToFahrenheit(k float64) float64  { return (k * 9 / 5) - 459.67 }

// Length conversion functions
func metersToKilometers(m float64) float64 { return m / 1000 }
func metersToFeet(m float64) float64       { return m * 3.28084 }
func metersToYards(m float64) float64      { return m * 1.09361 }
func metersToMiles(m float64) float64      { return m * 0.000621371 }

func kilometersToMeters(k float64) float64 { return k * 1000 }
func kilometersToFeet(k float64) float64   { return k * 3280.84 }
func kilometersToYards(k float64) float64  { return k * 1093.61 }
func kilometersToMiles(k float64) float64  { return k * 0.621371 }

func feetToMeters(f float64) float64     { return f / 3.28084 }
func feetToKilometers(f float64) float64 { return f / 3280.84 }
func feetToYards(f float64) float64      { return f / 3 }
func feetToMiles(f float64) float64      { return f / 5280 }

func yardsToMeters(y float64) float64     { return y / 1.09361 }
func yardsToKilometers(y float64) float64 { return y / 1093.61 }
func yardsToFeet(y float64) float64       { return y * 3 }
func yardsToMiles(y float64) float64      { return y / 1760 }

func milesToMeters(m float64) float64     { return m * 1609.34 }
func milesToKilometers(m float64) float64 { return m * 1.60934 }
func milesToFeet(m float64) float64       { return m * 5280 }
func milesToYards(m float64) float64      { return m * 1760 }

// Weight conversion functions
func milligramsToGrams(mg float64) float64     { return mg / 1000 }
func milligramsToKilograms(mg float64) float64 { return mg / 1000000 }
func milligramsToOunces(mg float64) float64    { return mg * 0.000035274 }
func milligramsToPounds(mg float64) float64    { return mg * 0.00000220462 }

func gramsToMilligrams(g float64) float64 { return g * 1000 }
func gramsToKilograms(g float64) float64  { return g / 1000 }
func gramsToOunces(g float64) float64     { return g * 0.035274 }
func gramsToPounds(g float64) float64     { return g * 0.00220462 }

func kilogramsToMilligrams(kg float64) float64 { return kg * 1000000 }
func kilogramsToGrams(kg float64) float64      { return kg * 1000 }
func kilogramsToOunces(kg float64) float64     { return kg * 35.274 }
func kilogramsToPounds(kg float64) float64     { return kg * 2.20462 }

func ouncesToMilligrams(oz float64) float64 { return oz * 28349.52 }
func ouncesToGrams(oz float64) float64      { return oz * 28.3495 }
func ouncesToKilograms(oz float64) float64  { return oz * 0.02835 }
func ouncesToPounds(oz float64) float64     { return oz * 0.0625 }

func poundsToMilligrams(lb float64) float64 { return lb * 453592.4 }
func poundsToGrams(lb float64) float64      { return lb * 453.5924 }
func poundsToKilograms(lb float64) float64  { return lb * 0.453592 }
func poundsToOunces(lb float64) float64     { return lb * 16 }
