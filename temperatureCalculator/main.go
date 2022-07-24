package main

import "fmt"

type Degree struct {
	degree float32
}

type Calculate interface {
	celsiusToFahrenheit()
	celsiusToKelvin()
}

func (c Degree) celsiusToFahrenheit() float32 {
	f := (float32(c.degree) * 1.8) + 32
	return f
}

func (c Degree) celsiusToKelvin() float32 {
	k := (float32(c.degree)) + 273.15
	return k
}

func main() {
	var c float32
	fmt.Print("Enter the temperature in Celcius:")
	fmt.Scanf("%f", &c)

	var d Degree = Degree{degree: c}
	f := d.celsiusToFahrenheit()
	k := d.celsiusToKelvin()
	fmt.Println("Temperature in Fahrenheit is:", f)
	fmt.Println("Temperature in Kelvin is:", k)
}
