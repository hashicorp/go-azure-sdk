package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInSignInIdentifierType string

const (
	SignInSignInIdentifierType_OnPremisesUserPrincipalName SignInSignInIdentifierType = "onPremisesUserPrincipalName"
	SignInSignInIdentifierType_PhoneNumber                 SignInSignInIdentifierType = "phoneNumber"
	SignInSignInIdentifierType_ProxyAddress                SignInSignInIdentifierType = "proxyAddress"
	SignInSignInIdentifierType_QrCode                      SignInSignInIdentifierType = "qrCode"
	SignInSignInIdentifierType_UserPrincipalName           SignInSignInIdentifierType = "userPrincipalName"
)

func PossibleValuesForSignInSignInIdentifierType() []string {
	return []string{
		string(SignInSignInIdentifierType_OnPremisesUserPrincipalName),
		string(SignInSignInIdentifierType_PhoneNumber),
		string(SignInSignInIdentifierType_ProxyAddress),
		string(SignInSignInIdentifierType_QrCode),
		string(SignInSignInIdentifierType_UserPrincipalName),
	}
}

func (s *SignInSignInIdentifierType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSignInSignInIdentifierType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSignInSignInIdentifierType(input string) (*SignInSignInIdentifierType, error) {
	vals := map[string]SignInSignInIdentifierType{
		"onpremisesuserprincipalname": SignInSignInIdentifierType_OnPremisesUserPrincipalName,
		"phonenumber":                 SignInSignInIdentifierType_PhoneNumber,
		"proxyaddress":                SignInSignInIdentifierType_ProxyAddress,
		"qrcode":                      SignInSignInIdentifierType_QrCode,
		"userprincipalname":           SignInSignInIdentifierType_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignInSignInIdentifierType(input)
	return &out, nil
}
