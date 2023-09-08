package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkforceIntegrationEligibilityFilteringEnabledEntities string

const (
	WorkforceIntegrationEligibilityFilteringEnabledEntitiesnone              WorkforceIntegrationEligibilityFilteringEnabledEntities = "None"
	WorkforceIntegrationEligibilityFilteringEnabledEntitiesofferShiftRequest WorkforceIntegrationEligibilityFilteringEnabledEntities = "OfferShiftRequest"
	WorkforceIntegrationEligibilityFilteringEnabledEntitiesswapRequest       WorkforceIntegrationEligibilityFilteringEnabledEntities = "SwapRequest"
	WorkforceIntegrationEligibilityFilteringEnabledEntitiestimeOffReason     WorkforceIntegrationEligibilityFilteringEnabledEntities = "TimeOffReason"
)

func PossibleValuesForWorkforceIntegrationEligibilityFilteringEnabledEntities() []string {
	return []string{
		string(WorkforceIntegrationEligibilityFilteringEnabledEntitiesnone),
		string(WorkforceIntegrationEligibilityFilteringEnabledEntitiesofferShiftRequest),
		string(WorkforceIntegrationEligibilityFilteringEnabledEntitiesswapRequest),
		string(WorkforceIntegrationEligibilityFilteringEnabledEntitiestimeOffReason),
	}
}

func parseWorkforceIntegrationEligibilityFilteringEnabledEntities(input string) (*WorkforceIntegrationEligibilityFilteringEnabledEntities, error) {
	vals := map[string]WorkforceIntegrationEligibilityFilteringEnabledEntities{
		"none":              WorkforceIntegrationEligibilityFilteringEnabledEntitiesnone,
		"offershiftrequest": WorkforceIntegrationEligibilityFilteringEnabledEntitiesofferShiftRequest,
		"swaprequest":       WorkforceIntegrationEligibilityFilteringEnabledEntitiesswapRequest,
		"timeoffreason":     WorkforceIntegrationEligibilityFilteringEnabledEntitiestimeOffReason,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkforceIntegrationEligibilityFilteringEnabledEntities(input)
	return &out, nil
}
