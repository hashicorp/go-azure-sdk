package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkforceIntegrationSupportedEntities string

const (
	WorkforceIntegrationSupportedEntitiesnone                 WorkforceIntegrationSupportedEntities = "None"
	WorkforceIntegrationSupportedEntitiesofferShiftRequest    WorkforceIntegrationSupportedEntities = "OfferShiftRequest"
	WorkforceIntegrationSupportedEntitiesopenShift            WorkforceIntegrationSupportedEntities = "OpenShift"
	WorkforceIntegrationSupportedEntitiesopenShiftRequest     WorkforceIntegrationSupportedEntities = "OpenShiftRequest"
	WorkforceIntegrationSupportedEntitiesshift                WorkforceIntegrationSupportedEntities = "Shift"
	WorkforceIntegrationSupportedEntitiesswapRequest          WorkforceIntegrationSupportedEntities = "SwapRequest"
	WorkforceIntegrationSupportedEntitiestimeCard             WorkforceIntegrationSupportedEntities = "TimeCard"
	WorkforceIntegrationSupportedEntitiestimeOff              WorkforceIntegrationSupportedEntities = "TimeOff"
	WorkforceIntegrationSupportedEntitiestimeOffReason        WorkforceIntegrationSupportedEntities = "TimeOffReason"
	WorkforceIntegrationSupportedEntitiestimeOffRequest       WorkforceIntegrationSupportedEntities = "TimeOffRequest"
	WorkforceIntegrationSupportedEntitiesuserShiftPreferences WorkforceIntegrationSupportedEntities = "UserShiftPreferences"
)

func PossibleValuesForWorkforceIntegrationSupportedEntities() []string {
	return []string{
		string(WorkforceIntegrationSupportedEntitiesnone),
		string(WorkforceIntegrationSupportedEntitiesofferShiftRequest),
		string(WorkforceIntegrationSupportedEntitiesopenShift),
		string(WorkforceIntegrationSupportedEntitiesopenShiftRequest),
		string(WorkforceIntegrationSupportedEntitiesshift),
		string(WorkforceIntegrationSupportedEntitiesswapRequest),
		string(WorkforceIntegrationSupportedEntitiestimeCard),
		string(WorkforceIntegrationSupportedEntitiestimeOff),
		string(WorkforceIntegrationSupportedEntitiestimeOffReason),
		string(WorkforceIntegrationSupportedEntitiestimeOffRequest),
		string(WorkforceIntegrationSupportedEntitiesuserShiftPreferences),
	}
}

func parseWorkforceIntegrationSupportedEntities(input string) (*WorkforceIntegrationSupportedEntities, error) {
	vals := map[string]WorkforceIntegrationSupportedEntities{
		"none":                 WorkforceIntegrationSupportedEntitiesnone,
		"offershiftrequest":    WorkforceIntegrationSupportedEntitiesofferShiftRequest,
		"openshift":            WorkforceIntegrationSupportedEntitiesopenShift,
		"openshiftrequest":     WorkforceIntegrationSupportedEntitiesopenShiftRequest,
		"shift":                WorkforceIntegrationSupportedEntitiesshift,
		"swaprequest":          WorkforceIntegrationSupportedEntitiesswapRequest,
		"timecard":             WorkforceIntegrationSupportedEntitiestimeCard,
		"timeoff":              WorkforceIntegrationSupportedEntitiestimeOff,
		"timeoffreason":        WorkforceIntegrationSupportedEntitiestimeOffReason,
		"timeoffrequest":       WorkforceIntegrationSupportedEntitiestimeOffRequest,
		"usershiftpreferences": WorkforceIntegrationSupportedEntitiesuserShiftPreferences,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkforceIntegrationSupportedEntities(input)
	return &out, nil
}
