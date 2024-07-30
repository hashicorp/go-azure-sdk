package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OutlookTaskImportance string

const (
	OutlookTaskImportance_High   OutlookTaskImportance = "high"
	OutlookTaskImportance_Low    OutlookTaskImportance = "low"
	OutlookTaskImportance_Normal OutlookTaskImportance = "normal"
)

func PossibleValuesForOutlookTaskImportance() []string {
	return []string{
		string(OutlookTaskImportance_High),
		string(OutlookTaskImportance_Low),
		string(OutlookTaskImportance_Normal),
	}
}

func (s *OutlookTaskImportance) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOutlookTaskImportance(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOutlookTaskImportance(input string) (*OutlookTaskImportance, error) {
	vals := map[string]OutlookTaskImportance{
		"high":   OutlookTaskImportance_High,
		"low":    OutlookTaskImportance_Low,
		"normal": OutlookTaskImportance_Normal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OutlookTaskImportance(input)
	return &out, nil
}
