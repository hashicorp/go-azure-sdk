package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FeatureRolloutPolicyFeature string

const (
	FeatureRolloutPolicyFeaturecertificateBasedAuthentication FeatureRolloutPolicyFeature = "CertificateBasedAuthentication"
	FeatureRolloutPolicyFeatureemailAsAlternateId             FeatureRolloutPolicyFeature = "EmailAsAlternateId"
	FeatureRolloutPolicyFeaturepassthroughAuthentication      FeatureRolloutPolicyFeature = "PassthroughAuthentication"
	FeatureRolloutPolicyFeaturepasswordHashSync               FeatureRolloutPolicyFeature = "PasswordHashSync"
	FeatureRolloutPolicyFeatureseamlessSso                    FeatureRolloutPolicyFeature = "SeamlessSso"
)

func PossibleValuesForFeatureRolloutPolicyFeature() []string {
	return []string{
		string(FeatureRolloutPolicyFeaturecertificateBasedAuthentication),
		string(FeatureRolloutPolicyFeatureemailAsAlternateId),
		string(FeatureRolloutPolicyFeaturepassthroughAuthentication),
		string(FeatureRolloutPolicyFeaturepasswordHashSync),
		string(FeatureRolloutPolicyFeatureseamlessSso),
	}
}

func parseFeatureRolloutPolicyFeature(input string) (*FeatureRolloutPolicyFeature, error) {
	vals := map[string]FeatureRolloutPolicyFeature{
		"certificatebasedauthentication": FeatureRolloutPolicyFeaturecertificateBasedAuthentication,
		"emailasalternateid":             FeatureRolloutPolicyFeatureemailAsAlternateId,
		"passthroughauthentication":      FeatureRolloutPolicyFeaturepassthroughAuthentication,
		"passwordhashsync":               FeatureRolloutPolicyFeaturepasswordHashSync,
		"seamlesssso":                    FeatureRolloutPolicyFeatureseamlessSso,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FeatureRolloutPolicyFeature(input)
	return &out, nil
}
