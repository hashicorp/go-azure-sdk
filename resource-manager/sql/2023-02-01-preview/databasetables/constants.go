package databasetables

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TableTemporalType string

const (
	TableTemporalTypeHistoryTable                 TableTemporalType = "HistoryTable"
	TableTemporalTypeNonTemporalTable             TableTemporalType = "NonTemporalTable"
	TableTemporalTypeSystemVersionedTemporalTable TableTemporalType = "SystemVersionedTemporalTable"
)

func PossibleValuesForTableTemporalType() []string {
	return []string{
		string(TableTemporalTypeHistoryTable),
		string(TableTemporalTypeNonTemporalTable),
		string(TableTemporalTypeSystemVersionedTemporalTable),
	}
}

func (s *TableTemporalType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTableTemporalType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTableTemporalType(input string) (*TableTemporalType, error) {
	vals := map[string]TableTemporalType{
		"historytable":                 TableTemporalTypeHistoryTable,
		"nontemporaltable":             TableTemporalTypeNonTemporalTable,
		"systemversionedtemporaltable": TableTemporalTypeSystemVersionedTemporalTable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TableTemporalType(input)
	return &out, nil
}
