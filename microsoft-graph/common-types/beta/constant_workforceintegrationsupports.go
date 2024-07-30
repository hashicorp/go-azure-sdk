package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkforceIntegrationSupports string

const (
	WorkforceIntegrationSupports_None                 WorkforceIntegrationSupports = "none"
	WorkforceIntegrationSupports_OfferShiftRequest    WorkforceIntegrationSupports = "offerShiftRequest"
	WorkforceIntegrationSupports_OpenShift            WorkforceIntegrationSupports = "openShift"
	WorkforceIntegrationSupports_OpenShiftRequest     WorkforceIntegrationSupports = "openShiftRequest"
	WorkforceIntegrationSupports_Shift                WorkforceIntegrationSupports = "shift"
	WorkforceIntegrationSupports_SwapRequest          WorkforceIntegrationSupports = "swapRequest"
	WorkforceIntegrationSupports_TimeCard             WorkforceIntegrationSupports = "timeCard"
	WorkforceIntegrationSupports_TimeOff              WorkforceIntegrationSupports = "timeOff"
	WorkforceIntegrationSupports_TimeOffReason        WorkforceIntegrationSupports = "timeOffReason"
	WorkforceIntegrationSupports_TimeOffRequest       WorkforceIntegrationSupports = "timeOffRequest"
	WorkforceIntegrationSupports_UserShiftPreferences WorkforceIntegrationSupports = "userShiftPreferences"
)

func PossibleValuesForWorkforceIntegrationSupports() []string {
	return []string{
		string(WorkforceIntegrationSupports_None),
		string(WorkforceIntegrationSupports_OfferShiftRequest),
		string(WorkforceIntegrationSupports_OpenShift),
		string(WorkforceIntegrationSupports_OpenShiftRequest),
		string(WorkforceIntegrationSupports_Shift),
		string(WorkforceIntegrationSupports_SwapRequest),
		string(WorkforceIntegrationSupports_TimeCard),
		string(WorkforceIntegrationSupports_TimeOff),
		string(WorkforceIntegrationSupports_TimeOffReason),
		string(WorkforceIntegrationSupports_TimeOffRequest),
		string(WorkforceIntegrationSupports_UserShiftPreferences),
	}
}

func (s *WorkforceIntegrationSupports) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWorkforceIntegrationSupports(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWorkforceIntegrationSupports(input string) (*WorkforceIntegrationSupports, error) {
	vals := map[string]WorkforceIntegrationSupports{
		"none":                 WorkforceIntegrationSupports_None,
		"offershiftrequest":    WorkforceIntegrationSupports_OfferShiftRequest,
		"openshift":            WorkforceIntegrationSupports_OpenShift,
		"openshiftrequest":     WorkforceIntegrationSupports_OpenShiftRequest,
		"shift":                WorkforceIntegrationSupports_Shift,
		"swaprequest":          WorkforceIntegrationSupports_SwapRequest,
		"timecard":             WorkforceIntegrationSupports_TimeCard,
		"timeoff":              WorkforceIntegrationSupports_TimeOff,
		"timeoffreason":        WorkforceIntegrationSupports_TimeOffReason,
		"timeoffrequest":       WorkforceIntegrationSupports_TimeOffRequest,
		"usershiftpreferences": WorkforceIntegrationSupports_UserShiftPreferences,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkforceIntegrationSupports(input)
	return &out, nil
}
