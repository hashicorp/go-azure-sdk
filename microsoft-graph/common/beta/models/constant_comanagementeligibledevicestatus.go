package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComanagementEligibleDeviceStatus string

const (
	ComanagementEligibleDeviceStatuscomanaged                   ComanagementEligibleDeviceStatus = "Comanaged"
	ComanagementEligibleDeviceStatuseligible                    ComanagementEligibleDeviceStatus = "Eligible"
	ComanagementEligibleDeviceStatuseligibleButNotAzureAdJoined ComanagementEligibleDeviceStatus = "EligibleButNotAzureAdJoined"
	ComanagementEligibleDeviceStatusineligible                  ComanagementEligibleDeviceStatus = "Ineligible"
	ComanagementEligibleDeviceStatusneedsOsUpdate               ComanagementEligibleDeviceStatus = "NeedsOsUpdate"
	ComanagementEligibleDeviceStatusscheduledForEnrollment      ComanagementEligibleDeviceStatus = "ScheduledForEnrollment"
)

func PossibleValuesForComanagementEligibleDeviceStatus() []string {
	return []string{
		string(ComanagementEligibleDeviceStatuscomanaged),
		string(ComanagementEligibleDeviceStatuseligible),
		string(ComanagementEligibleDeviceStatuseligibleButNotAzureAdJoined),
		string(ComanagementEligibleDeviceStatusineligible),
		string(ComanagementEligibleDeviceStatusneedsOsUpdate),
		string(ComanagementEligibleDeviceStatusscheduledForEnrollment),
	}
}

func parseComanagementEligibleDeviceStatus(input string) (*ComanagementEligibleDeviceStatus, error) {
	vals := map[string]ComanagementEligibleDeviceStatus{
		"comanaged":                   ComanagementEligibleDeviceStatuscomanaged,
		"eligible":                    ComanagementEligibleDeviceStatuseligible,
		"eligiblebutnotazureadjoined": ComanagementEligibleDeviceStatuseligibleButNotAzureAdJoined,
		"ineligible":                  ComanagementEligibleDeviceStatusineligible,
		"needsosupdate":               ComanagementEligibleDeviceStatusneedsOsUpdate,
		"scheduledforenrollment":      ComanagementEligibleDeviceStatusscheduledForEnrollment,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ComanagementEligibleDeviceStatus(input)
	return &out, nil
}
