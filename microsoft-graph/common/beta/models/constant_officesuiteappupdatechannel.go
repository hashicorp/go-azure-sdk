package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OfficeSuiteAppUpdateChannel string

const (
	OfficeSuiteAppUpdateChannelcurrent              OfficeSuiteAppUpdateChannel = "Current"
	OfficeSuiteAppUpdateChanneldeferred             OfficeSuiteAppUpdateChannel = "Deferred"
	OfficeSuiteAppUpdateChannelfirstReleaseCurrent  OfficeSuiteAppUpdateChannel = "FirstReleaseCurrent"
	OfficeSuiteAppUpdateChannelfirstReleaseDeferred OfficeSuiteAppUpdateChannel = "FirstReleaseDeferred"
	OfficeSuiteAppUpdateChannelmonthlyEnterprise    OfficeSuiteAppUpdateChannel = "MonthlyEnterprise"
	OfficeSuiteAppUpdateChannelnone                 OfficeSuiteAppUpdateChannel = "None"
)

func PossibleValuesForOfficeSuiteAppUpdateChannel() []string {
	return []string{
		string(OfficeSuiteAppUpdateChannelcurrent),
		string(OfficeSuiteAppUpdateChanneldeferred),
		string(OfficeSuiteAppUpdateChannelfirstReleaseCurrent),
		string(OfficeSuiteAppUpdateChannelfirstReleaseDeferred),
		string(OfficeSuiteAppUpdateChannelmonthlyEnterprise),
		string(OfficeSuiteAppUpdateChannelnone),
	}
}

func parseOfficeSuiteAppUpdateChannel(input string) (*OfficeSuiteAppUpdateChannel, error) {
	vals := map[string]OfficeSuiteAppUpdateChannel{
		"current":              OfficeSuiteAppUpdateChannelcurrent,
		"deferred":             OfficeSuiteAppUpdateChanneldeferred,
		"firstreleasecurrent":  OfficeSuiteAppUpdateChannelfirstReleaseCurrent,
		"firstreleasedeferred": OfficeSuiteAppUpdateChannelfirstReleaseDeferred,
		"monthlyenterprise":    OfficeSuiteAppUpdateChannelmonthlyEnterprise,
		"none":                 OfficeSuiteAppUpdateChannelnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OfficeSuiteAppUpdateChannel(input)
	return &out, nil
}
