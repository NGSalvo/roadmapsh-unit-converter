package services

import (
	"fmt"
	"math"
)

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
