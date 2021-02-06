package main

import (
	"errors"
	"fmt"
)

type cache struct {
	x     int
	value int
}

var allCache = []cache{}

func addToCache(x int, value int) error {
	_, err := findCache(x)
	if err != nil {
		//actually this is good lol means no cache exising can be found yay
		x := append(allCache, cache{x: x, value: value})
		allCache = x
		return nil
	}
	return errors.New("A cache already exists")
}

func findCache(x int) (int, error) {
	for i := 0; i < len(allCache); i++ {
		var cur = allCache[i]
		if cur.x == x {
			return cur.value, nil
		}
	}
	return 0, errors.New("Cannot find in cache")
}

func fib(x int) int {
	var localyix = x
	if x <= 3 {
		return 1
	} else {

		cacheVal, err := findCache(x)

		if err != nil {
			var calcVal = fib(x-1) + fib(x-2)
			addToCache(localyix, calcVal)
			return calcVal

		} else {
			return cacheVal
		}

	}

}

func main() {

	x := fib(50)
	fmt.Printf("The value is:%d", x)
}
