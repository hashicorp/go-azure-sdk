package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityIncidentClassification string

const (
	SecurityIncidentClassification_FalsePositive                 SecurityIncidentClassification = "falsePositive"
	SecurityIncidentClassification_InformationalExpectedActivity SecurityIncidentClassification = "informationalExpectedActivity"
	SecurityIncidentClassification_TruePositive                  SecurityIncidentClassification = "truePositive"
	SecurityIncidentClassification_Unknown                       SecurityIncidentClassification = "unknown"
)

func PossibleValuesForSecurityIncidentClassification() []string {
	return []string{
		string(SecurityIncidentClassification_FalsePositive),
		string(SecurityIncidentClassification_InformationalExpectedActivity),
		string(SecurityIncidentClassification_TruePositive),
		string(SecurityIncidentClassification_Unknown),
	}
}

func (s *SecurityIncidentClassification) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityIncidentClassification(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityIncidentClassification(input string) (*SecurityIncidentClassification, error) {
	vals := map[string]SecurityIncidentClassification{
		"falsepositive":                 SecurityIncidentClassification_FalsePositive,
		"informationalexpectedactivity": SecurityIncidentClassification_InformationalExpectedActivity,
		"truepositive":                  SecurityIncidentClassification_TruePositive,
		"unknown":                       SecurityIncidentClassification_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityIncidentClassification(input)
	return &out, nil
}
