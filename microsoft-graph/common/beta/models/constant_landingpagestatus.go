package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LandingPageStatus string

const (
	LandingPageStatusarchive LandingPageStatus = "Archive"
	LandingPageStatusdelete  LandingPageStatus = "Delete"
	LandingPageStatusdraft   LandingPageStatus = "Draft"
	LandingPageStatusready   LandingPageStatus = "Ready"
	LandingPageStatusunknown LandingPageStatus = "Unknown"
)

func PossibleValuesForLandingPageStatus() []string {
	return []string{
		string(LandingPageStatusarchive),
		string(LandingPageStatusdelete),
		string(LandingPageStatusdraft),
		string(LandingPageStatusready),
		string(LandingPageStatusunknown),
	}
}

func parseLandingPageStatus(input string) (*LandingPageStatus, error) {
	vals := map[string]LandingPageStatus{
		"archive": LandingPageStatusarchive,
		"delete":  LandingPageStatusdelete,
		"draft":   LandingPageStatusdraft,
		"ready":   LandingPageStatusready,
		"unknown": LandingPageStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LandingPageStatus(input)
	return &out, nil
}
