package set

import (
	"fmt"
	"math/rand"
	"strings"
)

type Set[T comparable] map[T]bool

/*
Constructs a new set
*/
func NewSet[T comparable]() Set[T] {
	return Set[T]{}
}

/*
Inserts an element in a set
*/
func (set Set[T]) Add(element T) {
	set[element] = true
}

/*
Returns a boolean asserting wether an element is present in a set
*/
func (set Set[T]) Has(element T) bool {
	_, exists := set[element]

	return exists
}

/*
Returns the intersection of 2 sets
*/
func (setA Set[T]) Intersection(setB Set[T]) Set[T] {
	result := NewSet[T]()
	for _, element := range setA.ToSlice() {
		if setB.Has(element) {
			result.Add(element)
		}
	}
	return result
}

/*
Returns set elements as a slice
*/
func (set Set[T]) ToSlice() []T {
	result := make([]T, 0, len(set))
	for element := range set {
		result = append(result, element)
	}
	return result
}

/*
Returns a string representing a set and its elements
*/
func (set Set[T]) ToString() string {
	return fmt.Sprintf("%v", set.ToSlice())
}

/*
Returns a string representing its elements (without square brackets)
*/
func (set Set[T]) ToElementsString() string {
	elements := set.ToSlice()
	elementsStrings := make([]string, len(elements))

	for i, element := range elements {
		elementsStrings[i] = fmt.Sprint(element)
	}

	result := strings.Join(elementsStrings, " ")

	return result
}

/*
Get a new set of random positive integers
*/
func GetRandIntSet(setLength int, max int) Set[int] {
	generatedNumbers := NewSet[int]()

	for len(generatedNumbers) < setLength {
		// Get an integer >= 0 and < max, add 1 to get a positive int including max
		randInt := rand.Intn(max) + 1

		if !generatedNumbers.Has(randInt) {
			generatedNumbers.Add(randInt)
		}
	}

	return generatedNumbers
}
