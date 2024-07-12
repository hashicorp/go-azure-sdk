package guestconfigurationassignmenthcrpreports

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComplianceStatus string

const (
	ComplianceStatusCompliant    ComplianceStatus = "Compliant"
	ComplianceStatusNonCompliant ComplianceStatus = "NonCompliant"
	ComplianceStatusPending      ComplianceStatus = "Pending"
)

func PossibleValuesForComplianceStatus() []string {
	return []string{
		string(ComplianceStatusCompliant),
		string(ComplianceStatusNonCompliant),
		string(ComplianceStatusPending),
	}
}

func (s *ComplianceStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseComplianceStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseComplianceStatus(input string) (*ComplianceStatus, error) {
	vals := map[string]ComplianceStatus{
		"compliant":    ComplianceStatusCompliant,
		"noncompliant": ComplianceStatusNonCompliant,
		"pending":      ComplianceStatusPending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ComplianceStatus(input)
	return &out, nil
}

type Type string

const (
	TypeConsistency Type = "Consistency"
	TypeInitial     Type = "Initial"
)

func PossibleValuesForType() []string {
	return []string{
		string(TypeConsistency),
		string(TypeInitial),
	}
}

func (s *Type) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseType(input string) (*Type, error) {
	vals := map[string]Type{
		"consistency": TypeConsistency,
		"initial":     TypeInitial,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Type(input)
	return &out, nil
}
