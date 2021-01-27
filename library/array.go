package library

func In_array(target interface{}, arr []interface{}) bool {
	for _, val := range arr {
		if target == val {
			return true
		}
	}
	return false
}
