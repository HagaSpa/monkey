package object

import "testing"

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "Hello World"}
	hello2 := &String{Value: "Hello World"}
	diff1 := &String{Value: "My name is johnny"}
	diff2 := &String{Value: "My name is johnny"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}
	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}
	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("strings with different content have same hash keys")
	}
}

func TestIntegerHashKey(t *testing.T) {
	int1 := &Integer{Value: 1}
	int2 := &Integer{Value: 1}
	diff1 := &Integer{Value: 0}

	if int1.HashKey() != int2.HashKey() {
		t.Errorf("integers with same content different hash keys")
	}
	if int1.HashKey() == diff1.HashKey() {
		t.Errorf("integers with different content have same hash keys")
	}
}

func TestBooleanHashKey(t *testing.T) {
	bool1 := &Boolean{Value: true}
	bool2 := &Boolean{Value: true}
	diff1 := &Boolean{Value: false}

	if bool1.HashKey() != bool2.HashKey() {
		t.Errorf("booleans with same content different hash keys")
	}
	if bool1.HashKey() == diff1.HashKey() {
		t.Errorf("booleans with different content have same hash keys")
	}
}
