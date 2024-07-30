package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComanagementEligibleDeviceStatus string

const (
	ComanagementEligibleDeviceStatus_Comanaged                   ComanagementEligibleDeviceStatus = "comanaged"
	ComanagementEligibleDeviceStatus_Eligible                    ComanagementEligibleDeviceStatus = "eligible"
	ComanagementEligibleDeviceStatus_EligibleButNotAzureAdJoined ComanagementEligibleDeviceStatus = "eligibleButNotAzureAdJoined"
	ComanagementEligibleDeviceStatus_Ineligible                  ComanagementEligibleDeviceStatus = "ineligible"
	ComanagementEligibleDeviceStatus_NeedsOsUpdate               ComanagementEligibleDeviceStatus = "needsOsUpdate"
	ComanagementEligibleDeviceStatus_ScheduledForEnrollment      ComanagementEligibleDeviceStatus = "scheduledForEnrollment"
)

func PossibleValuesForComanagementEligibleDeviceStatus() []string {
	return []string{
		string(ComanagementEligibleDeviceStatus_Comanaged),
		string(ComanagementEligibleDeviceStatus_Eligible),
		string(ComanagementEligibleDeviceStatus_EligibleButNotAzureAdJoined),
		string(ComanagementEligibleDeviceStatus_Ineligible),
		string(ComanagementEligibleDeviceStatus_NeedsOsUpdate),
		string(ComanagementEligibleDeviceStatus_ScheduledForEnrollment),
	}
}

func (s *ComanagementEligibleDeviceStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseComanagementEligibleDeviceStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseComanagementEligibleDeviceStatus(input string) (*ComanagementEligibleDeviceStatus, error) {
	vals := map[string]ComanagementEligibleDeviceStatus{
		"comanaged":                   ComanagementEligibleDeviceStatus_Comanaged,
		"eligible":                    ComanagementEligibleDeviceStatus_Eligible,
		"eligiblebutnotazureadjoined": ComanagementEligibleDeviceStatus_EligibleButNotAzureAdJoined,
		"ineligible":                  ComanagementEligibleDeviceStatus_Ineligible,
		"needsosupdate":               ComanagementEligibleDeviceStatus_NeedsOsUpdate,
		"scheduledforenrollment":      ComanagementEligibleDeviceStatus_ScheduledForEnrollment,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ComanagementEligibleDeviceStatus(input)
	return &out, nil
}
