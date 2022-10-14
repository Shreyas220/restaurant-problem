package main

import (
	"reflect"
	"testing"
)

func TestRestaurant(t *testing.T) {
	ans, err := restaurant("log.txt")
	checkans := []string{"5", "3", "2"}
	if err != nil && !reflect.DeepEqual(ans, checkans) {
		t.Error("test not working")
	}

}

//this test is to check wheather the function responds with err when given multiple similar input
func TestRestaurant2(t *testing.T) {
	ans, err := restaurant("log2.txt")
	if err == nil && ans != nil {
		t.Error("test not working")
	}
}
