package samples

func poc() {
	area = meter(4) * meter(4)                                 // l * l = l^2 or area
	volume = meter(4) * meter(4) * meter(4)                    // l * l * l = l^3 or volume
	volume = area * meter(4)                                   // l ^ 2 * l = l^3 or volume
	volume = metersSquared(4) * meter(4)                       // l ^ 2 * l = l^3 or volume
	length4thPower = meter(4) * meter(4) * meter(4) * meter(4) // l * l * l * l = l^4
	volume.ToCubicFeet()                                       // convert from SI units to cubic feet
}

/**
* UNIT
* length
* NAMES
* meter IS BASE
* foot IS BASE * 3.28084
*
* UNIT
* temperature
* NAMES
* celsius IS BASE + 273.15
* fahrenheit IS ((BASE - 32) * 5 / 9) + 273.15
 */

/*
* FUNCTION
* IdealGasLaw
* INPUTS
* P AS pressure
* V AS volume
* n AS mole
* R AS (pressure * volume) / (mole * temperature)
* RETURNS
* (P * V) / (n * R) AS temperatue
 */
