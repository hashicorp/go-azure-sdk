package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserRegistrationFeatureCountFeature string

const (
	UserRegistrationFeatureCountFeaturemfaCapable          UserRegistrationFeatureCountFeature = "MfaCapable"
	UserRegistrationFeatureCountFeaturepasswordlessCapable UserRegistrationFeatureCountFeature = "PasswordlessCapable"
	UserRegistrationFeatureCountFeaturessprCapable         UserRegistrationFeatureCountFeature = "SsprCapable"
	UserRegistrationFeatureCountFeaturessprEnabled         UserRegistrationFeatureCountFeature = "SsprEnabled"
	UserRegistrationFeatureCountFeaturessprRegistered      UserRegistrationFeatureCountFeature = "SsprRegistered"
)

func PossibleValuesForUserRegistrationFeatureCountFeature() []string {
	return []string{
		string(UserRegistrationFeatureCountFeaturemfaCapable),
		string(UserRegistrationFeatureCountFeaturepasswordlessCapable),
		string(UserRegistrationFeatureCountFeaturessprCapable),
		string(UserRegistrationFeatureCountFeaturessprEnabled),
		string(UserRegistrationFeatureCountFeaturessprRegistered),
	}
}

func parseUserRegistrationFeatureCountFeature(input string) (*UserRegistrationFeatureCountFeature, error) {
	vals := map[string]UserRegistrationFeatureCountFeature{
		"mfacapable":          UserRegistrationFeatureCountFeaturemfaCapable,
		"passwordlesscapable": UserRegistrationFeatureCountFeaturepasswordlessCapable,
		"ssprcapable":         UserRegistrationFeatureCountFeaturessprCapable,
		"ssprenabled":         UserRegistrationFeatureCountFeaturessprEnabled,
		"ssprregistered":      UserRegistrationFeatureCountFeaturessprRegistered,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserRegistrationFeatureCountFeature(input)
	return &out, nil
}
