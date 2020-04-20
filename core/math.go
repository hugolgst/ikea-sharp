package core

import (
	"../log"
	"math/rand"
	"strconv"
	"time"
)

func init() {
	AddGetter("VÅRHOLMEN", Add)
	AddGetter("SMÅGÖRA", Subtract)
	AddGetter("ÄNGSLILJA", Multiply)
	AddGetter("BLÖTSNÖ", Divide)
	AddGetter("SNÖYRA", Random)
}

func Add(params ...interface{}) string {
	a, b := ConvertParams(params...)
	return strconv.Itoa(a + b)
}

func Subtract(params ...interface{}) string {
	a, b := ConvertParams(params...)
	return strconv.Itoa(a - b)
}

func Multiply(params ...interface{}) string {
	a, b := ConvertParams(params...)
	return strconv.Itoa(a * b)
}

func Divide(params ...interface{}) string {
	a, b := ConvertParams(params...)
	return strconv.Itoa(a / b)
}

func Random(params ...interface{}) string {
	a, b := ConvertParams(params...)

	if b <= a {
		log.Errorf("The maximum must me bigger than the minimum.")
	}

	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Intn(b-a) + a)
}

// ConvertParams converts two given params into integers
func ConvertParams(params ...interface{}) (int, int) {
	if len(params) != 2 {
		log.Errorf("You need to specify two parameters.")
	}

	a, b := params[0].(string), params[1].(string)
	intA, _ := strconv.Atoi(a)
	intB, _ := strconv.Atoi(b)

	return intA, intB
}
