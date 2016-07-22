package data_structures

func findPeopleWithCommonInterest(data map[string][]string, interest string) []string {
	res := make([]string, 0)
	// v, ok := data["adsf"]
	for k,v := range data { // for each key,val
		if contains(v, interest) ==true{ // if v contains interest
			res = append(res, k)
		}
	}
	return res
}

func contains(src []string, elem string) bool {
	for s:= range src{
		if src[s] == elem{
			return true
		}
	}
	return false
}
