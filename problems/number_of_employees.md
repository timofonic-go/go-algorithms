# Find number of Employees Under every Employee

Given a dictionary that contains mapping of employee and his manager as a number of (employee, manager) pairs
like below:

```Go
{ "A", "C" },
{ "B", "C" },
{ "C", "F" },
{ "D", "E" },
{ "E", "F" },
{ "F", "F" }
```

In this example C is manager of A,
C is also manager of B, F is manager of C
and so on.

Write a function to get number of employees under each manager in the hierarchy; not just their direct reports.
It may be assumed that an employee directly reports to only one manager.
In the above dictionary the root node/CEO is listed as reporting to himself.


## ALGORITHM

### 1. Create a map of managers to their direct employees.
This will serve as:
- an entry point to drill into hierarchy of employees and accumulate employee count
- lookup table to know if a given employee has reports


We will get a map[string][]string "mngrEmpMap":

```Go
"C": []string{"A", "B"}
"E": []string{"D"}
"F": []string{"C", "E", "F"}
```

### 2. Iterate through a given map of employees to count their reports with the help of the map created on step #1.

In a loop:

2.1 If employee is not present in "mngrEmpMap":

	create a record with this employee as a key and its count 0.
	
2.2 If employee is present in "mngrEmpMap":
	
	2.2.1 Loop through a list of direct reports (i.e. {"C", "E", "F"} for a emp/key "F")
	
	2.2.2 Run recursive calculation of "emp"'s direct employees to get a total number for "emp".
	
	2.2.3 Store total count in a result map with a key being "emp".
