// Package vreme provides functions and structures for saving and printing weather data
package vreme

import "fmt"

// Postaja represents a struct with values for ime, kraj and meritve (names of propreties are lowercase so that they cannot be accesed from outside file)
type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

// DodajMeritev is a function that takes a map[string]Postaja, oznaka string and temperatura string. It takes this parameters and adds a new temperatura to already existing postaja named <oznaka>
func DodajMeritev(postaje map[string]Postaja, oznaka string, temperatura float64, minTemp float64, maxTemp float64) {

	// Preverimo, če je ključ v slovarju
	postaja, ok := postaje[oznaka]
	if !ok {
		fmt.Println("Postaja ne obstaja:", oznaka)
		return
	}

	if temperatura < minTemp || temperatura > maxTemp {
		fmt.Println("Temperatura je izven dovoljenega obmocja:", temperatura)
		return
	}

	// Lokalna kopija
	postaja.meritve = append(postaja.meritve, temperatura)
	postaje[oznaka] = postaja
	fmt.Println("V izbrano postajo je bila dodana temperatura")

}

// DodajPostajo is a function that adds a new station name "oznaka" to 'database'
func DodajPostajo(postaje map[string]Postaja, oznaka string, ime string, kraj string) {

	postaje[oznaka] = Postaja{ime, kraj, nil}

}

// statistika is a function that expects <oznaka> of a station and than prints out its statistics (max, min, average temperature)
func statistika(postaje map[string]Postaja, oznaka string, stMeritev int) (najnizja, najvisja, povprecna float64, ok bool) {

	// Preverimo, če je ključ v slovarju
	postaja, ok1 := postaje[oznaka]
	// fmt.Println(postaja.meritve[0])

	if !ok1 {
		fmt.Println("Postaja ne obstaja:", oznaka)
		return 0, 0, 0, false
	}

	// Ima manj kot stMeritev meritve
	if len(postaja.meritve) < stMeritev {
		// fmt.Println("Postaja ima premalo meritev")
		return 0, 0, 0, false
	}

	najnizja = postaja.meritve[0]
	najvisja = postaja.meritve[0]
	for _, v := range postaja.meritve {
		// fmt.Println(i, v)
		povprecna += v

		if v < najnizja {
			najnizja = v
		} else if v > najvisja {
			najvisja = v
		}
	}

	povprecna = povprecna / float64(len(postaja.meritve))
	ok = true

	return
}

// IzpisMeritev prints out all of the data from struct Postaja
func IzpisMeritev(postaje map[string]Postaja) {
	fmt.Println("MERITVE:")

	for l, pos := range postaje {

		fmt.Printf("%s - %s (%s): ", l, pos.ime, pos.kraj)
		fmt.Println(pos.meritve)
	}
}

// IzpisPorocila gives statistics and it's say in climate so what should this temperatures mean
func IzpisPorocila(postaje map[string]Postaja, stMeritev int) {

	for oznaka, pos := range postaje {

		najnizja, najvisja, povp, ok := statistika(postaje, oznaka, stMeritev)

		if !ok {
			fmt.Printf("%s: Premalo meritev:\n", pos.ime)
		} else {
			word := ""
			switch {
			case povp >= 25:
				word = "Vroče"
			case povp >= 15 && povp < 25:
				word = "Toplo"
			case povp >= 5 && povp < 15:
				word = "Hladno"
			case povp < 5:
				word = "Mrzlo"
			}

			fmt.Printf("%s: min %.1f °C, max %.1f °C, povprečje %.1f °C -> %s\n", pos.ime, najnizja, najvisja, povp, word)
		}
	}

}
