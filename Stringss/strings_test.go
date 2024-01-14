package Stringss

import "testing"

func TestCount(t *testing.T) {

	counted := findTimesString("holaaaa", "a")
	var expected int = 4
	if counted != expected {
		t.Errorf("this return %d and needed to return %d ", counted, expected)
	}

}

func TestJointStrings(t *testing.T) {

	stringsJointed := jointStrings("hola", "laura")
	expected := "hola laura"

	if expected != stringsJointed {
		t.Errorf("this return %s and needed to return %s ", stringsJointed, expected)
	}

}

func TestTrimThrowFunction(t *testing.T) {

	stringTrimeed := trimString("hola campeon", "campeon")
	expected := "hola "

	if expected != stringTrimeed {
		t.Errorf("this return -%s- and needed to return -%s- ", stringTrimeed, expected)
	}

}
