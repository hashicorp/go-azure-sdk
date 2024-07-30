package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestrictedAppsViolationRestrictedAppsState string

const (
	RestrictedAppsViolationRestrictedAppsState_NotApprovedApps RestrictedAppsViolationRestrictedAppsState = "notApprovedApps"
	RestrictedAppsViolationRestrictedAppsState_ProhibitedApps  RestrictedAppsViolationRestrictedAppsState = "prohibitedApps"
)

func PossibleValuesForRestrictedAppsViolationRestrictedAppsState() []string {
	return []string{
		string(RestrictedAppsViolationRestrictedAppsState_NotApprovedApps),
		string(RestrictedAppsViolationRestrictedAppsState_ProhibitedApps),
	}
}

func (s *RestrictedAppsViolationRestrictedAppsState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRestrictedAppsViolationRestrictedAppsState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRestrictedAppsViolationRestrictedAppsState(input string) (*RestrictedAppsViolationRestrictedAppsState, error) {
	vals := map[string]RestrictedAppsViolationRestrictedAppsState{
		"notapprovedapps": RestrictedAppsViolationRestrictedAppsState_NotApprovedApps,
		"prohibitedapps":  RestrictedAppsViolationRestrictedAppsState_ProhibitedApps,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RestrictedAppsViolationRestrictedAppsState(input)
	return &out, nil
}
