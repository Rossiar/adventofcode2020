package main

import "testing"

func Test_requiredFields(t *testing.T) {
	if requiredFields[BirthYear]("2003") {
		t.Errorf("byr invalid")
	}

	if requiredFields[Height]("190in") || requiredFields[Height]("190") {
		t.Errorf("hgt invalid")
	}

	if requiredFields[HairColour]("#123abz") || requiredFields[HairColour]("123abc") {
		t.Errorf("hcl invalid")
	}

	if requiredFields[EyeColour]("wat") {
		t.Errorf("ecl invalid")
	}

	if requiredFields[PassportId]("0123456789") {
		t.Errorf("pid invalid")
	}

}
