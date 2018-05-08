package main

import "testing"

func TestGetAPIKey(t *testing.T) {
	input := `asd
apiKey: 'dwe-1-01-Gq565g-asdf-ng-55-hccnn-2cA3wUralIW',
`
	got := getAPIKey(input)
	if got != "dwe-1-01-Gq565g-asdf-ng-55-hccnn-2cA3wUralIW" {
		t.Fatalf("got: %s\n", got)
	}
}

func TestMatchDay(t *testing.T) {
	input := `currentMatchday: {
		'DFL-COM-000001': 33, 'DFL-COM-000002': 33
	},`

	got := getMatchDay(input)
	if got != "33" {
		t.Fatalf("got: %s\n", got)
	}
}
