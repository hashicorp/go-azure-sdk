package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CredentialUsageSummaryFeature string

const (
	CredentialUsageSummaryFeatureregistration CredentialUsageSummaryFeature = "Registration"
	CredentialUsageSummaryFeaturereset        CredentialUsageSummaryFeature = "Reset"
)

func PossibleValuesForCredentialUsageSummaryFeature() []string {
	return []string{
		string(CredentialUsageSummaryFeatureregistration),
		string(CredentialUsageSummaryFeaturereset),
	}
}

func parseCredentialUsageSummaryFeature(input string) (*CredentialUsageSummaryFeature, error) {
	vals := map[string]CredentialUsageSummaryFeature{
		"registration": CredentialUsageSummaryFeatureregistration,
		"reset":        CredentialUsageSummaryFeaturereset,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CredentialUsageSummaryFeature(input)
	return &out, nil
}
