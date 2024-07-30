package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInIncomingTokenType string

const (
	SignInIncomingTokenType_None                SignInIncomingTokenType = "none"
	SignInIncomingTokenType_PrimaryRefreshToken SignInIncomingTokenType = "primaryRefreshToken"
	SignInIncomingTokenType_RemoteDesktopToken  SignInIncomingTokenType = "remoteDesktopToken"
	SignInIncomingTokenType_Saml11              SignInIncomingTokenType = "saml11"
	SignInIncomingTokenType_Saml20              SignInIncomingTokenType = "saml20"
)

func PossibleValuesForSignInIncomingTokenType() []string {
	return []string{
		string(SignInIncomingTokenType_None),
		string(SignInIncomingTokenType_PrimaryRefreshToken),
		string(SignInIncomingTokenType_RemoteDesktopToken),
		string(SignInIncomingTokenType_Saml11),
		string(SignInIncomingTokenType_Saml20),
	}
}

func (s *SignInIncomingTokenType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSignInIncomingTokenType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSignInIncomingTokenType(input string) (*SignInIncomingTokenType, error) {
	vals := map[string]SignInIncomingTokenType{
		"none":                SignInIncomingTokenType_None,
		"primaryrefreshtoken": SignInIncomingTokenType_PrimaryRefreshToken,
		"remotedesktoptoken":  SignInIncomingTokenType_RemoteDesktopToken,
		"saml11":              SignInIncomingTokenType_Saml11,
		"saml20":              SignInIncomingTokenType_Saml20,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignInIncomingTokenType(input)
	return &out, nil
}
