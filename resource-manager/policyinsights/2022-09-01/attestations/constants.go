package attestations

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
