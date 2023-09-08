package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkforceIntegrationSupports string

const (
	WorkforceIntegrationSupportsnone                 WorkforceIntegrationSupports = "None"
	WorkforceIntegrationSupportsofferShiftRequest    WorkforceIntegrationSupports = "OfferShiftRequest"
	WorkforceIntegrationSupportsopenShift            WorkforceIntegrationSupports = "OpenShift"
	WorkforceIntegrationSupportsopenShiftRequest     WorkforceIntegrationSupports = "OpenShiftRequest"
	WorkforceIntegrationSupportsshift                WorkforceIntegrationSupports = "Shift"
	WorkforceIntegrationSupportsswapRequest          WorkforceIntegrationSupports = "SwapRequest"
	WorkforceIntegrationSupportstimeCard             WorkforceIntegrationSupports = "TimeCard"
	WorkforceIntegrationSupportstimeOff              WorkforceIntegrationSupports = "TimeOff"
	WorkforceIntegrationSupportstimeOffReason        WorkforceIntegrationSupports = "TimeOffReason"
	WorkforceIntegrationSupportstimeOffRequest       WorkforceIntegrationSupports = "TimeOffRequest"
	WorkforceIntegrationSupportsuserShiftPreferences WorkforceIntegrationSupports = "UserShiftPreferences"
)

func PossibleValuesForWorkforceIntegrationSupports() []string {
	return []string{
		string(WorkforceIntegrationSupportsnone),
		string(WorkforceIntegrationSupportsofferShiftRequest),
		string(WorkforceIntegrationSupportsopenShift),
		string(WorkforceIntegrationSupportsopenShiftRequest),
		string(WorkforceIntegrationSupportsshift),
		string(WorkforceIntegrationSupportsswapRequest),
		string(WorkforceIntegrationSupportstimeCard),
		string(WorkforceIntegrationSupportstimeOff),
		string(WorkforceIntegrationSupportstimeOffReason),
		string(WorkforceIntegrationSupportstimeOffRequest),
		string(WorkforceIntegrationSupportsuserShiftPreferences),
	}
}

func parseWorkforceIntegrationSupports(input string) (*WorkforceIntegrationSupports, error) {
	vals := map[string]WorkforceIntegrationSupports{
		"none":                 WorkforceIntegrationSupportsnone,
		"offershiftrequest":    WorkforceIntegrationSupportsofferShiftRequest,
		"openshift":            WorkforceIntegrationSupportsopenShift,
		"openshiftrequest":     WorkforceIntegrationSupportsopenShiftRequest,
		"shift":                WorkforceIntegrationSupportsshift,
		"swaprequest":          WorkforceIntegrationSupportsswapRequest,
		"timecard":             WorkforceIntegrationSupportstimeCard,
		"timeoff":              WorkforceIntegrationSupportstimeOff,
		"timeoffreason":        WorkforceIntegrationSupportstimeOffReason,
		"timeoffrequest":       WorkforceIntegrationSupportstimeOffRequest,
		"usershiftpreferences": WorkforceIntegrationSupportsuserShiftPreferences,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkforceIntegrationSupports(input)
	return &out, nil
}
