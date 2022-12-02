package backupusagesummaries

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UsagesUnit string

const (
	UsagesUnitBytes          UsagesUnit = "Bytes"
	UsagesUnitBytesPerSecond UsagesUnit = "BytesPerSecond"
	UsagesUnitCount          UsagesUnit = "Count"
	UsagesUnitCountPerSecond UsagesUnit = "CountPerSecond"
	UsagesUnitPercent        UsagesUnit = "Percent"
	UsagesUnitSeconds        UsagesUnit = "Seconds"
)

func PossibleValuesForUsagesUnit() []string {
	return []string{
		string(UsagesUnitBytes),
		string(UsagesUnitBytesPerSecond),
		string(UsagesUnitCount),
		string(UsagesUnitCountPerSecond),
		string(UsagesUnitPercent),
		string(UsagesUnitSeconds),
	}
}

func parseUsagesUnit(input string) (*UsagesUnit, error) {
	vals := map[string]UsagesUnit{
		"bytes":          UsagesUnitBytes,
		"bytespersecond": UsagesUnitBytesPerSecond,
		"count":          UsagesUnitCount,
		"countpersecond": UsagesUnitCountPerSecond,
		"percent":        UsagesUnitPercent,
		"seconds":        UsagesUnitSeconds,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UsagesUnit(input)
	return &out, nil
}
