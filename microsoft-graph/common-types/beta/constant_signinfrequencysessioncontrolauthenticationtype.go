package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInFrequencySessionControlAuthenticationType string

const (
	SignInFrequencySessionControlAuthenticationType_PrimaryAndSecondaryAuthentication SignInFrequencySessionControlAuthenticationType = "primaryAndSecondaryAuthentication"
	SignInFrequencySessionControlAuthenticationType_SecondaryAuthentication           SignInFrequencySessionControlAuthenticationType = "secondaryAuthentication"
)

func PossibleValuesForSignInFrequencySessionControlAuthenticationType() []string {
	return []string{
		string(SignInFrequencySessionControlAuthenticationType_PrimaryAndSecondaryAuthentication),
		string(SignInFrequencySessionControlAuthenticationType_SecondaryAuthentication),
	}
}

func (s *SignInFrequencySessionControlAuthenticationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSignInFrequencySessionControlAuthenticationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSignInFrequencySessionControlAuthenticationType(input string) (*SignInFrequencySessionControlAuthenticationType, error) {
	vals := map[string]SignInFrequencySessionControlAuthenticationType{
		"primaryandsecondaryauthentication": SignInFrequencySessionControlAuthenticationType_PrimaryAndSecondaryAuthentication,
		"secondaryauthentication":           SignInFrequencySessionControlAuthenticationType_SecondaryAuthentication,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignInFrequencySessionControlAuthenticationType(input)
	return &out, nil
}
