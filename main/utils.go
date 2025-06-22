package main

func RemoveItemsFromList(original []string, removedList []string) []string {
	var finalList []string
	if len(original) == 0 {
		return original
	}

	if len(removedList) == 0 {
		return original
	}

	removedListMap := make(map[string]string)
	for _, value := range removedList {
		removedListMap[value] = value
	}

	for _, value := range original {
		if removedListMap[value] != value {
			finalList = append(finalList, value)
		}
	}

	if len(finalList) > 0 {
		return finalList
	}

	return original
}

func Contains(list []string, value string) bool {
	if len(list) == 0 {
		return false
	}
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

func CheckDuplicacyInList(original []string) bool {
	keys := make(map[string]bool)
	flag := false
	for _, entry := range original {
		if _, value := keys[entry]; value {
			flag = true
			break
		} else {
			keys[entry] = true
		}

	}
	return flag
}

func RemoveDuplicateFromList(original []string) []string {
	keys := make(map[string]bool)
	finalList := []string{}

	for _, entry := range original {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			finalList = append(finalList, entry)
		}
	}

	return finalList
}
