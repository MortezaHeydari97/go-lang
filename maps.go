package main

import "fmt"

type Coordinates struct {
	Lat, Long float64
}

func main() {
	/*
	 * A map, maps keys to values
	 * The zero of a map is nil.
	 * A nil map has no keys, nor can keys be added.
	 * The `make` function returns a map of the given type initialized and ready to use.
	*/

	gps := make(map[string]Coordinates)
	fmt.Println("test => ", gps)

	gps["Bell Labs"] = Coordinates{40.68433, -74.39967}


	// Map literals continued
	gps = map[string]Coordinates{
		"Google": Coordinates{37.42202, -122.08408},
		"Bell Labs": Coordinates{40.68433, -74.39967},
	}

	fmt.Println(gps)


	// Mutating maps

	// Insert of update an element in map:
	m[key] = elem
	
	// Retrieve an element:
	elem = m[key]

	// Delete an element:
	delete(m, key)

	// Test that a key is present with a two-value assignment:
	elem, ok = m[key]

	/*
	 * If key is in m, ok is true. If not, ok is false.
	 * If key is not in the map, then elem is the zero value for the map's element type.
	 * Note: If elem or ok have not yet been declared you could use a short declaration form:
	*/
	elem, ok := m[key]
}