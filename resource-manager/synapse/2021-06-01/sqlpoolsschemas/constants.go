package sqlpoolsschemas

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VulnerabilityAssessmentScanState string

const (
	VulnerabilityAssessmentScanStateFailed      VulnerabilityAssessmentScanState = "Failed"
	VulnerabilityAssessmentScanStateFailedToRun VulnerabilityAssessmentScanState = "FailedToRun"
	VulnerabilityAssessmentScanStateInProgress  VulnerabilityAssessmentScanState = "InProgress"
	VulnerabilityAssessmentScanStatePassed      VulnerabilityAssessmentScanState = "Passed"
)

func PossibleValuesForVulnerabilityAssessmentScanState() []string {
	return []string{
		string(VulnerabilityAssessmentScanStateFailed),
		string(VulnerabilityAssessmentScanStateFailedToRun),
		string(VulnerabilityAssessmentScanStateInProgress),
		string(VulnerabilityAssessmentScanStatePassed),
	}
}

func parseVulnerabilityAssessmentScanState(input string) (*VulnerabilityAssessmentScanState, error) {
	vals := map[string]VulnerabilityAssessmentScanState{
		"failed":      VulnerabilityAssessmentScanStateFailed,
		"failedtorun": VulnerabilityAssessmentScanStateFailedToRun,
		"inprogress":  VulnerabilityAssessmentScanStateInProgress,
		"passed":      VulnerabilityAssessmentScanStatePassed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VulnerabilityAssessmentScanState(input)
	return &out, nil
}

type VulnerabilityAssessmentScanTriggerType string

const (
	VulnerabilityAssessmentScanTriggerTypeOnDemand  VulnerabilityAssessmentScanTriggerType = "OnDemand"
	VulnerabilityAssessmentScanTriggerTypeRecurring VulnerabilityAssessmentScanTriggerType = "Recurring"
)

func PossibleValuesForVulnerabilityAssessmentScanTriggerType() []string {
	return []string{
		string(VulnerabilityAssessmentScanTriggerTypeOnDemand),
		string(VulnerabilityAssessmentScanTriggerTypeRecurring),
	}
}

func parseVulnerabilityAssessmentScanTriggerType(input string) (*VulnerabilityAssessmentScanTriggerType, error) {
	vals := map[string]VulnerabilityAssessmentScanTriggerType{
		"ondemand":  VulnerabilityAssessmentScanTriggerTypeOnDemand,
		"recurring": VulnerabilityAssessmentScanTriggerTypeRecurring,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VulnerabilityAssessmentScanTriggerType(input)
	return &out, nil
}
