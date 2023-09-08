package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCredentialUsageDetailsFeature string

const (
	UserCredentialUsageDetailsFeatureregistration UserCredentialUsageDetailsFeature = "Registration"
	UserCredentialUsageDetailsFeaturereset        UserCredentialUsageDetailsFeature = "Reset"
)

func PossibleValuesForUserCredentialUsageDetailsFeature() []string {
	return []string{
		string(UserCredentialUsageDetailsFeatureregistration),
		string(UserCredentialUsageDetailsFeaturereset),
	}
}

func parseUserCredentialUsageDetailsFeature(input string) (*UserCredentialUsageDetailsFeature, error) {
	vals := map[string]UserCredentialUsageDetailsFeature{
		"registration": UserCredentialUsageDetailsFeatureregistration,
		"reset":        UserCredentialUsageDetailsFeaturereset,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserCredentialUsageDetailsFeature(input)
	return &out, nil
}
