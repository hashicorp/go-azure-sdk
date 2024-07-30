package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FeatureRolloutPolicyFeature string

const (
	FeatureRolloutPolicyFeature_CertificateBasedAuthentication FeatureRolloutPolicyFeature = "certificateBasedAuthentication"
	FeatureRolloutPolicyFeature_EmailAsAlternateId             FeatureRolloutPolicyFeature = "emailAsAlternateId"
	FeatureRolloutPolicyFeature_PassthroughAuthentication      FeatureRolloutPolicyFeature = "passthroughAuthentication"
	FeatureRolloutPolicyFeature_PasswordHashSync               FeatureRolloutPolicyFeature = "passwordHashSync"
	FeatureRolloutPolicyFeature_SeamlessSso                    FeatureRolloutPolicyFeature = "seamlessSso"
)

func PossibleValuesForFeatureRolloutPolicyFeature() []string {
	return []string{
		string(FeatureRolloutPolicyFeature_CertificateBasedAuthentication),
		string(FeatureRolloutPolicyFeature_EmailAsAlternateId),
		string(FeatureRolloutPolicyFeature_PassthroughAuthentication),
		string(FeatureRolloutPolicyFeature_PasswordHashSync),
		string(FeatureRolloutPolicyFeature_SeamlessSso),
	}
}

func (s *FeatureRolloutPolicyFeature) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFeatureRolloutPolicyFeature(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFeatureRolloutPolicyFeature(input string) (*FeatureRolloutPolicyFeature, error) {
	vals := map[string]FeatureRolloutPolicyFeature{
		"certificatebasedauthentication": FeatureRolloutPolicyFeature_CertificateBasedAuthentication,
		"emailasalternateid":             FeatureRolloutPolicyFeature_EmailAsAlternateId,
		"passthroughauthentication":      FeatureRolloutPolicyFeature_PassthroughAuthentication,
		"passwordhashsync":               FeatureRolloutPolicyFeature_PasswordHashSync,
		"seamlesssso":                    FeatureRolloutPolicyFeature_SeamlessSso,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FeatureRolloutPolicyFeature(input)
	return &out, nil
}
