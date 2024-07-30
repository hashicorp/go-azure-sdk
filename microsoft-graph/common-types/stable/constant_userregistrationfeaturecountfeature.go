package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserRegistrationFeatureCountFeature string

const (
	UserRegistrationFeatureCountFeature_MfaCapable          UserRegistrationFeatureCountFeature = "mfaCapable"
	UserRegistrationFeatureCountFeature_PasswordlessCapable UserRegistrationFeatureCountFeature = "passwordlessCapable"
	UserRegistrationFeatureCountFeature_SsprCapable         UserRegistrationFeatureCountFeature = "ssprCapable"
	UserRegistrationFeatureCountFeature_SsprEnabled         UserRegistrationFeatureCountFeature = "ssprEnabled"
	UserRegistrationFeatureCountFeature_SsprRegistered      UserRegistrationFeatureCountFeature = "ssprRegistered"
)

func PossibleValuesForUserRegistrationFeatureCountFeature() []string {
	return []string{
		string(UserRegistrationFeatureCountFeature_MfaCapable),
		string(UserRegistrationFeatureCountFeature_PasswordlessCapable),
		string(UserRegistrationFeatureCountFeature_SsprCapable),
		string(UserRegistrationFeatureCountFeature_SsprEnabled),
		string(UserRegistrationFeatureCountFeature_SsprRegistered),
	}
}

func (s *UserRegistrationFeatureCountFeature) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserRegistrationFeatureCountFeature(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserRegistrationFeatureCountFeature(input string) (*UserRegistrationFeatureCountFeature, error) {
	vals := map[string]UserRegistrationFeatureCountFeature{
		"mfacapable":          UserRegistrationFeatureCountFeature_MfaCapable,
		"passwordlesscapable": UserRegistrationFeatureCountFeature_PasswordlessCapable,
		"ssprcapable":         UserRegistrationFeatureCountFeature_SsprCapable,
		"ssprenabled":         UserRegistrationFeatureCountFeature_SsprEnabled,
		"ssprregistered":      UserRegistrationFeatureCountFeature_SsprRegistered,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserRegistrationFeatureCountFeature(input)
	return &out, nil
}
