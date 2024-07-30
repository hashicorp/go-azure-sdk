package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInAuthenticationProtocol string

const (
	SignInAuthenticationProtocol_AuthenticationTransfer SignInAuthenticationProtocol = "authenticationTransfer"
	SignInAuthenticationProtocol_DeviceCode             SignInAuthenticationProtocol = "deviceCode"
	SignInAuthenticationProtocol_None                   SignInAuthenticationProtocol = "none"
	SignInAuthenticationProtocol_OAuth2                 SignInAuthenticationProtocol = "oAuth2"
	SignInAuthenticationProtocol_Ropc                   SignInAuthenticationProtocol = "ropc"
	SignInAuthenticationProtocol_Saml20                 SignInAuthenticationProtocol = "saml20"
	SignInAuthenticationProtocol_WsFederation           SignInAuthenticationProtocol = "wsFederation"
)

func PossibleValuesForSignInAuthenticationProtocol() []string {
	return []string{
		string(SignInAuthenticationProtocol_AuthenticationTransfer),
		string(SignInAuthenticationProtocol_DeviceCode),
		string(SignInAuthenticationProtocol_None),
		string(SignInAuthenticationProtocol_OAuth2),
		string(SignInAuthenticationProtocol_Ropc),
		string(SignInAuthenticationProtocol_Saml20),
		string(SignInAuthenticationProtocol_WsFederation),
	}
}

func (s *SignInAuthenticationProtocol) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSignInAuthenticationProtocol(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSignInAuthenticationProtocol(input string) (*SignInAuthenticationProtocol, error) {
	vals := map[string]SignInAuthenticationProtocol{
		"authenticationtransfer": SignInAuthenticationProtocol_AuthenticationTransfer,
		"devicecode":             SignInAuthenticationProtocol_DeviceCode,
		"none":                   SignInAuthenticationProtocol_None,
		"oauth2":                 SignInAuthenticationProtocol_OAuth2,
		"ropc":                   SignInAuthenticationProtocol_Ropc,
		"saml20":                 SignInAuthenticationProtocol_Saml20,
		"wsfederation":           SignInAuthenticationProtocol_WsFederation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignInAuthenticationProtocol(input)
	return &out, nil
}
