package samples

func poc() {
	area = meter(4) * meter(4)                                 // l * l = l^2 or area
	volume = meter(4) * meter(4) * meter(4)                    // l * l * l = l^3 or volume
	volume = area * meter(4)                                   // l ^ 2 * l = l^3 or volume
	volume = metersSquared(4) * meter(4)                       // l ^ 2 * l = l^3 or volume
	length4thPower = meter(4) * meter(4) * meter(4) * meter(4) // l * l * l * l = l^4
	volume.ToCubicFeet()                                       // convert from SI units to cubic feet
}

/*
* FUNCTION
* IdealGasLaw
* INPUTS
* P as pressure
* V as volume
* n as mole
* R as (pressure * volume) / (mole * temperature)
* RETURNS
* (P * V) / (n * R) as temperatue
 */
