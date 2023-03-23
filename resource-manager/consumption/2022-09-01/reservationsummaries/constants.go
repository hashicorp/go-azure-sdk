package reservationsummaries

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Datagrain string

const (
	DatagrainDaily   Datagrain = "daily"
	DatagrainMonthly Datagrain = "monthly"
)

func PossibleValuesForDatagrain() []string {
	return []string{
		string(DatagrainDaily),
		string(DatagrainMonthly),
	}
}

func parseDatagrain(input string) (*Datagrain, error) {
	vals := map[string]Datagrain{
		"daily":   DatagrainDaily,
		"monthly": DatagrainMonthly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Datagrain(input)
	return &out, nil
}
