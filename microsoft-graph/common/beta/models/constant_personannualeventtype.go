package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PersonAnnualEventType string

const (
	PersonAnnualEventTypebirthday PersonAnnualEventType = "Birthday"
	PersonAnnualEventTypeother    PersonAnnualEventType = "Other"
	PersonAnnualEventTypewedding  PersonAnnualEventType = "Wedding"
	PersonAnnualEventTypework     PersonAnnualEventType = "Work"
)

func PossibleValuesForPersonAnnualEventType() []string {
	return []string{
		string(PersonAnnualEventTypebirthday),
		string(PersonAnnualEventTypeother),
		string(PersonAnnualEventTypewedding),
		string(PersonAnnualEventTypework),
	}
}

func parsePersonAnnualEventType(input string) (*PersonAnnualEventType, error) {
	vals := map[string]PersonAnnualEventType{
		"birthday": PersonAnnualEventTypebirthday,
		"other":    PersonAnnualEventTypeother,
		"wedding":  PersonAnnualEventTypewedding,
		"work":     PersonAnnualEventTypework,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PersonAnnualEventType(input)
	return &out, nil
}
