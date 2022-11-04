package usages

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UsageUnit string

const (
	UsageUnitCount UsageUnit = "Count"
)

func PossibleValuesForUsageUnit() []string {
	return []string{
		string(UsageUnitCount),
	}
}

func parseUsageUnit(input string) (*UsageUnit, error) {
	vals := map[string]UsageUnit{
		"count": UsageUnitCount,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UsageUnit(input)
	return &out, nil
}
