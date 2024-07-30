package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkforceIntegrationEligibilityFilteringEnabledEntities string

const (
	WorkforceIntegrationEligibilityFilteringEnabledEntities_None              WorkforceIntegrationEligibilityFilteringEnabledEntities = "none"
	WorkforceIntegrationEligibilityFilteringEnabledEntities_OfferShiftRequest WorkforceIntegrationEligibilityFilteringEnabledEntities = "offerShiftRequest"
	WorkforceIntegrationEligibilityFilteringEnabledEntities_SwapRequest       WorkforceIntegrationEligibilityFilteringEnabledEntities = "swapRequest"
	WorkforceIntegrationEligibilityFilteringEnabledEntities_TimeOffReason     WorkforceIntegrationEligibilityFilteringEnabledEntities = "timeOffReason"
)

func PossibleValuesForWorkforceIntegrationEligibilityFilteringEnabledEntities() []string {
	return []string{
		string(WorkforceIntegrationEligibilityFilteringEnabledEntities_None),
		string(WorkforceIntegrationEligibilityFilteringEnabledEntities_OfferShiftRequest),
		string(WorkforceIntegrationEligibilityFilteringEnabledEntities_SwapRequest),
		string(WorkforceIntegrationEligibilityFilteringEnabledEntities_TimeOffReason),
	}
}

func (s *WorkforceIntegrationEligibilityFilteringEnabledEntities) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWorkforceIntegrationEligibilityFilteringEnabledEntities(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWorkforceIntegrationEligibilityFilteringEnabledEntities(input string) (*WorkforceIntegrationEligibilityFilteringEnabledEntities, error) {
	vals := map[string]WorkforceIntegrationEligibilityFilteringEnabledEntities{
		"none":              WorkforceIntegrationEligibilityFilteringEnabledEntities_None,
		"offershiftrequest": WorkforceIntegrationEligibilityFilteringEnabledEntities_OfferShiftRequest,
		"swaprequest":       WorkforceIntegrationEligibilityFilteringEnabledEntities_SwapRequest,
		"timeoffreason":     WorkforceIntegrationEligibilityFilteringEnabledEntities_TimeOffReason,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkforceIntegrationEligibilityFilteringEnabledEntities(input)
	return &out, nil
}
