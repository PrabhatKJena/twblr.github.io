package data_structures

import "testing"

func TestMapLikes(t *testing.T) {
	data := map[string][]string{
		"superman":  []string{"flying", "x-ray vision"},
		"spiderman": []string{"swinging", "super hearing", "running"},
		"flash":     []string{"running", "color red"},
	}
	people := findPeopleWithCommonInterest(data, "running")

	if !contains(people, "spiderman") {
		t.Errorf("Expected %v to contain spiderman", people)
	}

	if !contains(people, "flash") {
		t.Errorf("Expected %v to contain flash", people)
	}
}
