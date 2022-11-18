package attestations

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComplianceState string

const (
	ComplianceStateCompliant    ComplianceState = "Compliant"
	ComplianceStateNonCompliant ComplianceState = "NonCompliant"
	ComplianceStateUnknown      ComplianceState = "Unknown"
)

func PossibleValuesForComplianceState() []string {
	return []string{
		string(ComplianceStateCompliant),
		string(ComplianceStateNonCompliant),
		string(ComplianceStateUnknown),
	}
}

func parseComplianceState(input string) (*ComplianceState, error) {
	vals := map[string]ComplianceState{
		"compliant":    ComplianceStateCompliant,
		"noncompliant": ComplianceStateNonCompliant,
		"unknown":      ComplianceStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ComplianceState(input)
	return &out, nil
}
