package testlines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCanceled),
		string(ProvisioningStateFailed),
		string(ProvisioningStateSucceeded),
	}
}

type TestLinePurpose string

const (
	TestLinePurposeAutomated TestLinePurpose = "Automated"
	TestLinePurposeManual    TestLinePurpose = "Manual"
)

func PossibleValuesForTestLinePurpose() []string {
	return []string{
		string(TestLinePurposeAutomated),
		string(TestLinePurposeManual),
	}
}
