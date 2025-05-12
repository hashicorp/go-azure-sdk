package vaultusages

import (
	"encoding/json"
	"fmt"
	"strings"
)

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

func (s *UsagesUnit) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUsagesUnit(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
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
