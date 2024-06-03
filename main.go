package main

import (
	"fmt"
)

// Single row interface
type Row interface {
	Key() string      // returns a key when searching for a row
	Subtract(Row) Row // subtracts the value1s from the given named row from the current one returning the difference
}

type Rows []Row // to simplify slice naming

// Subtract compares 2 slices by Key()
// - for each matching key calculates initial[x] - reduce[y]
// - if we do not find a matching row initial[x] is returned unchanged
// - return the full list of initial with subtractions applied.
func Subtract[T Rows](initial T, reduce T) T {
	subtraction := make(T, len(reduce))

	// iterate over rows to provide an index by name of the rows slice
	reduceIndexByKey := make(map[string]int)
	for i := range reduce {
		reduceIndexByKey[reduce[i].Key()] = i
	}

	for i := range initial {
		if _, ok := reduceIndexByKey[initial[i].Key()]; ok {
			reduceIndex := reduceIndexByKey[initial[i].Key()]
			subtraction[i] = initial[i].Subtract(reduce[reduceIndex])
		} else {
			subtraction[i] = initial[i]
		}
	}
	return subtraction
}

// ------- test struct --------

type SampleRow struct {
	name   string
	value1 int
	// ...
	valuen int
}
type SampleRows []SampleRow

func (s SampleRow) Key() string {
	return s.name
}

func (s SampleRow) Subtract(other SampleRow) SampleRow {
	return SampleRow{
		name:   s.name,
		value1: s.value1 - other.value1,
		// ...
		valuen: s.valuen - other.valuen,
	}
}

// ------- testing to see if we can use Subtract() on SampleRows --------

func main() {
	a := SampleRows{
		{name: "a", value1: 10 /* ... */, valuen: 0},
		{name: "b", value1: 20 /* ... */, valuen: 0},
		{name: "c", value1: 30 /* ... */, valuen: 0},
	}
	b := SampleRows{
		{name: "c", value1: 3 /* ... */, valuen: 0},
		{name: "b", value1: 2 /* ... */, valuen: 0},
	}

	fmt.Printf("a: %+v\n", a)
	fmt.Printf("b: %+v\n", b)

	// SampleRows does now satisfy Rows (SampleRows missing in whatever.Rows)
	c := Subtract[SampleRows](a, b)
	fmt.Printf("c: %+v\n", c)

	/* Expected output would be something like:
	[]SampleRows{
		{"value1": 10, ..., "valuen": 0},
		{"value1": 18, ..., "valuen": 0},
		{"value1": 27, ..., "valuen": 0},
	}
	*/
}
