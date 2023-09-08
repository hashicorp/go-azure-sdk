package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInClientCredentialType string

const (
	SignInClientCredentialTypecertificate                 SignInClientCredentialType = "Certificate"
	SignInClientCredentialTypeclientAssertion             SignInClientCredentialType = "ClientAssertion"
	SignInClientCredentialTypeclientSecret                SignInClientCredentialType = "ClientSecret"
	SignInClientCredentialTypefederatedIdentityCredential SignInClientCredentialType = "FederatedIdentityCredential"
	SignInClientCredentialTypemanagedIdentity             SignInClientCredentialType = "ManagedIdentity"
	SignInClientCredentialTypenone                        SignInClientCredentialType = "None"
)

func PossibleValuesForSignInClientCredentialType() []string {
	return []string{
		string(SignInClientCredentialTypecertificate),
		string(SignInClientCredentialTypeclientAssertion),
		string(SignInClientCredentialTypeclientSecret),
		string(SignInClientCredentialTypefederatedIdentityCredential),
		string(SignInClientCredentialTypemanagedIdentity),
		string(SignInClientCredentialTypenone),
	}
}

func parseSignInClientCredentialType(input string) (*SignInClientCredentialType, error) {
	vals := map[string]SignInClientCredentialType{
		"certificate":                 SignInClientCredentialTypecertificate,
		"clientassertion":             SignInClientCredentialTypeclientAssertion,
		"clientsecret":                SignInClientCredentialTypeclientSecret,
		"federatedidentitycredential": SignInClientCredentialTypefederatedIdentityCredential,
		"managedidentity":             SignInClientCredentialTypemanagedIdentity,
		"none":                        SignInClientCredentialTypenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignInClientCredentialType(input)
	return &out, nil
}
