package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCredentialUsageDetailsFeature string

const (
	UserCredentialUsageDetailsFeature_Registration UserCredentialUsageDetailsFeature = "registration"
	UserCredentialUsageDetailsFeature_Reset        UserCredentialUsageDetailsFeature = "reset"
)

func PossibleValuesForUserCredentialUsageDetailsFeature() []string {
	return []string{
		string(UserCredentialUsageDetailsFeature_Registration),
		string(UserCredentialUsageDetailsFeature_Reset),
	}
}

func (s *UserCredentialUsageDetailsFeature) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserCredentialUsageDetailsFeature(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserCredentialUsageDetailsFeature(input string) (*UserCredentialUsageDetailsFeature, error) {
	vals := map[string]UserCredentialUsageDetailsFeature{
		"registration": UserCredentialUsageDetailsFeature_Registration,
		"reset":        UserCredentialUsageDetailsFeature_Reset,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserCredentialUsageDetailsFeature(input)
	return &out, nil
}
