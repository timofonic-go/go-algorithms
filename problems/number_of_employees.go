package problems

// NumberOfReports implements solution for a problem "Find number of Employees Under every Employee".
// Based on http://www.geeksforgeeks.org/find-number-of-employees-under-every-manager/
func NumberOfReports(dataset map[string]string) map[string]int {

	// Resulting map of manager -> direct reports quantity,
	// i.e. "A" -> 0, "B" -> 4, etc.
	result := make(map[string]int)

	// 1. Prepare a map of Manager -> Employee relations
	// To store reverse of original map, each key will have 0 to multiple values
	mngrEmpMap := make(map[string][]string)

	// To fill mngrEmpMap, iterate through the given map
	for emp, mngr := range dataset {

		// skip emp-emp entry
		if emp == mngr {
			continue
		}

		// Get the previous list of direct reports under
		// current 'mngr' and add the current 'emp' to the list
		directReportList, ok := mngrEmpMap[mngr]

		// If 'emp' is the first employee under 'mngr'
		if !ok {
			directReportList = []string{}
		}

		directReportList = append(directReportList, emp)

		// Replace old value for 'mngr' with new directReportList
		mngrEmpMap[mngr] = directReportList
	}

	// 2. Use manager->employee map mngrEmpMap built above to populate result using getManagerReports()
	for mngr := range mngrEmpMap {
		_ = getManagerReports(&result, mngr, &mngrEmpMap)
	}

	return result
}

// getManagerReports computes a total number of reports for a given manager.
func getManagerReports(result *map[string]int, mngr string, mngrEmpMap *map[string][]string) int {

	// Check if a given manager ("mngr" var) has any reports.
	// Need this to understand if need to calculate reports for this person.
	directReports, ok := (*mngrEmpMap)[mngr]
	if !ok {
		// If not, create a zero record for this person and return.
		(*result)[mngr] = 0
		return 0
	}

	// If yes, let's count his/her direct reports.
	// Staying within the current function until we get the total number.

	// Check if already counted.
	cnt, ok := (*result)[mngr]

	// We set only a final result (total count) into the result map. Thus, if a person is found in the result map,
	// then we can just return the value ("cnt" var).
	if ok {
		return cnt
	}

	// If not found, this means he/she were not processed yet.

	// We start accumulating a total result with a number of direct reports of the 1st level,
	// ex. "C" and "E" for manager "F" ("F" self-report is skipped above).
	(*result)[mngr] = len(directReports)

	// Additionally, increment the count started above by a total of lower level reports.
	// Ex. "A" and "B" for report "C" (ex. above), and "D" for report "E" (ex. above).
	for _, emp := range directReports {
		(*result)[mngr] += getManagerReports(result, emp, mngrEmpMap)
	}

	// We get to this return only when we have a final, total count of reports for a given manager ("mngr" var).
	return (*result)[mngr]
}
