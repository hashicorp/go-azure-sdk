package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkSoftwareUpdateStatusSoftwareFreshness string

const (
	TeamworkSoftwareUpdateStatusSoftwareFreshnesslatest          TeamworkSoftwareUpdateStatusSoftwareFreshness = "Latest"
	TeamworkSoftwareUpdateStatusSoftwareFreshnessunknown         TeamworkSoftwareUpdateStatusSoftwareFreshness = "Unknown"
	TeamworkSoftwareUpdateStatusSoftwareFreshnessupdateAvailable TeamworkSoftwareUpdateStatusSoftwareFreshness = "UpdateAvailable"
)

func PossibleValuesForTeamworkSoftwareUpdateStatusSoftwareFreshness() []string {
	return []string{
		string(TeamworkSoftwareUpdateStatusSoftwareFreshnesslatest),
		string(TeamworkSoftwareUpdateStatusSoftwareFreshnessunknown),
		string(TeamworkSoftwareUpdateStatusSoftwareFreshnessupdateAvailable),
	}
}

func parseTeamworkSoftwareUpdateStatusSoftwareFreshness(input string) (*TeamworkSoftwareUpdateStatusSoftwareFreshness, error) {
	vals := map[string]TeamworkSoftwareUpdateStatusSoftwareFreshness{
		"latest":          TeamworkSoftwareUpdateStatusSoftwareFreshnesslatest,
		"unknown":         TeamworkSoftwareUpdateStatusSoftwareFreshnessunknown,
		"updateavailable": TeamworkSoftwareUpdateStatusSoftwareFreshnessupdateAvailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamworkSoftwareUpdateStatusSoftwareFreshness(input)
	return &out, nil
}
