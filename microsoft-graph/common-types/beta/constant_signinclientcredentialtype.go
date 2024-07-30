package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInClientCredentialType string

const (
	SignInClientCredentialType_Certificate                 SignInClientCredentialType = "certificate"
	SignInClientCredentialType_ClientAssertion             SignInClientCredentialType = "clientAssertion"
	SignInClientCredentialType_ClientSecret                SignInClientCredentialType = "clientSecret"
	SignInClientCredentialType_FederatedIdentityCredential SignInClientCredentialType = "federatedIdentityCredential"
	SignInClientCredentialType_ManagedIdentity             SignInClientCredentialType = "managedIdentity"
	SignInClientCredentialType_None                        SignInClientCredentialType = "none"
)

func PossibleValuesForSignInClientCredentialType() []string {
	return []string{
		string(SignInClientCredentialType_Certificate),
		string(SignInClientCredentialType_ClientAssertion),
		string(SignInClientCredentialType_ClientSecret),
		string(SignInClientCredentialType_FederatedIdentityCredential),
		string(SignInClientCredentialType_ManagedIdentity),
		string(SignInClientCredentialType_None),
	}
}

func (s *SignInClientCredentialType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSignInClientCredentialType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSignInClientCredentialType(input string) (*SignInClientCredentialType, error) {
	vals := map[string]SignInClientCredentialType{
		"certificate":                 SignInClientCredentialType_Certificate,
		"clientassertion":             SignInClientCredentialType_ClientAssertion,
		"clientsecret":                SignInClientCredentialType_ClientSecret,
		"federatedidentitycredential": SignInClientCredentialType_FederatedIdentityCredential,
		"managedidentity":             SignInClientCredentialType_ManagedIdentity,
		"none":                        SignInClientCredentialType_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignInClientCredentialType(input)
	return &out, nil
}
