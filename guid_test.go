package goutil

import (
	"fmt"
	"testing"
)

func TestGetGUID(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(GetGUID())
	}
}

func TestGetXUID(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(GetXUID())
	}
}

func TestGetSnowFlakeID(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(GetSnowFlakeID())
	}
}