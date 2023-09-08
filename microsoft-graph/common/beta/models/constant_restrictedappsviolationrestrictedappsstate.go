package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestrictedAppsViolationRestrictedAppsState string

const (
	RestrictedAppsViolationRestrictedAppsStatenotApprovedApps RestrictedAppsViolationRestrictedAppsState = "NotApprovedApps"
	RestrictedAppsViolationRestrictedAppsStateprohibitedApps  RestrictedAppsViolationRestrictedAppsState = "ProhibitedApps"
)

func PossibleValuesForRestrictedAppsViolationRestrictedAppsState() []string {
	return []string{
		string(RestrictedAppsViolationRestrictedAppsStatenotApprovedApps),
		string(RestrictedAppsViolationRestrictedAppsStateprohibitedApps),
	}
}

func parseRestrictedAppsViolationRestrictedAppsState(input string) (*RestrictedAppsViolationRestrictedAppsState, error) {
	vals := map[string]RestrictedAppsViolationRestrictedAppsState{
		"notapprovedapps": RestrictedAppsViolationRestrictedAppsStatenotApprovedApps,
		"prohibitedapps":  RestrictedAppsViolationRestrictedAppsStateprohibitedApps,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RestrictedAppsViolationRestrictedAppsState(input)
	return &out, nil
}
