package client

import (
	"go5/util"
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log(util.Squrae(5))
	t.Log(util.GetFibonacciSerie(10))
	t.Log(util.Squrae(5))
	t.Log(util.GetFibonacciSerie(10))
}
