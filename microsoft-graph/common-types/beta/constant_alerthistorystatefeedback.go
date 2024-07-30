package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertHistoryStateFeedback string

const (
	AlertHistoryStateFeedback_BenignPositive AlertHistoryStateFeedback = "benignPositive"
	AlertHistoryStateFeedback_FalsePositive  AlertHistoryStateFeedback = "falsePositive"
	AlertHistoryStateFeedback_TruePositive   AlertHistoryStateFeedback = "truePositive"
	AlertHistoryStateFeedback_Unknown        AlertHistoryStateFeedback = "unknown"
)

func PossibleValuesForAlertHistoryStateFeedback() []string {
	return []string{
		string(AlertHistoryStateFeedback_BenignPositive),
		string(AlertHistoryStateFeedback_FalsePositive),
		string(AlertHistoryStateFeedback_TruePositive),
		string(AlertHistoryStateFeedback_Unknown),
	}
}

func (s *AlertHistoryStateFeedback) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAlertHistoryStateFeedback(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAlertHistoryStateFeedback(input string) (*AlertHistoryStateFeedback, error) {
	vals := map[string]AlertHistoryStateFeedback{
		"benignpositive": AlertHistoryStateFeedback_BenignPositive,
		"falsepositive":  AlertHistoryStateFeedback_FalsePositive,
		"truepositive":   AlertHistoryStateFeedback_TruePositive,
		"unknown":        AlertHistoryStateFeedback_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AlertHistoryStateFeedback(input)
	return &out, nil
}
