package main

import (
	"encoding/json"
	"testing"
)

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

func TestParseTables(t *testing.T) {
	src := `[{"wettbewerbsId":"DFL-COM-000002","wettbewerb":"2. Bundesliga","spieltagsId":null,"spieltag":null,"eintraege":[{"platzierung":1,"unterPlatzierung":null,"club":"1. FC Nürnberg","clubId":"DFL-CLU-000005"},{"platzierung":2,"unterPlatzierung":null,"club":"Fortuna Düsseldorf","clubId":"DFL-CLU-00000P"},{"platzierung":3,"unterPlatzierung":null,"club":"Holstein Kiel","clubId":"DFL-CLU-000N5P"},{"platzierung":4,"unterPlatzierung":null,"club":"DSC Arminia Bielefeld","clubId":"DFL-CLU-000015"},{"platzierung":5,"unterPlatzierung":null,"club":"SSV Jahn Regensburg","clubId":"DFL-CLU-000011"},{"platzierung":6,"unterPlatzierung":null,"club":"VfL Bochum 1848","clubId":"DFL-CLU-00000S"},{"platzierung":7,"unterPlatzierung":null,"club":"FC Ingolstadt 04","clubId":"DFL-CLU-00000J"},{"platzierung":8,"unterPlatzierung":null,"club":"MSV Duisburg","clubId":"DFL-CLU-00000R"},{"platzierung":9,"unterPlatzierung":null,"club":"1. FC Union Berlin","clubId":"DFL-CLU-00000V"},{"platzierung":10,"unterPlatzierung":null,"club":"FC St. Pauli","clubId":"DFL-CLU-00000H"},{"platzierung":11,"unterPlatzierung":null,"club":"SV Sandhausen","clubId":"DFL-CLU-000012"},{"platzierung":12,"unterPlatzierung":null,"club":"1. FC Heidenheim 1846","clubId":"DFL-CLU-000018"},{"platzierung":13,"unterPlatzierung":null,"club":"SG Dynamo Dresden","clubId":"DFL-CLU-00000N"},{"platzierung":14,"unterPlatzierung":null,"club":"SV Darmstadt 98","clubId":"DFL-CLU-000016"},{"platzierung":15,"unterPlatzierung":null,"club":"FC Erzgebirge Aue","clubId":"DFL-CLU-00000Y"},{"platzierung":16,"unterPlatzierung":null,"club":"Eintracht Braunschweig","clubId":"DFL-CLU-00000X"},{"platzierung":17,"unterPlatzierung":null,"club":"SpVgg Greuther Fürth","clubId":"DFL-CLU-00000W"},{"platzierung":18,"unterPlatzierung":null,"club":"1. FC Kaiserslautern","clubId":"DFL-CLU-00000I"}]},{"wettbewerbsId":"DFL-COM-000001","wettbewerb":"Bundesliga","spieltagsId":null,"spieltag":null,"eintraege":[{"platzierung":2,"unterPlatzierung":null,"club":"FC Schalke 04","clubId":"DFL-CLU-000009"},{"platzierung":3,"unterPlatzierung":null,"club":"Borussia Dortmund","clubId":"DFL-CLU-000007"},{"platzierung":4,"unterPlatzierung":null,"club":"TSG 1899 Hoffenheim","clubId":"DFL-CLU-000002"},{"platzierung":5,"unterPlatzierung":null,"club":"Bayer 04 Leverkusen","clubId":"DFL-CLU-00000B"},{"platzierung":6,"unterPlatzierung":null,"club":"RB Leipzig","clubId":"DFL-CLU-000017"},{"platzierung":7,"unterPlatzierung":null,"club":"Eintracht Frankfurt","clubId":"DFL-CLU-00000F"},{"platzierung":8,"unterPlatzierung":null,"club":"VfB Stuttgart","clubId":"DFL-CLU-00000D"},{"platzierung":9,"unterPlatzierung":null,"club":"Borussia Mönchengladbach","clubId":"DFL-CLU-000004"},{"platzierung":10,"unterPlatzierung":null,"club":"Hertha BSC","clubId":"DFL-CLU-00000Z"},{"platzierung":11,"unterPlatzierung":null,"club":"FC Augsburg","clubId":"DFL-CLU-000010"},{"platzierung":12,"unterPlatzierung":null,"club":"SV Werder Bremen","clubId":"DFL-CLU-00000E"},{"platzierung":13,"unterPlatzierung":null,"club":"Hannover 96","clubId":"DFL-CLU-000001"},{"platzierung":14,"unterPlatzierung":null,"club":"1. FSV Mainz 05","clubId":"DFL-CLU-000006"},{"platzierung":15,"unterPlatzierung":null,"club":"Sport-Club Freiburg","clubId":"DFL-CLU-00000A"},{"platzierung":16,"unterPlatzierung":null,"club":"VfL Wolfsburg","clubId":"DFL-CLU-000003"},{"platzierung":17,"unterPlatzierung":null,"club":"Hamburger SV","clubId":"DFL-CLU-00000C"},{"platzierung":18,"unterPlatzierung":null,"club":"1. FC Köln","clubId":"DFL-CLU-000008"},{"platzierung":1,"unterPlatzierung":null,"club":"FC Bayern München","clubId":"DFL-CLU-00000G"}]}]`

	var tt []Tables
	err := json.Unmarshal([]byte(src), &tt)
	if err != nil {
		t.Fatalf("error while unmarshaling: %v\n", err)
	}

}
