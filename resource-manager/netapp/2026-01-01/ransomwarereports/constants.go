package ransomwarereports

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RansomwareReportSeverity string

const (
	RansomwareReportSeverityHigh     RansomwareReportSeverity = "High"
	RansomwareReportSeverityLow      RansomwareReportSeverity = "Low"
	RansomwareReportSeverityModerate RansomwareReportSeverity = "Moderate"
	RansomwareReportSeverityNone     RansomwareReportSeverity = "None"
)

func PossibleValuesForRansomwareReportSeverity() []string {
	return []string{
		string(RansomwareReportSeverityHigh),
		string(RansomwareReportSeverityLow),
		string(RansomwareReportSeverityModerate),
		string(RansomwareReportSeverityNone),
	}
}

func (s *RansomwareReportSeverity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRansomwareReportSeverity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRansomwareReportSeverity(input string) (*RansomwareReportSeverity, error) {
	vals := map[string]RansomwareReportSeverity{
		"high":     RansomwareReportSeverityHigh,
		"low":      RansomwareReportSeverityLow,
		"moderate": RansomwareReportSeverityModerate,
		"none":     RansomwareReportSeverityNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RansomwareReportSeverity(input)
	return &out, nil
}

type RansomwareReportState string

const (
	RansomwareReportStateActive   RansomwareReportState = "Active"
	RansomwareReportStateResolved RansomwareReportState = "Resolved"
)

func PossibleValuesForRansomwareReportState() []string {
	return []string{
		string(RansomwareReportStateActive),
		string(RansomwareReportStateResolved),
	}
}

func (s *RansomwareReportState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRansomwareReportState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRansomwareReportState(input string) (*RansomwareReportState, error) {
	vals := map[string]RansomwareReportState{
		"active":   RansomwareReportStateActive,
		"resolved": RansomwareReportStateResolved,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RansomwareReportState(input)
	return &out, nil
}

type RansomwareSuspectResolution string

const (
	RansomwareSuspectResolutionFalsePositive   RansomwareSuspectResolution = "FalsePositive"
	RansomwareSuspectResolutionPotentialThreat RansomwareSuspectResolution = "PotentialThreat"
)

func PossibleValuesForRansomwareSuspectResolution() []string {
	return []string{
		string(RansomwareSuspectResolutionFalsePositive),
		string(RansomwareSuspectResolutionPotentialThreat),
	}
}

func (s *RansomwareSuspectResolution) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRansomwareSuspectResolution(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRansomwareSuspectResolution(input string) (*RansomwareSuspectResolution, error) {
	vals := map[string]RansomwareSuspectResolution{
		"falsepositive":   RansomwareSuspectResolutionFalsePositive,
		"potentialthreat": RansomwareSuspectResolutionPotentialThreat,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RansomwareSuspectResolution(input)
	return &out, nil
}
